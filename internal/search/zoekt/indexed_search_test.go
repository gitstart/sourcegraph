package zoekt

import (
	"context"
	"crypto/md5"
	"encoding/binary"
	"testing"
	"time"

	"github.com/cockroachdb/errors"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/google/zoekt"

	"github.com/sourcegraph/sourcegraph/internal/api"
	"github.com/sourcegraph/sourcegraph/internal/search"
	searchbackend "github.com/sourcegraph/sourcegraph/internal/search/backend"
	"github.com/sourcegraph/sourcegraph/internal/search/query"
	"github.com/sourcegraph/sourcegraph/internal/search/result"
	"github.com/sourcegraph/sourcegraph/internal/search/streaming"
	"github.com/sourcegraph/sourcegraph/internal/trace"
	"github.com/sourcegraph/sourcegraph/internal/types"
)

func TestIndexedSearch(t *testing.T) {
	zeroTimeoutCtx, cancel := context.WithTimeout(context.Background(), 0)
	defer cancel()
	type args struct {
		ctx             context.Context
		query           string
		patternInfo     *search.TextPatternInfo
		repos           *search.Repos
		useFullDeadline bool
		results         []zoekt.FileMatch
		since           func(time.Time) time.Duration
	}

	reposHEAD := mkSearchRepos("foo/bar", "foo/foobar")
	zoektRepos := []*zoekt.RepoListEntry{{
		Repository: zoekt.Repository{
			Name:     "foo/bar",
			Branches: []zoekt.RepositoryBranch{{Name: "HEAD", Version: "barHEADSHA"}, {Name: "dev", Version: "bardevSHA"}, {Name: "main", Version: "barmainSHA"}},
		},
	}, {
		Repository: zoekt.Repository{
			Name:     "foo/foobar",
			Branches: []zoekt.RepositoryBranch{{Name: "HEAD", Version: "foobarHEADSHA"}},
		},
	}}

	tests := []struct {
		name               string
		args               args
		wantMatchCount     int
		wantMatchKeys      []result.Key
		wantMatchInputRevs []string
		wantUnindexed      *search.Repos
		wantCommon         streaming.Stats
		wantErr            bool
	}{
		{
			name: "no matches",
			args: args{
				ctx:             context.Background(),
				patternInfo:     &search.TextPatternInfo{},
				repos:           reposHEAD,
				useFullDeadline: false,
				since:           func(time.Time) time.Duration { return time.Second - time.Millisecond },
			},
			wantErr: false,
		},
		{
			name: "no matches timeout",
			args: args{
				ctx:             context.Background(),
				patternInfo:     &search.TextPatternInfo{},
				repos:           reposHEAD,
				useFullDeadline: false,
				since:           func(time.Time) time.Duration { return time.Minute },
			},
			wantCommon: streaming.Stats{
				Status: mkStatusMap(map[string]search.RepoStatus{
					"foo/bar":    search.RepoStatusTimedout,
					"foo/foobar": search.RepoStatusTimedout,
				}),
			},
		},
		{
			name: "context timeout",
			args: args{
				ctx:             zeroTimeoutCtx,
				patternInfo:     &search.TextPatternInfo{},
				repos:           reposHEAD,
				useFullDeadline: true,
				since:           func(time.Time) time.Duration { return 0 },
			},
			wantErr: true,
		},
		{
			name: "results",
			args: args{
				ctx:             context.Background(),
				patternInfo:     &search.TextPatternInfo{FileMatchLimit: 100},
				repos:           mkSearchRepos("foo/bar", "foo/foobar"),
				useFullDeadline: false,
				results: []zoekt.FileMatch{
					{
						Repository: "foo/bar",
						Branches:   []string{"HEAD"},
						Version:    "1",
						FileName:   "baz.go",
						LineMatches: []zoekt.LineMatch{
							{
								Line: []byte("I'm like 1.5+ hours into writing this test :'("),
								LineFragments: []zoekt.LineFragmentMatch{
									{LineOffset: 0, MatchLength: 5},
								},
							},
							{
								Line: []byte("I'm ready for the rain to stop."),
								LineFragments: []zoekt.LineFragmentMatch{
									{LineOffset: 0, MatchLength: 5},
									{LineOffset: 5, MatchLength: 10},
								},
							},
						},
					},
					{
						Repository: "foo/foobar",
						Branches:   []string{"HEAD"},
						Version:    "2",
						FileName:   "baz.go",
						LineMatches: []zoekt.LineMatch{
							{
								Line: []byte("s/rain/pain"),
								LineFragments: []zoekt.LineFragmentMatch{
									{LineOffset: 0, MatchLength: 5},
									{LineOffset: 5, MatchLength: 2},
								},
							},
						},
					},
				},
				since: func(time.Time) time.Duration { return 0 },
			},
			wantMatchCount: 5,
			wantMatchKeys: []result.Key{
				{Repo: "foo/bar", Commit: "1", Path: "baz.go"},
				{Repo: "foo/foobar", Commit: "2", Path: "baz.go"},
			},
			wantMatchInputRevs: []string{
				"",
				"",
			},
			wantErr: false,
		},
		{
			name: "results multi-branch",
			args: args{
				ctx:             context.Background(),
				patternInfo:     &search.TextPatternInfo{FileMatchLimit: 100},
				repos:           mkSearchRepos("foo/bar@HEAD:dev:main"),
				useFullDeadline: false,
				results: []zoekt.FileMatch{
					{
						Repository: "foo/bar",
						// baz.go is the same in HEAD and dev
						Branches: []string{"HEAD", "dev"},
						FileName: "baz.go",
						Version:  "1",
					},
					{
						Repository: "foo/bar",
						Branches:   []string{"dev"},
						FileName:   "bam.go",
						Version:    "2",
					},
				},
				since: func(time.Time) time.Duration { return 0 },
			},
			wantMatchCount: 3,
			wantMatchKeys: []result.Key{
				{Repo: "foo/bar", Commit: "1", Path: "baz.go"},
				{Repo: "foo/bar", Commit: "1", Path: "baz.go"},
				{Repo: "foo/bar", Commit: "2", Path: "bam.go"},
			},
			wantMatchInputRevs: []string{
				"HEAD",
				"dev",
				"dev",
			},
			wantErr: false,
		},
		{
			// if we search a branch that is indexed and unindexed, we should
			// split the repository revision into the indexed and unindexed
			// parts.
			name: "split branch",
			args: args{
				ctx:             context.Background(),
				patternInfo:     &search.TextPatternInfo{FileMatchLimit: 100},
				repos:           mkSearchRepos("foo/bar@HEAD:unindexed"),
				useFullDeadline: false,
				results: []zoekt.FileMatch{
					{
						Repository: "foo/bar",
						Branches:   []string{"HEAD"},
						FileName:   "baz.go",
						Version:    "1",
					},
				},
			},
			wantUnindexed: mkSearchRepos("foo/bar@unindexed"),
			wantMatchKeys: []result.Key{
				{Repo: "foo/bar", Commit: "1", Path: "baz.go"},
			},
			wantMatchCount:     1,
			wantMatchInputRevs: []string{"HEAD"},
		},
		{
			// Fallback to unindexed search if the query contains ref-globs.
			name: "ref-glob with explicit /*",
			args: args{
				ctx:             context.Background(),
				query:           "repo:foo/bar@*refs/heads/*",
				patternInfo:     &search.TextPatternInfo{FileMatchLimit: 100},
				repos:           mkSearchRepos("foo/bar@HEAD"),
				useFullDeadline: false,
				results:         []zoekt.FileMatch{},
			},
			wantUnindexed:      mkSearchRepos("foo/bar@HEAD"),
			wantMatchKeys:      nil,
			wantMatchInputRevs: nil,
		},
		{
			name: "ref-glob with implicit /*",
			args: args{
				ctx:             context.Background(),
				query:           "repo:foo/bar@*refs/tags",
				patternInfo:     &search.TextPatternInfo{FileMatchLimit: 100},
				repos:           mkSearchRepos("foo/bar@HEAD"),
				useFullDeadline: false,
				results:         []zoekt.FileMatch{},
			},
			wantUnindexed:      mkSearchRepos("foo/bar@HEAD"),
			wantMatchKeys:      nil,
			wantMatchInputRevs: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q, err := query.ParseLiteral(tt.args.query)
			if err != nil {
				t.Fatal(err)
			}

			args := &search.TextParameters{
				Query:           q,
				PatternInfo:     tt.args.patternInfo,
				Repos:           tt.args.repos,
				UseFullDeadline: tt.args.useFullDeadline,
				Zoekt: &searchbackend.Zoekt{
					Client: &searchbackend.FakeSearcher{
						Result: &zoekt.SearchResult{Files: tt.args.results},
						Repos:  zoektRepos,
					},
					DisableCache: true,
				},
			}

			indexed, err := NewIndexedSubsetSearchRequest(context.Background(), args, search.TextRequest, MissingRepoRevStatus(streaming.StreamFunc(func(streaming.SearchEvent) {})))
			if err != nil {
				t.Fatal(err)
			}

			wantUnindexed := tt.wantUnindexed
			if wantUnindexed == nil {
				wantUnindexed = search.NewRepos()
			}

			if diff := cmp.Diff(wantUnindexed, indexed.Unindexed, cmpopts.EquateEmpty()); diff != "" {
				t.Errorf("unindexed mismatch (-want +got):\n%s", diff)
			}

			indexed.since = tt.args.since

			// This is a quick fix which will break once we enable the zoekt client for true streaming.
			// Once we return more than one event we have to account for the proper order of results
			// in the tests.
			gotMatches, gotCommon, err := streaming.CollectStream(func(stream streaming.Sender) error {
				return indexed.Search(tt.args.ctx, stream)
			})
			if (err != nil) != tt.wantErr {
				t.Errorf("zoektSearchHEAD() error = %v, wantErr = %v", err, tt.wantErr)
				return
			}

			gotFm, err := matchesToFileMatches(gotMatches)
			if err != nil {
				t.Fatal(err)
			}

			if diff := cmp.Diff(&tt.wantCommon, &gotCommon, cmpopts.EquateEmpty()); diff != "" {
				t.Errorf("common mismatch (-want +got):\n%s", diff)
			}

			var gotMatchCount int
			var gotMatchKeys []result.Key
			var gotMatchInputRevs []string
			for _, m := range gotFm {
				gotMatchCount += m.ResultCount()
				gotMatchKeys = append(gotMatchKeys, m.Key())
				if m.InputRev != nil {
					gotMatchInputRevs = append(gotMatchInputRevs, *m.InputRev)
				}
			}
			if diff := cmp.Diff(tt.wantMatchKeys, gotMatchKeys); diff != "" {
				t.Errorf("match URLs mismatch (-want +got):\n%s", diff)
			}
			if diff := cmp.Diff(tt.wantMatchInputRevs, gotMatchInputRevs); diff != "" {
				t.Errorf("match InputRevs mismatch (-want +got):\n%s", diff)
			}
			if gotMatchCount != tt.wantMatchCount {
				t.Errorf("gotMatchCount = %v, want %v", gotMatchCount, tt.wantMatchCount)
			}
		})
	}
}

