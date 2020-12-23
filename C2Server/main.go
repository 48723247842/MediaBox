package main

import (
	fiber "github.com/gofiber/fiber/v2"
	buttons_route_handler "c2server/routes/buttons"
	state_spotify_route_handler "c2server/routes/states/spotify"
	state_local_tvshow_route_handler "c2server/routes/states/local/tvshow"
)

func main() {

	app := fiber.New()

	// Top Level Control
	// app.Get( "/pause" )
	// app.Get( "/resume" )
	// app.Get( "/play" )
	// app.Get( "/stop" )
	// app.Get( "/next" )
	// app.Get( "/previous" )
	// app.Get( "/restart" )

	// Named Top Level Macro States
	app.Get( "/spotify" , buttons_route_handler.Button1 )

	// Button Routes
	buttons := app.Group( "/button" )
	buttons.Get( "/1" , buttons_route_handler.Button1 )
	buttons.Get( "/spotify" , buttons_route_handler.Button1 )

	// Micro State Routes
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

	states_local := states.Group( "/local" )
	states_local_tvshow := states_local.Group( "/tvshow" )
	states_local_tvshow.Get( "/start" , state_local_tvshow_route_handler.Start )
	states_local_tvshow.Get( "/teardown" , state_local_tvshow_route_handler.Teardown )
	states_local_tvshow.Get( "/play" , state_local_tvshow_route_handler.Play )
	states_local_tvshow.Get( "/stop" , state_local_tvshow_route_handler.Stop )
	states_local_tvshow.Get( "/pause" , state_local_tvshow_route_handler.Pause )
	states_local_tvshow.Get( "/resume" , state_local_tvshow_route_handler.Resume )
	states_local_tvshow.Get( "/next" , state_local_tvshow_route_handler.Next )
	states_local_tvshow.Get( "/previous" , state_local_tvshow_route_handler.Previous )
	states_local_tvshow.Get( "/status" , state_local_tvshow_route_handler.Status )

	app.Listen( ":9363" )

}