package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Health(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, "Status OK")
}

func Spec(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, "Status OK")
}