func mkStatusMap(m map[string]search.RepoStatus) search.RepoStatusMap {
	var rsm search.RepoStatusMap
	for name, status := range m {
		rsm.Update(mkRepos(name)[0].ID, status)
	}
	return rsm
}

func TestZoektIndexedRepos(t *testing.T) {
	repos := []string{
		"foo/indexed-one@",
		"foo/indexed-two@",
		"foo/indexed-three@",
		"foo/unindexed-one",
		"foo/unindexed-two",
		"foo/multi-rev@a:b",
	}

	zoektRepos := map[string]*zoekt.Repository{}
	for _, r := range []*zoekt.Repository{{
		Name:     "foo/indexed-one",
		Branches: []zoekt.RepositoryBranch{{Name: "HEAD", Version: "deadbeef"}},
	}, {
		Name:     "foo/indexed-two",
		Branches: []zoekt.RepositoryBranch{{Name: "HEAD", Version: "deadbeef"}},
	}, {
		Name: "foo/indexed-three",
		Branches: []zoekt.RepositoryBranch{
			{Name: "HEAD", Version: "deadbeef"},
			{Name: "foobar", Version: "deadcow"},
		},
	}} {
		zoektRepos[r.Name] = r
	}

	mkIndexedRepoRevs := func(repos []string) *IndexedRepos {
		idx := &IndexedRepos{
			Repos:    mkSearchRepos(repos...),
			RevSpecs: map[api.RepoName]search.RevSpecs{},
			Branches: map[string][]string{},
		}

		for repo, revs := range idx.RepoRevs {
			idx.RevSpecs[repo] = revs
		}

		return idx
	}

	cases := []struct {
		name             string
		indexed          *IndexedRepos
		repos, unindexed *search.Repos
	}{{
		name:      "all",
		repos:     mkSearchRepos(repos...),
		indexed:   mkIndexedRepoRevs(repos[:3]),
		unindexed: mkSearchRepos(repos[3:]...),
	}, {
		name:      "one unindexed",
		repos:     mkSearchRepos(repos[3:4]...),
		indexed:   mkIndexedRepoRevs(repos[:0]),
		unindexed: mkSearchRepos(repos[3:4]...),
	}, {
		name:      "one indexed",
		repos:     mkSearchRepos(repos[:1]...),
		indexed:   mkIndexedRepoRevs(repos[:1]),
		unindexed: mkSearchRepos(repos[:0]...),
	}}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			indexed, unindexed := zoektIndexedRepos(zoektRepos, tc.repos, nil)

			if diff := cmp.Diff(tc.indexed.RevSpecs, indexed.RevSpecs); diff != "" {
				t.Error("unexpected indexed:", diff)
			}
			if diff := cmp.Diff(tc.unindexed, unindexed); diff != "" {
				t.Error("unexpected unindexed:", diff)
			}
		})
	}
}

