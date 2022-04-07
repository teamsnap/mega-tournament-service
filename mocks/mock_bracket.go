package mocks

import (
	"context"
	"errors"

	"mega-tournament-api/services"
)

type mockBracketService struct{}

func (f *mockBracketService) Create(ctx context.Context) (*services.BracketResponse, error) {
	return nil, nil
}

func (f *mockBracketService) Update(ctx context.Context) (*services.BracketResponse, error) {
	return nil, nil
}

func (f *mockBracketService) Delete(ctx context.Context) (*services.BracketResponse, error) {
	return nil, nil
}

func (f *mockBracketService) Get(ctx context.Context) (*services.BracketResponse, error) {
	return nil, nil
}

type mockBracketServiceWithErr struct{}

func (f *mockBracketServiceWithErr) Create(ctx context.Context) (*services.BracketResponse, error) {
	return nil, errors.New(FakePostErr)
}

func (f *mockBracketServiceWithErr) Update(ctx context.Context) (*services.BracketResponse, error) {
	return nil, errors.New(FakePatchErr)
}

func (f *mockBracketServiceWithErr) Get(ctx context.Context) (*services.BracketResponse, error) {
	return nil, errors.New(FakeGetErr)
}

func (f *mockBracketServiceWithErr) Delete(ctx context.Context) (*services.BracketResponse, error) {
	return nil, errors.New(FakeDeleteErr)
}
