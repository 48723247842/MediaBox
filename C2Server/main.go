package main

import (
	fiber "github.com/gofiber/fiber/v2"
	state_spotify_route_handler "c2server/routes/states/spotify"
)

func main() {

	app := fiber.New()

	// States Routes
	states := app.Group( "/states" )

	// States Spotify Routes
	states_spotify := states.Group( "/spotify" )
	states_spotify.Get( "/start" , state_spotify_route_handler.Start )
	states_spotify.Get( "/teardown" , state_spotify_route_handler.Teardown )
	states_spotify.Get( "/play" , state_spotify_route_handler.Play )
	states_spotify.Get( "/stop" , state_spotify_route_handler.Stop )
	states_spotify.Get( "/pause" , state_spotify_route_handler.Pause )
	states_spotify.Get( "/resume" , state_spotify_route_handler.Resume )
	states_spotify.Get( "/next" , state_spotify_route_handler.Next )
	states_spotify.Get( "/previous" , state_spotify_route_handler.Previous )
	states_spotify.Get( "/status" , state_spotify_route_handler.Status )

	 app.Listen( ":9363" )

}