package service

import "Service-for-assigning-reviewers-for-Pull-Requests/internal/entity"

type UserService struct {
	repo UserRepository
}

func NewUserService(r UserRepository) *UserService {
	return &UserService{repo: r}
}

func (s *UserService) SetUserActive(userID int64, active bool) (*entity.User, error) {
	return s.repo.UpdateActive(userID, active)
}

func (s *UserService) GetPRsAssignedTo(userID int64) ([]entity.PullRequest, error) {
	return s.repo.GetPRsForReviewer(userID)
}
