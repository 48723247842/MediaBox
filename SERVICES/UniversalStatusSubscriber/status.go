package main

import (
	"fmt"
	"reflect"
	//"context"
	//logrus "github.com/sirupsen/logrus"
	redis "github.com/0187773933/RedisManagerUtils/manager"
	gabs "github.com/Jeffail/gabs/v2"
	//spotify_dbus "github.com/0187773933/SpotifyDBUSController/controller"
)

// func GenericSpotify() ( result string ) {
// 	fmt.Println( "GenericSpotify()" )
// 	result = "failed"
// 	spotify := spotify_dbus.Controller{}
// 	spotify.Connect()
// 	spotify.UpdateStatus()
// 	fmt.Println( spotify.Status )
// 	result = spotify.PlaybackStatus()
// 	if result != "Playing" {
// 		var context = context.Background()
// 		redis_connection := redis_lib.NewClient( &redis_lib.Options{
// 			Addr: "localhost:6379" ,
// 			DB: 3 ,
// 			Password: "" ,
// 		})
// 		get_current_state_restart_on_fail_flag , get_current_state_restart_on_fail_flag_error := redis_connection.Get( context , "STATE.CURRENT.NAME" ).Result()
// 		if get_current_state_restart_on_fail_flag_error != nil {} else {
// 			if get_current_state_restart_on_fail_flag == "true" {
// 				// GET http://localhost:C2ServerPort/state/restart?statename=spotify
// 			}
// 		}

// 	}
// 	return
// }

func log_all_handler( message string ) {
	//fmt.Println( message )
	json_parsed , err := gabs.ParseJSON( []byte( message ) )
	if err != nil { fmt.Println( err ); return }
	// fmt.Println( json_parsed )
	//for key , child := range json_parsed.ChildrenMap() {
	// for key , child := range json_parsed.Search("Fields").ChildrenMap() {
	// 	fmt.Println( key )
	// 	fmt.Println( child )

	// }

	// 1.) Get Command Key
	command := json_parsed.Search( "Fields" , "command" ).Data()
	if command == nil { return }
	command_string := string( command.(string) )
	fmt.Println( command_string )

	// 2.) Get Data Stored In Command Key
	// command_result := json_parsed.Search( "Fields" , command_string ).Data()
	// if command_result == nil { return }
	// if reflect.TypeOf( command_result ).String() == "string" {
	// 	command_result_string := command_result.(string)
	// 	//fmt.Println( command_result_string )
	// 	command_json_parsed , command_json_parsed_error := gabs.ParseJSON( []byte( command_result_string ) )
	// 	if command_json_parsed_error != nil { fmt.Println( command_json_parsed_error ); return }
	// 	fmt.Println( command_json_parsed )
	// } else{
	// 	fmt.Println( command_result )
	// }

	switch state_name {
		case "spotify_status":
			GenericSpotify()
		case "Spotify":
			GenericSpotify()
		default:
			fmt.Println( "No Active States" )
			return
	}
}

func main() {
	redis := redis.Manager{}
	redis.Connect( "localhost:6379" , 3 , "" )
	redis.Subscribe( "LOG.ALL" , log_all_handler )
}