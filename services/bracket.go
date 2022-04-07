package services

import "context"

type BracketService struct{}

type BracketPayload struct{}
type BracketResponse struct{}

func NewBracketService() (BracketService, error) {
	return BracketService{}, nil
}

func (b *BracketService) Get(ctx context.Context) (*BracketResponse, error) {
	return nil, nil
}

func (b BracketService) Create(ctx context.Context) (*BracketResponse, error) {
	return nil, nil
}

func (b *BracketService) Modify(ctx context.Context) (*BracketResponse, error) {
	return nil, nil
}

func (b *BracketService) Delete(ctx context.Context) (*BracketResponse, error) {
	return nil, nil
}
