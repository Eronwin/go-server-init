package repository

import "go-server-init/internal/model"

//数据访问层

type Repository struct{}

func New() *Repository {
	return &Repository{}
}

func (r *Repository) FetchBase() (*model.PingResponse, error) {
	return &model.PingResponse{
		Message: "pong",
	}, nil
}
