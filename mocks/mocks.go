package mocks

import "mega-tournament-api/services"

const FakeGetErr string = "FAKE GET ERROR"
const FakePostErr string = "FAKE POST ERROR"
const FakeDeleteErr string = "FAKE DELETE ERROR"
const FakePatchErr string = "FAKE PATCH ERROR"

func GetBracketMock(shouldError bool) services.BracketService {
	if shouldError {
		return &mockBracketServiceWithErr{}
	}

	return &mockBracketService{}
}

func GetPlayerMock(shouldError bool) services.PlayerService {
	if shouldError {
		return &mockPlayerServiceWithErr{}
	}

	return &mockPlayerService{}
}

func GetTeamMock(shouldError bool) services.TeamService {
	if shouldError {
		return &mockTeamServiceWithErr{}
	}

	return &mockTeamService{}
}
