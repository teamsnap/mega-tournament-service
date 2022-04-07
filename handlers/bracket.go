package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (c *Container) GetBracket(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, "Status OK")
}

func (c *Container) CreateBracket(ctx echo.Context) error {
	return ctx.JSON(http.StatusCreated, "Status OK")
}

func (c *Container) UpdateBracket(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, "Status OK")
}

func (c *Container) DeleteBracket(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, "Status OK")
}
