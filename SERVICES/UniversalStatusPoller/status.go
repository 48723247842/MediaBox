package main

import (
	"fmt"
	redis "github.com/0187773933/RedisManagerUtils/manager"
	gabs "github.com/Jeffail/gabs/v2"
	spotify_dbus "github.com/0187773933/SpotifyDBUSController/controller"
)

func GenericSpotify( current_state *gabs.Container ) ( result string ) {
	fmt.Println( "GenericSpotify()" )
	result = "failed"
	spotify := spotify_dbus.Controller{}
	spotify.Connect()
	spotify.UpdateStatus()
	fmt.Println( spotify.Status )
	result = spotify.PlaybackStatus()
	// if result != "Playing" {
	// 	var context = context.Background()
	// 	redis_connection := redis_lib.NewClient( &redis_lib.Options{
	// 		Addr: "localhost:6379" ,
	// 		DB: 3 ,
	// 		Password: "" ,
	// 	})
	// 	get_current_state_restart_on_fail_flag , get_current_state_restart_on_fail_flag_error := redis_connection.Get( context , "STATE.CURRENT.NAME" ).Result()
	// 	if get_current_state_restart_on_fail_flag_error != nil {} else {
	// 		if get_current_state_restart_on_fail_flag == "true" {
	// 			// GET http://localhost:C2ServerPort/state/restart?statename=spotify
	// 		}
	// 	}

	// }
	return
}

func GenericLocalTVShow( current_state *gabs.Container ) {
	fmt.Println( current_state )
}

// Switch To Just This , and then try to start in PM2
// https://github.com/onatm/clockwerk
func main() {
	redis := redis.Manager{}
	redis.Connect( "localhost:6379" , 3 , "" )

	current_state_string := redis.Get( "STATE.CURRENT" )
	current_state , current_state_error := gabs.ParseJSON( []byte( current_state_string ) )
	if current_state_error != nil { fmt.Println( current_state_error ); return }
	//fmt.Println( current_state )

	generic_type_data := current_state.Search( "GenericType" ).Data()
	if generic_type_data == nil { return }
	generic_type := generic_type_data.(string)
	switch generic_type {
		case "LocalTVShow":
			GenericLocalTVShow( current_state )
		case "Spotify":
			GenericSpotify( current_state )
		default:
			fmt.Println( "No Active States" )
			return
	}

}