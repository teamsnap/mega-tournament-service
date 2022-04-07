package services

import "context"

type TeamService struct{}
type TeamPayload interface{}
type TeamResponse struct{}

func NewTeamService() (TeamService, error) {
	return TeamService{}, nil
}

func (b *TeamService) Get(ctx context.Context) (*TeamService, error) {
	return nil, nil
}

func (b *TeamService) Create(ctx context.Context) (*TeamService, error) {
	return nil, nil
}

func (b *TeamService) Modify(ctx context.Context) (*TeamService, error) {
	return nil, nil
}

func (b *TeamService) Delete(ctx context.Context) (*TeamService, error) {
	return nil, nil
}
