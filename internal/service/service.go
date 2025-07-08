package service

import (
	"go-server-init/internal/model"
	"go-server-init/internal/repository"
	"go-server-init/internal/utils"
	"go-server-init/pkg/errors"
)

// 业务层

type Service struct {
	repo *repository.Repository
}

func New(repo *repository.Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Ping() (*model.PingResponse, error) {
	base, err := s.repo.FetchBase()
	if err != nil {
		return nil, errors.Wrap(err, "repo.FetchBase failed")
	}
	base.UUID = utils.GenerateUUID()
	return base, nil
}
