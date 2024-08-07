package store

import (
	"context"
	"fmt"
	"time"

	"github.com/keegancsmith/sqlf"
	"go.opentelemetry.io/otel/attribute"

	"github.com/sourcegraph/sourcegraph/internal/api"
	"github.com/sourcegraph/sourcegraph/internal/codeintel/core"
	"github.com/sourcegraph/sourcegraph/internal/database/basestore"
	"github.com/sourcegraph/sourcegraph/internal/database/dbutil"
	"github.com/sourcegraph/sourcegraph/internal/observation"
)

type CommitWithDate struct {
	Commit        api.CommitID
	CommitterDate time.Time
}

// GetCommitDateForOldestUpload returns the oldest commit date for all uploads for the given repository.
//   - If there are any null values, the commit date backfill job has not yet completed and
//     an error is returned to prevent downstream expiration errors being made due to outdated commit graph data.
//   - Otherwise if there are no non-nil timestamps, Some is returned.
func (s *store) GetCommitAndDateForOldestUpload(ctx context.Context, repositoryID int) (_ core.Option[CommitWithDate], err error) {
	ctx, _, endObservation := s.operations.getCommitAndDateForOldestUpload.With(ctx, &err, observation.Args{Attrs: []attribute.KeyValue{
		attribute.Int("repositoryID", repositoryID),
	}})
	defer endObservation(1, observation.Args{})

	var none = core.None[CommitWithDate]()
	type commitWithNullDate struct {
		commit string
		t      *time.Time
	}
	data, ok, err := basestore.NewFirstScanner(func(scanner dbutil.Scanner) (commitWithNullDate, error) {
		var commit string
		var t *time.Time
		err := scanner.Scan(&commit, &t)
		return commitWithNullDate{commit, t}, err
	})(s.db.Query(ctx, sqlf.Sprintf(getCommitAndDateForOldestUploadQuery, repositoryID)))

	if err != nil || !ok {
		return none, err
	}
	if data.t == nil {
		return none, &backfillIncompleteError{repositoryID}
	}

	return core.Some(CommitWithDate{Commit: api.CommitID(data.commit), CommitterDate: *data.t}), nil
}

// Note: we check against '-infinity' here, as the backfill operation will use this sentinel value in the case
// that the commit is no longer know by gitserver. This allows the backfill migration to make progress without
// having pristine database.
const getCommitAndDateForOldestUploadQuery = `
SELECT
	u.commit, cd.committed_at
FROM lsif_uploads u
LEFT JOIN codeintel_commit_dates cd ON cd.repository_id = u.repository_id AND cd.commit_bytea = decode(u.commit, 'hex')
WHERE
	u.repository_id = %s AND
	u.state = 'completed' AND
	(cd.committed_at != '-infinity' OR cd.committed_at IS NULL)
ORDER BY cd.committed_at NULLS FIRST
LIMIT 1
`

// UpdateCommittedAt updates the committed_at column for upload matching the given repository and commit.
func (s *store) UpdateCommittedAt(ctx context.Context, repositoryID int, commit, commitDateString string) (err error) {
	ctx, _, endObservation := s.operations.updateCommittedAt.With(ctx, &err, observation.Args{Attrs: []attribute.KeyValue{
		attribute.Int("repositoryID", repositoryID),
		attribute.String("commit", commit),
	}})
	defer func() { endObservation(1, observation.Args{}) }()

	return s.db.Exec(ctx, sqlf.Sprintf(updateCommittedAtQuery, repositoryID, dbutil.CommitBytea(commit), commitDateString))
}

const updateCommittedAtQuery = `
INSERT INTO codeintel_commit_dates(repository_id, commit_bytea, committed_at) VALUES (%s, %s, %s) ON CONFLICT DO NOTHING
`

// SourcedCommitsWithoutCommittedAt returns the repository and commits of uploads that do not have an
// associated commit date value.
func (s *store) SourcedCommitsWithoutCommittedAt(ctx context.Context, batchSize int) (_ []SourcedCommits, err error) {
	ctx, _, endObservation := s.operations.sourcedCommitsWithoutCommittedAt.With(ctx, &err, observation.Args{Attrs: []attribute.KeyValue{
		attribute.Int("batchSize", batchSize),
	}})
	defer func() { endObservation(1, observation.Args{}) }()

	batchOfCommits, err := scanSourcedCommits(s.db.Query(ctx, sqlf.Sprintf(sourcedCommitsWithoutCommittedAtQuery, batchSize)))
	if err != nil {
		return nil, err
	}

	return batchOfCommits, nil
}

const sourcedCommitsWithoutCommittedAtQuery = `
SELECT u.repository_id, r.name, u.commit
FROM lsif_uploads u
JOIN repo r ON r.id = u.repository_id
LEFT JOIN codeintel_commit_dates cd ON cd.repository_id = u.repository_id AND cd.commit_bytea = decode(u.commit, 'hex')
WHERE u.state = 'completed' AND cd.committed_at IS NULL
GROUP BY u.repository_id, r.name, u.commit
ORDER BY repository_id, commit
LIMIT %s
`

//
//

type backfillIncompleteError struct {
	repositoryID int
}

func (e backfillIncompleteError) Error() string {
	return fmt.Sprintf("repository %d has not yet completed its backfill of commit dates", e.repositoryID)
}
