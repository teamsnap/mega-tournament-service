package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (c *Container) GetPlayer(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, "Status OK")
}

func (c *Container) CreatePlayer(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, "Status OK")
}

func (c *Container) UpdatePlayer(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, "Status OK")
}

func (c *Container) DeletePlayer(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, "Status OK")
}
