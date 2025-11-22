package entity

import "time"

type PRReviewer struct {
	PullRequestID string    `db:"pull_request_id"`
	ReviewerID    string    `db:"reviewer_id"`
	AssignedAt    time.Time `db:"assigned_at"`
}
