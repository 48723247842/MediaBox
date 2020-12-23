package buttonsroutehandler

import (
	"fmt"
	"reflect"
	logrus "github.com/sirupsen/logrus"
	try "github.com/manucorporat/try"
	async "github.com/rafaeldias/async"
	fiber "github.com/gofiber/fiber/v2"
	types "c2server/types"
	utils "c2server/utils"
	spotify "c2server/states/spotify"
	//localtvshow "c2server/states/local_tv_show"
)

var logger *logrus.Entry = utils.BuildLogger( "Buttons" )

// Spotify Random Currated Playlist
func Button1( context *fiber.Ctx ) ( error ) {
	result := "failed"
	status := types.SpotifyStatus{}
	teardown_result := "failed"
	tv_preparation_result := "failed"
	try.This( func() {
		async_result , async_error := async.Parallel( async.MapTasks{
			"teardown": func() {
				teardown_result = utils.TeardownCurrentState()
			} ,
			"tv_prep": func() {
				tv_preparation_result = utils.PrepareTV()
			} ,
		})
		if async_error != nil { fmt.Println( async_error ) }
		fmt.Println( reflect.TypeOf( async_result ) )
		status = spotify.Start()
		result = "success"
	}).Catch( func ( e try.E ) {
		fmt.Println( e )
	})
	return context.JSON( fiber.Map{
		"route": "/states/spotify/start" ,
		"previous_state_teardown_result": teardown_result ,
		"tv_preparation_result": tv_preparation_result ,
		"result": result ,
		"status": status ,
	})
}