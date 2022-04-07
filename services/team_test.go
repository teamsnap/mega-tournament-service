package services

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type TeamTestSuite struct {
	suite.Suite
	ctx context.Context
	svc TeamService
}

func (suite *TeamTestSuite) SetupTest() {
	suite.ctx = context.Background()
	suite.svc, _ = NewTeamService()
}

func (suite *TeamTestSuite) TestGetTeam() {
	_, err := suite.svc.Get(suite.ctx)
	assert.NoError(suite.T(), err, "Failed getting Team.")
}

func (suite *TeamTestSuite) TestUpdateTeam() {
	_, err := suite.svc.Update(suite.ctx)
	assert.NoError(suite.T(), err, "Failed updating Team.")
}

func (suite *TeamTestSuite) TestDeleteTeam() {
	_, err := suite.svc.Delete(suite.ctx)
	assert.NoError(suite.T(), err, "Failed deleting Team.")
}

func (suite *TeamTestSuite) TestCreateTeam() {
	_, err := suite.svc.Create(suite.ctx)
	assert.NoError(suite.T(), err, "Failed creating Team.")
}

func TestTeamTestSuite(t *testing.T) {
	suite.Run(t, new(TeamTestSuite))
}
