package app

import "goprojstructtest/internal/domain"

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) GetUser(id domain.UserID) (*domain.User, error) {
	return nil, nil
}
