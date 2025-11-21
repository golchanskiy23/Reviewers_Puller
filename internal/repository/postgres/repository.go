package postgres

import "Service-for-assigning-reviewers-for-Pull-Requests/pkg/database/postgres"

type Repository struct {
	Teams        TeamRepository
	Users        UserRepository
	PullRequests PullRequestRepository
}

func NewRepository(db *postgres.DatabaseSource) *Repository {
	return &Repository{
		Teams:        NewTeamPGRepository(db),
		Users:        NewUserPGRepository(db),
		PullRequests: NewPullRequestPGRepository(db),
	}
}
