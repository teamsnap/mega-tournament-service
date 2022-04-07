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

type BracketTestSuite struct {
	suite.Suite
	service  Container
	echo     *echo.Echo
	recorder *httptest.ResponseRecorder
}

func (suite *BracketTestSuite) SetupTest() {
	underTest, _ := services.NewBracketService()

	e := echo.New()
	h, _ := NewContainer(underTest, services.PlayerService{}, services.TeamService{})

	suite.service = h
	suite.echo = e
	suite.recorder = httptest.NewRecorder()
}

func (suite *BracketTestSuite) TestGetBracket() {
	// Mock context
	req := httptest.NewRequest(
		http.MethodGet,
		"/bracket",
		strings.NewReader(""),
	)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	ctx := suite.echo.NewContext(req, suite.recorder)

	// Assertions
	assert.NoError(suite.T(), suite.service.GetBracket(ctx), "Failed getting Bracket.")
	assert.Equal(suite.T(), 200, suite.recorder.Code, "Invalid Response code.")
}

func (suite *BracketTestSuite) TestUpdateBracket() {
	req := httptest.NewRequest(
		http.MethodPatch,
		"/bracket",
		strings.NewReader("{}"),
	)

	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	ctx := suite.echo.NewContext(req, suite.recorder)

	assert.NoError(suite.T(), suite.service.UpdateBracket(ctx), "Failed getting Bracket.")
	assert.Equal(suite.T(), 200, suite.recorder.Code, "Invalid Response code.")
}

func (suite *BracketTestSuite) TestCreateBracket() {
	req := httptest.NewRequest(
		http.MethodPost,
		"/bracket",
		strings.NewReader("{}"),
	)

	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	ctx := suite.echo.NewContext(req, suite.recorder)

	assert.NoError(suite.T(), suite.service.CreateBracket(ctx), "Failed getting Bracket.")
	assert.Equal(suite.T(), 201, suite.recorder.Code, "Invalid Response code.")
}

func (suite *BracketTestSuite) TestDeleteBracket() {
	req := httptest.NewRequest(
		http.MethodDelete,
		"/bracket",
		strings.NewReader("{}"),
	)

	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	ctx := suite.echo.NewContext(req, suite.recorder)

	assert.NoError(suite.T(), suite.service.DeleteBracket(ctx), "Failed getting Bracket.")
	assert.Equal(suite.T(), 200, suite.recorder.Code, "Invalid Response code.")
}

func TestBracketTestSuite(t *testing.T) {
	suite.Run(t, new(BracketTestSuite))
}
