package entity

import "time"

type PullRequest struct {
	ID        string     `db:"pull_request_id"`
	Name      string     `db:"pull_request_name"`
	AuthorID  string     `db:"author_id"`
	Status    string     `db:"status"`
	CreatedAt time.Time  `db:"created_at"`
	MergedAt  *time.Time `db:"merged_at"`

	Reviewers []string
}
