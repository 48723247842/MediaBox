package statespotifyroutehandler

import (
	"fmt"
	types "c2server/types"
	try "github.com/manucorporat/try"
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
	result := "failed"
	status := types.SpotifyStatus{}
	try.This( func() {
		status = spotify.Start()
		result = "success"
	}).Catch( func ( e try.E ) {
		fmt.Println( e )
	})
	return context.JSON(fiber.Map{
		"route": "/states/spotify/start" ,
		"result": result ,
		"status": status ,
	})
}

func Teardown( context *fiber.Ctx ) ( error ) {
	result := "failed"
	status := types.SpotifyStatus{}
	try.This( func() {
		status = spotify.Stop()
		result = "success"
	}).Catch( func ( e try.E ) {
		fmt.Println( e )
	})
	return context.JSON(fiber.Map{
		"route": "/states/spotify/teardown" ,
		"result": result ,
		"status": status,
	})
}

func Play( context *fiber.Ctx ) ( error ) {
	result := "failed"
	status := types.SpotifyStatus{}
	try.This( func() {
		status = spotify.Play()
		result = "success"
	}).Catch( func ( e try.E ) {
		fmt.Println( e )
	})
	return context.JSON(fiber.Map{
		"route": "/states/spotify/teardown" ,
		"result": result ,
		"status": status,
	})
}

func Stop( context *fiber.Ctx ) ( error ) {
	result := "failed"
	status := types.SpotifyStatus{}
	try.This( func() {
		status = spotify.Stop()
		result = "success"
	}).Catch( func ( e try.E ) {
		fmt.Println( e )
	})
	return context.JSON(fiber.Map{
		"route": "/states/spotify/stop" ,
		"result": result ,
		"status": status ,
	})
}

func Pause( context *fiber.Ctx ) ( error ) {
	result := "failed"
	status := types.SpotifyStatus{}
	try.This( func() {
		status = spotify.Pause()
		result = "success"
	}).Catch( func ( e try.E ) {
		fmt.Println( e )
	})
	return context.JSON(fiber.Map{
		"route": "/states/spotify/pause" ,
		"result": result ,
		"status": status ,
	})
}

func Resume( context *fiber.Ctx ) ( error ) {
	result := "failed"
	status := types.SpotifyStatus{}
	try.This( func() {
		status = spotify.Play()
		result = "success"
	}).Catch( func ( e try.E ) {
		fmt.Println( e )
	})
	return context.JSON(fiber.Map{
		"route": "/states/spotify/resume" ,
		"result": result ,
		"status": status ,
	})
}

func Next( context *fiber.Ctx ) ( error ) {
	result := "failed"
	status := types.SpotifyStatus{}
	try.This( func() {
		status = spotify.Next()
		result = "success"
	}).Catch( func ( e try.E ) {
		fmt.Println( e )
	})
	return context.JSON(fiber.Map{
		"route": "/states/spotify/next" ,
		"result": result ,
		"status": status ,
	})
}

func Previous( context *fiber.Ctx ) ( error ) {
	result := "failed"
	status := types.SpotifyStatus{}
	try.This( func() {
		status = spotify.Previous()
		result = "success"
	}).Catch( func ( e try.E ) {
		fmt.Println( e )
	})
	return context.JSON(fiber.Map{
		"route": "/states/spotify/previous" ,
		"result": result ,
		"status": status ,
	})
}

func Status( context *fiber.Ctx ) ( error ) {
	result := "failed"
	status := types.SpotifyStatus{}
	try.This( func() {
		status = spotify.Status()
		fmt.Println( status )
		result = "success"
	}).Catch( func ( e try.E ) {
		fmt.Println( e )
	})
	return context.JSON(fiber.Map{
		"route": "/states/spotify/stop" ,
		"status": status ,
		"result": result ,
	})
}