func TestZoektResultCountFactor(t *testing.T) {
	cases := []struct {
		name         string
		numRepos     int
		globalSearch bool
		pattern      *search.TextPatternInfo
		want         int
	}{
		{
			name:     "One repo implies max scaling factor",
			numRepos: 1,
			pattern:  &search.TextPatternInfo{},
			want:     100,
		},
		{
			name:     "Eleven repos implies a scaling factor between min and max",
			numRepos: 11,
			pattern:  &search.TextPatternInfo{},
			want:     8,
		},
		{
			name:     "More than 500 repos implies a min scaling factor",
			numRepos: 501,
			pattern:  &search.TextPatternInfo{},
			want:     1,
		},
		{
			name:     "Setting a count greater than defautl max results (30) adapts scaling factor",
			numRepos: 501,
			pattern:  &search.TextPatternInfo{FileMatchLimit: 100},
			want:     10,
		},
		{
			name:         "for global searches, k should be 1",
			numRepos:     0,
			globalSearch: true,
			pattern:      &search.TextPatternInfo{},
			want:         1,
		},
		{
			name:         "for global searches, k should be 1, adjusted by the FileMatchLimit",
			numRepos:     0,
			globalSearch: true,
			pattern:      &search.TextPatternInfo{FileMatchLimit: 100},
			want:         10,
		},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			got := ResultCountFactor(tt.numRepos, tt.pattern.FileMatchLimit, tt.globalSearch)
			if tt.want != got {
				t.Fatalf("Want scaling factor %d but got %d", tt.want, got)
			}
		})
	}
}

