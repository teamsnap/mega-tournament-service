package services

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type PlayerTestSuite struct {
	suite.Suite
	ctx context.Context
	svc PlayerService
}

func (suite *PlayerTestSuite) SetupTest() {
	suite.ctx = context.Background()
	suite.svc, _ = NewPlayerService()
}

func (suite *PlayerTestSuite) TestGetPlayer() {
	_, err := suite.svc.Get(suite.ctx)
	assert.NoError(suite.T(), err, "Failed getting Player.")
}

func (suite *PlayerTestSuite) TestUpdatePlayer() {
	_, err := suite.svc.Update(suite.ctx)
	assert.NoError(suite.T(), err, "Failed updating Player.")
}

func (suite *PlayerTestSuite) TestDeletePlayer() {
	_, err := suite.svc.Delete(suite.ctx)
	assert.NoError(suite.T(), err, "Failed deleting Player.")
}

func (suite *PlayerTestSuite) TestCreatePlayer() {
	_, err := suite.svc.Create(suite.ctx)
	assert.NoError(suite.T(), err, "Failed creating Player.")
}

func TestPlayerTestSuite(t *testing.T) {
	suite.Run(t, new(PlayerTestSuite))
}
