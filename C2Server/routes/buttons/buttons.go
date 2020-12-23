package buttonsroutehandler

import (
	"fmt"
	logrus "github.com/sirupsen/logrus"
	try "github.com/manucorporat/try"
	fiber "github.com/gofiber/fiber/v2"
	types "c2server/types"
	utils "c2server/utils"
	spotify "c2server/states/spotify"
	//localtvshow "c2server/states/local_tv_show"
)

var logger *logrus.Entry = utils.BuildLogger( "Buttons" )

// Spotify Random Currated Playlist
func Button1( context *fiber.Ctx ) ( error ) {

	teardown_result := utils.TeardownCurrentState()
	tv_preparation_result := utils.PrepareTV()

	result := "failed"
	status := types.SpotifyStatus{}
	try.This( func() {
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