func TestZoektIndexedRepos_single(t *testing.T) {
	repoRev := func(revSpec string) *search.Repos {
		rs := search.NewRepos()
		rs.Add(
			&types.RepoName{ID: api.RepoID(0), Name: "test/repo"},
			search.RevisionSpecifier{RevSpec: revSpec},
		)
		return rs
	}

	zoektRepos := map[string]*zoekt.Repository{
		"test/repo": {
			Name: "test/repo",
			Branches: []zoekt.RepositoryBranch{
				{
					Name:    "HEAD",
					Version: "df3f4e499698e48152b39cd655d8901eaf583fa5",
				},
				{
					Name:    "NOT-HEAD",
					Version: "8ec975423738fe7851676083ebf660a062ed1578",
				},
			},
		},
	}
	cases := []struct {
		rev           string
		wantIndexed   *search.Repos
		wantUnindexed *search.Repos
	}{
		{
			rev:           "",
			wantIndexed:   repoRev(""),
			wantUnindexed: search.NewRepos(),
		},
		{
			rev:           "HEAD",
			wantIndexed:   repoRev("HEAD"),
			wantUnindexed: search.NewRepos(),
		},
		{
			rev:           "df3f4e499698e48152b39cd655d8901eaf583fa5",
			wantIndexed:   repoRev("df3f4e499698e48152b39cd655d8901eaf583fa5"),
			wantUnindexed: search.NewRepos(),
		},
		{
			rev:           "df3f4e",
			wantIndexed:   repoRev("df3f4e"),
			wantUnindexed: search.NewRepos(),
		},
		{
			rev:           "d",
			wantIndexed:   search.NewRepos(),
			wantUnindexed: repoRev("d"),
		},
		{
			rev:           "HEAD^1",
			wantIndexed:   search.NewRepos(),
			wantUnindexed: repoRev("HEAD^1"),
		},
		{
			rev:           "8ec975423738fe7851676083ebf660a062ed1578",
			wantUnindexed: search.NewRepos(),
			wantIndexed:   repoRev("8ec975423738fe7851676083ebf660a062ed1578"),
		},
	}

	type ret struct {
		Indexed   map[api.RepoName]search.RevSpecs
		Unindexed *search.Repos
	}

	for _, tt := range cases {
		indexed, unindexed := zoektIndexedRepos(zoektRepos, repoRev(tt.rev), nil)
		got := ret{
			Indexed:   indexed.RevSpecs,
			Unindexed: unindexed,
		}
		want := ret{
			Indexed:   tt.wantIndexed.RepoRevs,
			Unindexed: tt.wantUnindexed,
		}
		if !cmp.Equal(want, got) {
			t.Errorf("%s mismatch (-want +got):\n%s", tt.rev, cmp.Diff(want, got))
		}
	}
}

