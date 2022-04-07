package services

import "context"

type TeamService interface {
	Get(ctx context.Context) (*TeamResponse, error)
	Create(ctx context.Context) (*TeamResponse, error)
	Update(ctx context.Context) (*TeamResponse, error)
	Delete(ctx context.Context) (*TeamResponse, error)
}

type teamService struct{}

type TeamPayload interface{}
type TeamResponse struct{}

func NewTeamService() (TeamService, error) {
	return &teamService{}, nil
}

func (b *teamService) Get(ctx context.Context) (*TeamResponse, error) {
	return nil, nil
}

func (b *teamService) Create(ctx context.Context) (*TeamResponse, error) {
	return nil, nil
}

func (b *teamService) Update(ctx context.Context) (*TeamResponse, error) {
	return nil, nil
}

func (b *teamService) Delete(ctx context.Context) (*TeamResponse, error) {
	return nil, nil
}
