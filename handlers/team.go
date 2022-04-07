package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (c *Container) GetTeam(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, "Status OK")
}

func (c *Container) CreateTeam(ctx echo.Context) error {
	return ctx.JSON(http.StatusCreated, "Status OK")
}

func (c *Container) UpdateTeam(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, "Status OK")
}

func (c *Container) DeleteTeam(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, "Status OK")
}
