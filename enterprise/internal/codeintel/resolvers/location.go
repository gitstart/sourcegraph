package resolvers

import (
	"context"

	"github.com/sourcegraph/go-langserver/pkg/lsp"
	"github.com/sourcegraph/sourcegraph/cmd/frontend/graphqlbackend"
	"github.com/sourcegraph/sourcegraph/cmd/frontend/graphqlbackend/graphqlutil"
	"github.com/sourcegraph/sourcegraph/cmd/frontend/types"
	"github.com/sourcegraph/sourcegraph/internal/api"
	"github.com/sourcegraph/sourcegraph/internal/lsif"
)

type locationConnectionResolver struct {
	repo      *types.Repo
	commit    api.CommitID
	locations []*lsif.LSIFLocation
	endCursor string
}

var _ graphqlbackend.LocationConnectionResolver = &locationConnectionResolver{}

func (r *locationConnectionResolver) Nodes(ctx context.Context) ([]graphqlbackend.LocationResolver, error) {
	collectionResolver := &repositoryCollectionResolver{
		commitCollectionResolvers: map[api.RepoID]*commitCollectionResolver{},
	}

	var l []graphqlbackend.LocationResolver
	for _, location := range r.locations {
		adjustedCommit, adjustedRange, err := r.adjustLocation(ctx, location)
		if err != nil {
			return nil, err
		}

		treeResolver, err := collectionResolver.resolve(ctx, location.RepositoryID, adjustedCommit, location.Path)
		if err != nil {
			return nil, err
		}

		if treeResolver == nil {
			continue
		}

		l = append(l, graphqlbackend.NewLocationResolver(treeResolver, &adjustedRange))
	}

	return l, nil
}

func (r *locationConnectionResolver) PageInfo(ctx context.Context) (*graphqlutil.PageInfo, error) {
	if r.endCursor != "" {
		return graphqlutil.NextPageCursor(r.endCursor), nil
	}
	return graphqlutil.HasNextPage(false), nil
}

// adjustLocation attempts to transform the source range of location into a corresponding
// range of the same file at the user's requested commit.
//
// If location has no corresponding range at the requested commit or is located in a different
// repository, it returns the location's current commit and range without modification.
// Otherwise, it returns the user's requested commit along with the transformed range.
//
// A non-nil error means the connection resolver was unable to load the diff between
// the requested commit and location's commit.
func (r *locationConnectionResolver) adjustLocation(ctx context.Context, location *lsif.LSIFLocation) (string, lsp.Range, error) {
	if location.RepositoryID != r.repo.ID {
		return location.Commit, location.Range, nil
	}

	adjuster, err := newPositionAdjuster(ctx, r.repo, location.Commit, string(r.commit), location.Path)
	if err != nil {
		return "", lsp.Range{}, err
	}

	if adjustedRange, ok := adjuster.adjustRange(location.Range); ok {
		return string(r.commit), adjustedRange, nil
	}

	// Couldn't adjust range, return original result which is precise but
	// jump the user to another into another commit context on navigation.
	return location.Commit, location.Range, nil
}