func TestZoektFileMatchToSymbolResults(t *testing.T) {
	symbolInfo := func(sym string) *zoekt.Symbol {
		return &zoekt.Symbol{
			Sym:        sym,
			Kind:       "kind",
			Parent:     "parent",
			ParentKind: "parentkind",
		}
	}

	file := &zoekt.FileMatch{
		FileName:   "bar.go",
		Repository: "foo",
		Language:   "go",
		Version:    "deadbeef",
		LineMatches: []zoekt.LineMatch{{
			// Skips missing symbol info (shouldn't happen in practice).
			Line:          []byte(""),
			LineNumber:    5,
			LineFragments: []zoekt.LineFragmentMatch{{}},
		}, {
			Line:       []byte("symbol a symbol b"),
			LineNumber: 10,
			LineFragments: []zoekt.LineFragmentMatch{{
				SymbolInfo: symbolInfo("a"),
			}, {
				SymbolInfo: symbolInfo("b"),
			}},
		}, {
			Line:       []byte("symbol c"),
			LineNumber: 15,
			LineFragments: []zoekt.LineFragmentMatch{{
				SymbolInfo: symbolInfo("c"),
			}},
		}, {
			Line:       []byte(`bar() { var regex = /.*\//; function baz() { }  } `),
			LineNumber: 20,
			LineFragments: []zoekt.LineFragmentMatch{{
				SymbolInfo: symbolInfo("baz"),
			}},
		}},
	}

	results := zoektFileMatchToSymbolResults(&types.RepoName{Name: "foo"}, "master", file)
	var symbols []result.Symbol
	for _, res := range results {
		symbols = append(symbols, res.Symbol)
	}

	want := []result.Symbol{
		{
			Name:    "a",
			Line:    10,
			Pattern: "/^symbol a symbol b$/",
		}, {
			Name:    "b",
			Line:    10,
			Pattern: "/^symbol a symbol b$/",
		}, {
			Name:    "c",
			Line:    15,
			Pattern: "/^symbol c$/",
		}, {
			Name:    "baz",
			Line:    20,
			Pattern: `/^bar() { var regex = \/.*\\\/\/; function baz() { }  } $/`,
		},
	}
	for i := range want {
		want[i].Kind = "kind"
		want[i].Parent = "parent"
		want[i].ParentKind = "parentkind"
		want[i].Path = "bar.go"
		want[i].Language = "go"
	}

	if diff := cmp.Diff(want, symbols); diff != "" {
		t.Fatalf("symbol mismatch (-want +got):\n%s", diff)
	}
}

