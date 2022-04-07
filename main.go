package main

import (
	"fmt"
	"os"

	"mega-tournament-api/handlers"
	"mega-tournament-api/services"

	"github.com/labstack/echo/v4"
)

// Injected from the ENV
// Pretend that we're injected these!
const PORT string = ":3000"

// End: Injected from the ENV

func main() {
	e := echo.New()

	bracketService, err := services.NewBracketService()
	if err != nil {
		fmt.Println("failed to create bracket service, err: ", err)
		os.Exit(1)
	}

	playerService, err := services.NewPlayerService()
	if err != nil {
		fmt.Println("failed to create player service, err: ", err)
		os.Exit(1)
	}

	teamService, err := services.NewTeamService()
	if err != nil {
		fmt.Println("failed to create team service, err: ", err)
		os.Exit(1)
	}

	c, err := handlers.NewContainer(
		bracketService,
		playerService,
		teamService,
	)

	if err != nil {
		fmt.Println("failed to create new container, err: ", err)
		os.Exit(1)
	}

	// Player
	e.GET("/player", c.GetPlayer)
	e.POST("/player", c.CreatePlayer)
	e.PATCH("/player", c.UpdatePlayer)
	e.DELETE("/player", c.DeletePlayer)

	// Bracket
	e.GET("/bracket", c.GetBracket)
	e.POST("/bracket", c.CreateBracket)
	e.PATCH("/bracket", c.UpdateBracket)
	e.DELETE("/bracket", c.DeleteBracket)

	// Team
	e.GET("/team", c.GetTeam)
	e.POST("/team", c.CreateTeam)
	e.PATCH("/team", c.UpdateTeam)
	e.DELETE("/team", c.DeleteTeam)

	// Internal
	e.GET("/spec", handlers.Spec)
	e.GET("/health", handlers.Health)

	e.Logger.Fatal(e.Start(PORT))
}
