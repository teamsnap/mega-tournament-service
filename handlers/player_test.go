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

type PlayerTestSuite struct {
	suite.Suite
	service  Container
	echo     *echo.Echo
	recorder *httptest.ResponseRecorder
}

func (suite *PlayerTestSuite) SetupTest() {
	underTest, _ := services.NewPlayerService()

	e := echo.New()
	h, _ := NewContainer(services.BracketService{}, underTest, services.TeamService{})

	suite.service = h
	suite.echo = e
	suite.recorder = httptest.NewRecorder()
}

func (suite *PlayerTestSuite) TestGetPlayer() {
	// Mock context
	req := httptest.NewRequest(
		http.MethodGet,
		"/bracket",
		strings.NewReader(""),
	)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	ctx := suite.echo.NewContext(req, suite.recorder)

	// Assertions
	assert.NoError(suite.T(), suite.service.GetPlayer(ctx), "Failed getting Player.")
	assert.Equal(suite.T(), 200, suite.recorder.Code, "Invalid Response code.")
}

func (suite *PlayerTestSuite) TestUpdatePlayer() {
	req := httptest.NewRequest(
		http.MethodPatch,
		"/bracket",
		strings.NewReader("{}"),
	)

	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	ctx := suite.echo.NewContext(req, suite.recorder)

	assert.NoError(suite.T(), suite.service.UpdatePlayer(ctx), "Failed getting Player.")
	assert.Equal(suite.T(), 200, suite.recorder.Code, "Invalid Response code.")
}

func (suite *PlayerTestSuite) TestCreatePlayer() {
	req := httptest.NewRequest(
		http.MethodPost,
		"/bracket",
		strings.NewReader("{}"),
	)

	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	ctx := suite.echo.NewContext(req, suite.recorder)

	assert.NoError(suite.T(), suite.service.CreatePlayer(ctx), "Failed getting Player.")
	assert.Equal(suite.T(), 201, suite.recorder.Code, "Invalid Response code.")
}

func (suite *PlayerTestSuite) TestDeletePlayer() {
	req := httptest.NewRequest(
		http.MethodDelete,
		"/bracket",
		strings.NewReader("{}"),
	)

	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	ctx := suite.echo.NewContext(req, suite.recorder)

	assert.NoError(suite.T(), suite.service.DeletePlayer(ctx), "Failed getting Player.")
	assert.Equal(suite.T(), 200, suite.recorder.Code, "Invalid Response code.")
}

func TestPlayerTestSuite(t *testing.T) {
	suite.Run(t, new(PlayerTestSuite))
}