func TestContextWithoutDeadline(t *testing.T) {
	ctxWithDeadline, cancelWithDeadline := context.WithTimeout(context.Background(), time.Minute)
	defer cancelWithDeadline()

	tr, ctxWithDeadline := trace.New(ctxWithDeadline, "", "")

	if _, ok := ctxWithDeadline.Deadline(); !ok {
		t.Fatal("expected context to have deadline")
	}

	ctxNoDeadline, cancelNoDeadline := contextWithoutDeadline(ctxWithDeadline)
	defer cancelNoDeadline()

	if _, ok := ctxNoDeadline.Deadline(); ok {
		t.Fatal("expected context to not have deadline")
	}

	// We want to keep trace info
	if tr2 := trace.TraceFromContext(ctxNoDeadline); tr != tr2 {
		t.Error("trace information not propogated")
	}

	// Calling cancelWithDeadline should cancel ctxNoDeadline
	cancelWithDeadline()
	select {
	case <-ctxNoDeadline.Done():
	case <-time.After(10 * time.Second):
		t.Fatal("expected context to be done")
	}
}

func TestContextWithoutDeadline_cancel(t *testing.T) {
	ctxWithDeadline, cancelWithDeadline := context.WithTimeout(context.Background(), time.Minute)
	defer cancelWithDeadline()
	ctxNoDeadline, cancelNoDeadline := contextWithoutDeadline(ctxWithDeadline)

	cancelNoDeadline()
	select {
	case <-ctxNoDeadline.Done():
	case <-time.After(10 * time.Second):
		t.Fatal("expected context to be done")
	}
}

func mkSearchRepos(repos ...string) *search.Repos {
	resolved := search.NewRepos()
	for _, repospec := range repos {
		repoName, revs := search.ParseRepositoryRevisions(repospec)
		if len(revs) == 0 {
			// treat empty list as preferring master
			revs = search.RevSpecs{{RevSpec: ""}}
		}
		rs := mkRepos(repoName)
		resolved.Add(&rs[0], revs...)
	}
	return resolved
}

func mkRepos(names ...string) []types.RepoName {
	var repos []types.RepoName
	for _, name := range names {
		sum := md5.Sum([]byte(name))
		id := api.RepoID(binary.BigEndian.Uint64(sum[:]))
		if id < 0 {
			id = -(id / 2)
		}
		if id == 0 {
			id++
		}
		repos = append(repos, types.RepoName{ID: id, Name: api.RepoName(name)})
	}
	return repos
}

func matchesToFileMatches(matches []result.Match) ([]*result.FileMatch, error) {
	fms := make([]*result.FileMatch, 0, len(matches))
	for _, match := range matches {
		fm, ok := match.(*result.FileMatch)
		if !ok {
			return nil, errors.Errorf("expected only file match results")
		}
		fms = append(fms, fm)
	}
	return fms, nil
}
