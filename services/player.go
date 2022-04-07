package services

import "context"

type PlayerService struct{}
type PlayerPayload interface{}
type PlayerResponse struct{}

func NewPlayerService() (PlayerService, error) {
	return PlayerService{}, nil
}

func (b *PlayerService) Get(ctx context.Context) (*PlayerResponse, error) {
	return nil, nil
}

func (b *PlayerService) Create(ctx context.Context) (*PlayerResponse, error) {
	return nil, nil
}

func (b *PlayerService) Modify(ctx context.Context) (*PlayerResponse, error) {
	return nil, nil
}

func (b *PlayerService) Delete(ctx context.Context) (*PlayerResponse, error) {
	return nil, nil
}
