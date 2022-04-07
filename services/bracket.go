package services

import "context"

type BracketService interface {
	Get(ctx context.Context) (*BracketResponse, error)
	Create(ctx context.Context) (*BracketResponse, error)
	Update(ctx context.Context) (*BracketResponse, error)
	Delete(ctx context.Context) (*BracketResponse, error)
}

type bracketService struct{}

type BracketPayload struct{}
type BracketResponse struct{}

func NewBracketService() (BracketService, error) {
	return &bracketService{}, nil
}

func (b *bracketService) Get(ctx context.Context) (*BracketResponse, error) {
	return nil, nil
}

func (b bracketService) Create(ctx context.Context) (*BracketResponse, error) {
	return nil, nil
}

func (b *bracketService) Update(ctx context.Context) (*BracketResponse, error) {
	return nil, nil
}

func (b *bracketService) Delete(ctx context.Context) (*BracketResponse, error) {
	return nil, nil
}
