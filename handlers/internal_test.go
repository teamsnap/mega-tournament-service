package handlers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type InternalTestSuite struct {
	suite.Suite
	echo     *echo.Echo
	recorder *httptest.ResponseRecorder
}

func (suite *InternalTestSuite) SetupTest() {
	e := echo.New()
	suite.echo = e
	suite.recorder = httptest.NewRecorder()
}

func (suite *InternalTestSuite) TestGetHealth() {
	// Mock context
	req := httptest.NewRequest(
		http.MethodGet,
		"/bracket",
		strings.NewReader(""),
	)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	ctx := suite.echo.NewContext(req, suite.recorder)

	// Assertions
	assert.NoError(suite.T(), Health(ctx), "Failed getting Internal.")
	assert.Equal(suite.T(), 200, suite.recorder.Code, "Invalid Response code.")
}

func (suite *InternalTestSuite) TestGetSpec() {
	// Mock context
	req := httptest.NewRequest(
		http.MethodGet,
		"/bracket",
		strings.NewReader(""),
	)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	ctx := suite.echo.NewContext(req, suite.recorder)

	// Assertions
	assert.NoError(suite.T(), Spec(ctx), "Failed getting Internal.")
	assert.Equal(suite.T(), 200, suite.recorder.Code, "Invalid Response code.")
}

func TestInternalTestSuite(t *testing.T) {
	suite.Run(t, new(InternalTestSuite))
}
