package services

import "context"

type PlayerService interface {
	Get(ctx context.Context) (*PlayerResponse, error)
	Create(ctx context.Context) (*PlayerResponse, error)
	Update(ctx context.Context) (*PlayerResponse, error)
	Delete(ctx context.Context) (*PlayerResponse, error)
}

type playerService struct{}

type PlayerPayload interface{}
type PlayerResponse struct{}

func NewPlayerService() (PlayerService, error) {
	return &playerService{}, nil
}

func (b *playerService) Get(ctx context.Context) (*PlayerResponse, error) {
	return nil, nil
}

func (b *playerService) Create(ctx context.Context) (*PlayerResponse, error) {
	return nil, nil
}

func (b *playerService) Update(ctx context.Context) (*PlayerResponse, error) {
	return nil, nil
}

func (b *playerService) Delete(ctx context.Context) (*PlayerResponse, error) {
	return nil, nil
}
