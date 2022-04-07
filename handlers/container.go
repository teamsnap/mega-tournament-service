package handlers

import "mega-tournament-api/services"

// Container will hold all dependencies for your application.
type Container struct {
	bracketService services.BracketService
	playerService  services.PlayerService
	teamService    services.TeamService
}

// NewContainer returns an empty or an initialized container for your handlers.
func NewContainer(
	bracketService services.BracketService,
	playerService services.PlayerService,
	teamService services.TeamService,
) (Container, error) {
	c := Container{
		bracketService,
		playerService,
		teamService,
	}
	return c, nil
}
