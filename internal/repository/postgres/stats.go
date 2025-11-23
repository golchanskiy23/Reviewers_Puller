package postgres

import (
	"context"
	"time"

	"Service-for-assigning-reviewers-for-Pull-Requests/pkg/database"
)

// StatsRepository provides aggregated read-only queries used for metrics
type StatsRepository interface {
	GetAssignedReviewersCountPerPR(ctx context.Context) (map[string]int, error)
	GetOpenPRCountPerUser(ctx context.Context) (map[string]int, error)
}

type statsPGRepository struct {
	db *database.DatabaseSource
}

func NewStatsPGRepository(db *database.DatabaseSource) StatsRepository {
	return &statsPGRepository{db: db}
}

// GetAssignedReviewersCountPerPR returns number of assigned reviewers for each OPEN pull request
func (r *statsPGRepository) GetAssignedReviewersCountPerPR(ctx context.Context) (map[string]int, error) {
	// Use a short timeout for metrics queries to avoid long-running scrapes
	qctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	rows, err := r.db.Pool.Query(qctx,
		`SELECT pr.pull_request_id, COUNT(prr.reviewer_id) AS cnt
		 FROM pull_requests pr
		 LEFT JOIN pr_reviewers prr ON pr.pull_request_id = prr.pull_request_id
		 WHERE pr.status = 'OPEN'
		 GROUP BY pr.pull_request_id`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	res := make(map[string]int)
	for rows.Next() {
		var prID string
		var cnt int
		if err := rows.Scan(&prID, &cnt); err != nil {
			return nil, err
		}
		res[prID] = cnt
	}

	return res, nil
}

// GetOpenPRCountPerUser returns number of OPEN PRs for which each user is currently assigned as reviewer
func (r *statsPGRepository) GetOpenPRCountPerUser(ctx context.Context) (map[string]int, error) {
	// Use a short timeout for metrics queries to avoid long-running scrapes
	qctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	rows, err := r.db.Pool.Query(qctx,
		`SELECT prr.reviewer_id, COUNT(DISTINCT pr.pull_request_id) AS cnt
		 FROM pr_reviewers prr
		 JOIN pull_requests pr ON pr.pull_request_id = prr.pull_request_id
		 WHERE pr.status = 'OPEN'
		 GROUP BY prr.reviewer_id`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	res := make(map[string]int)
	for rows.Next() {
		var userID string
		var cnt int
		if err := rows.Scan(&userID, &cnt); err != nil {
			return nil, err
		}
		res[userID] = cnt
	}

	return res, nil
}
