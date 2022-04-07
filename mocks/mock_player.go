package mocks

import (
	"context"
	"errors"

	"mega-tournament-api/services"
)

type mockPlayerService struct{}

func (f *mockPlayerService) Create(ctx context.Context) (*services.PlayerResponse, error) {
	return nil, nil
}

func (f *mockPlayerService) Update(ctx context.Context) (*services.PlayerResponse, error) {
	return nil, nil
}

func (f *mockPlayerService) Delete(ctx context.Context) (*services.PlayerResponse, error) {
	return nil, nil
}

func (f *mockPlayerService) Get(ctx context.Context) (*services.PlayerResponse, error) {
	return nil, nil
}

type mockPlayerServiceWithErr struct{}

func (f *mockPlayerServiceWithErr) Create(ctx context.Context) (*services.PlayerResponse, error) {
	return nil, errors.New(FakePostErr)
}

func (f *mockPlayerServiceWithErr) Update(ctx context.Context) (*services.PlayerResponse, error) {
	return nil, errors.New(FakePatchErr)
}

func (f *mockPlayerServiceWithErr) Get(ctx context.Context) (*services.PlayerResponse, error) {
	return nil, errors.New(FakeGetErr)
}

func (f *mockPlayerServiceWithErr) Delete(ctx context.Context) (*services.PlayerResponse, error) {
	return nil, errors.New(FakeDeleteErr)
}
