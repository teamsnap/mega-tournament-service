package mocks

import (
	"context"
	"errors"

	"mega-tournament-api/services"
)

type mockTeamService struct{}

func (f *mockTeamService) Create(ctx context.Context) (*services.TeamResponse, error) {
	return nil, nil
}

func (f *mockTeamService) Update(ctx context.Context) (*services.TeamResponse, error) {
	return nil, nil
}

func (f *mockTeamService) Delete(ctx context.Context) (*services.TeamResponse, error) {
	return nil, nil
}

func (f *mockTeamService) Get(ctx context.Context) (*services.TeamResponse, error) {
	return nil, nil
}

type mockTeamServiceWithErr struct{}

func (f *mockTeamServiceWithErr) Create(ctx context.Context) (*services.TeamResponse, error) {
	return nil, errors.New(FakePostErr)
}

func (f *mockTeamServiceWithErr) Update(ctx context.Context) (*services.TeamResponse, error) {
	return nil, errors.New(FakePatchErr)
}

func (f *mockTeamServiceWithErr) Get(ctx context.Context) (*services.TeamResponse, error) {
	return nil, errors.New(FakeGetErr)
}

func (f *mockTeamServiceWithErr) Delete(ctx context.Context) (*services.TeamResponse, error) {
	return nil, errors.New(FakeDeleteErr)
}
