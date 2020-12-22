package statespotifyroutehandler

import (
	"fmt"
	fiber "github.com/gofiber/fiber/v2"
	utils "c2server/utils"
	logrus "github.com/sirupsen/logrus"
	spotify "c2server/states/spotify"
)

var logger *logrus.Entry = utils.BuildLogger( "State-Spotify" )

// Spotify Test
// spotify.Start()
// time.Sleep( 1 * time.Second )
// spotify.Next()
// time.Sleep( 1 * time.Second )

func Start( context *fiber.Ctx ) ( error ) {
	fmt.Println( "/states/spotify/start" )
	result := spotify.Start()
	fmt.Println( result )
	return context.JSON( fiber.Map{
		"route": "/states/spotify/start" ,
		"result": "success",
	})
}

func Teardown( context *fiber.Ctx ) ( error ) {
	fmt.Println( "/states/spotify/teardown" )
	return context.JSON( fiber.Map{
		"route": "/states/spotify/teardown" ,
		"result": "success",
	})
}

func Play( context *fiber.Ctx ) ( error ) {
	fmt.Println( "/states/spotify/play" )
	return context.JSON( fiber.Map{
		"route": "/states/spotify/play" ,
		"result": "success",
	})
}

func Stop( context *fiber.Ctx ) ( error ) {
	fmt.Println( "/states/spotify/stop" )
	return context.JSON( fiber.Map{
		"route": "/states/spotify/stop" ,
		"result": "success",
	})
}

func Pause( context *fiber.Ctx ) ( error ) {
	fmt.Println( "/states/spotify/pause" )
	return context.JSON( fiber.Map{
		"route": "/states/spotify/pause" ,
		"result": "success",
	})
}

func Resume( context *fiber.Ctx ) ( error ) {
	fmt.Println( "/states/spotify/resume" )
	return context.JSON( fiber.Map{
		"route": "/states/spotify/resume" ,
		"result": "success",
	})
}

func Next( context *fiber.Ctx ) ( error ) {
	fmt.Println( "/states/spotify/next" )
	return context.JSON( fiber.Map{
		"route": "/states/spotify/next" ,
		"result": "success",
	})
}

func Previous( context *fiber.Ctx ) ( error ) {
	fmt.Println( "/states/spotify/previous" )
	return context.JSON( fiber.Map{
		"route": "/states/spotify/previous" ,
		"result": "success",
	})
}

func Status( context *fiber.Ctx ) ( error ) {
	fmt.Println( "/states/spotify/status" )
	return context.JSON( fiber.Map{
		"route": "/states/spotify/status" ,
		"result": "success",
	})
}