package statelocaltvshowroutehandler

import (
	"fmt"
	try "github.com/manucorporat/try"
	fiber "github.com/gofiber/fiber/v2"
	utils "c2server/utils"
	logrus "github.com/sirupsen/logrus"
	localtvshow "c2server/states/local_tv_show"
)

var logger *logrus.Entry = utils.BuildLogger( "State-LocalTVShow" )

// localtvshow Test
// localtvshow.Start()
// time.Sleep( 1 * time.Second )
// localtvshow.Next()
// time.Sleep( 1 * time.Second )

func Start( context *fiber.Ctx ) ( error ) {
	fmt.Println( "/states/localtvshow/start" )
	result := localtvshow.Start()
	fmt.Println( result )
	return context.JSON( fiber.Map{
		"route": "/states/localtvshow/start" ,
		"result": "success",
	})
}

func Teardown( context *fiber.Ctx ) ( error ) {
	fmt.Println( "/states/localtvshow/teardown" )
	return context.JSON( fiber.Map{
		"route": "/states/localtvshow/teardown" ,
		"result": "success",
	})
}

func Play( context *fiber.Ctx ) ( error ) {
	fmt.Println( "/states/localtvshow/play" )
	return context.JSON( fiber.Map{
		"route": "/states/localtvshow/play" ,
		"result": "success",
	})
}

func Stop( context *fiber.Ctx ) ( error ) {
	fmt.Println( "/states/localtvshow/stop" )
	return context.JSON( fiber.Map{
		"route": "/states/localtvshow/stop" ,
		"result": "success",
	})
}

func Pause( context *fiber.Ctx ) ( error ) {
	fmt.Println( "/states/localtvshow/pause" )
	return context.JSON( fiber.Map{
		"route": "/states/localtvshow/pause" ,
		"result": "success",
	})
}

func Resume( context *fiber.Ctx ) ( error ) {
	fmt.Println( "/states/localtvshow/resume" )
	return context.JSON( fiber.Map{
		"route": "/states/localtvshow/resume" ,
		"result": "success",
	})
}

func Next( context *fiber.Ctx ) ( error ) {
	fmt.Println( "/states/localtvshow/next" )
	return context.JSON( fiber.Map{
		"route": "/states/localtvshow/next" ,
		"result": "success",
	})
}

func Previous( context *fiber.Ctx ) ( error ) {
	fmt.Println( "/states/localtvshow/previous" )
	return context.JSON( fiber.Map{
		"route": "/states/localtvshow/previous" ,
		"result": "success",
	})
}

func Status( context *fiber.Ctx ) ( error ) {
	fmt.Println( "/states/localtvshow/status" )
	result := fiber.Map{
		"route": "/states/localtvshow/status" ,
		"result": "failed" ,
	}
	try.This( func() {
		status := localtvshow.Status()
		result["status"] = status
		result["result"] = "success"
	}).Catch( func ( e try.E ) {
		fmt.Println( e )
	})
	return context.JSON( result )
}