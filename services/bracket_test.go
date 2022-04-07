package services

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type BracketTestSuite struct {
	suite.Suite
	ctx context.Context
	svc BracketService
}

func (suite *BracketTestSuite) SetupTest() {
	suite.ctx = context.Background()
	suite.svc, _ = NewBracketService()
}

func (suite *BracketTestSuite) TestGetBracket() {
	_, err := suite.svc.Get(suite.ctx)
	assert.NoError(suite.T(), err, "Failed getting Bracket.")
}

func (suite *BracketTestSuite) TestUpdateBracket() {
	_, err := suite.svc.Update(suite.ctx)
	assert.NoError(suite.T(), err, "Failed updating Bracket.")
}

func (suite *BracketTestSuite) TestDeleteBracket() {
	_, err := suite.svc.Delete(suite.ctx)
	assert.NoError(suite.T(), err, "Failed deleting Bracket.")
}

func (suite *BracketTestSuite) TestCreateBracket() {
	_, err := suite.svc.Create(suite.ctx)
	assert.NoError(suite.T(), err, "Failed creating Bracket.")
}

func TestBracketTestSuite(t *testing.T) {
	suite.Run(t, new(BracketTestSuite))
}
