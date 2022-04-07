package handlers

import (
	"mega-tournament-api/services"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type TeamTestSuite struct {
	suite.Suite
	service  Container
	echo     *echo.Echo
	recorder *httptest.ResponseRecorder
}

func (suite *TeamTestSuite) SetupTest() {
	underTest, _ := services.NewTeamService()

	e := echo.New()
	h, _ := NewContainer(services.BracketService{}, services.PlayerService{}, underTest)

	suite.service = h
	suite.echo = e
	suite.recorder = httptest.NewRecorder()
}

func (suite *TeamTestSuite) TestGetTeam() {
	// Mock context
	req := httptest.NewRequest(
		http.MethodGet,
		"/bracket",
		strings.NewReader(""),
	)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	ctx := suite.echo.NewContext(req, suite.recorder)

	// Assertions
	assert.NoError(suite.T(), suite.service.GetTeam(ctx), "Failed getting Team.")
	assert.Equal(suite.T(), 200, suite.recorder.Code, "Invalid Response code.")
}

func (suite *TeamTestSuite) TestUpdateTeam() {
	req := httptest.NewRequest(
		http.MethodPatch,
		"/bracket",
		strings.NewReader("{}"),
	)

	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	ctx := suite.echo.NewContext(req, suite.recorder)

	assert.NoError(suite.T(), suite.service.UpdateTeam(ctx), "Failed getting Team.")
	assert.Equal(suite.T(), 200, suite.recorder.Code, "Invalid Response code.")
}

func (suite *TeamTestSuite) TestCreateTeam() {
	req := httptest.NewRequest(
		http.MethodPost,
		"/bracket",
		strings.NewReader("{}"),
	)

	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	ctx := suite.echo.NewContext(req, suite.recorder)

	assert.NoError(suite.T(), suite.service.CreateTeam(ctx), "Failed getting Team.")
	assert.Equal(suite.T(), 201, suite.recorder.Code, "Invalid Response code.")
}

func (suite *TeamTestSuite) TestDeleteTeam() {
	req := httptest.NewRequest(
		http.MethodDelete,
		"/bracket",
		strings.NewReader("{}"),
	)

	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	ctx := suite.echo.NewContext(req, suite.recorder)

	assert.NoError(suite.T(), suite.service.DeleteTeam(ctx), "Failed getting Team.")
	assert.Equal(suite.T(), 200, suite.recorder.Code, "Invalid Response code.")
}

func TestTeamTestSuite(t *testing.T) {
	suite.Run(t, new(TeamTestSuite))
}
