package statelocaltvshowroutehandler

import (
	"fmt"
	try "github.com/manucorporat/try"
	fiber "github.com/gofiber/fiber/v2"
	utils "c2server/utils"
	types "c2server/types"
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
	result := "failed"
	var status types.VLCCommonStatus
	try.This( func() {
		status = localtvshow.Start()
		result = "success"
	}).Catch( func ( e try.E ) {
		fmt.Println( e )
	})
	logger.WithFields( logrus.Fields {
		"command": "route_local_tv_show_teardown_result" ,
		"route_local_tv_show_teardown_result": result ,
		"status": status ,
	}).Info( "Route === LocalTVShow === Start()" )
	return context.JSON( fiber.Map{
		"route": "/states/localtvshow/start" ,
		"result": "success" ,
		"status": status ,
	})
}

func Teardown( context *fiber.Ctx ) ( error ) {
	result := "failed"
	var status types.VLCCommonStatus
	try.This( func() {
		status = localtvshow.Teardown()
		result = "success"
	}).Catch( func ( e try.E ) {
		fmt.Println( e )
	})
	logger.WithFields( logrus.Fields {
		"command": "route_local_tv_show_teardown_result" ,
		"route_local_tv_show_teardown_result": result ,
	}).Info( "Route === LocalTVShow === Teardown()" )
	return context.JSON( fiber.Map{
		"route": "/states/localtvshow/teardown" ,
		"result": "success" ,
		"status": status ,
	})
}

func Play( context *fiber.Ctx ) ( error ) {
	result := "failed"
	var status types.VLCCommonStatus
	try.This( func() {
		status = localtvshow.Play()
		result = "success"
	}).Catch( func ( e try.E ) {
		fmt.Println( e )
	})
	logger.WithFields( logrus.Fields {
		"command": "route_local_tv_show_play_result" ,
		"route_local_tv_show_play_result": result ,
	}).Info( "Route === LocalTVShow === Play()")
	return context.JSON( fiber.Map{
		"route": "/states/localtvshow/play" ,
		"result": "success",
	})
}

func Stop( context *fiber.Ctx ) ( error ) {
	result := "failed"
	var status types.VLCCommonStatus
	try.This( func() {
		status = localtvshow.Stop()
		result = "success"
	}).Catch( func ( e try.E ) {
		fmt.Println( e )
	})
	logger.WithFields( logrus.Fields {
		"command": "route_local_tv_show_stop_result" ,
		"route_local_tv_show_stop_result": result ,
		"status": status ,
	}).Info( "Route === LocalTVShow === Stop()" )
	return context.JSON( fiber.Map{
		"route": "/states/localtvshow/stop" ,
		"result": "success" ,
		"status": status ,
	})
}

func Pause( context *fiber.Ctx ) ( error ) {
	result := "failed"
	var status types.VLCCommonStatus
	try.This( func() {
		status = localtvshow.Pause()
		result = "success"
	}).Catch( func ( e try.E ) {
		fmt.Println( e )
	})
	logger.WithFields( logrus.Fields {
		"command": "route_local_tv_show_pause_result" ,
		"route_local_tv_show_pause_result": result ,
		"status": status ,
	}).Info( "Route === LocalTVShow === Pause()" )
	return context.JSON( fiber.Map{
		"route": "/states/localtvshow/pause" ,
		"result": "success" ,
		"status": status ,
	})
}

func Resume( context *fiber.Ctx ) ( error ) {
	result := "failed"
	var status types.VLCCommonStatus
	try.This( func() {
		status = localtvshow.Resume()
		result = "success"
	}).Catch( func ( e try.E ) {
		fmt.Println( e )
	})
	logger.WithFields( logrus.Fields {
		"command": "route_local_tv_show_resume_result" ,
		"route_local_tv_show_resume_result": result ,
		"status": status ,
	}).Info( "Route === LocalTVShow === Resume()" )
	return context.JSON( fiber.Map{
		"route": "/states/localtvshow/resume" ,
		"result": "success" ,
		"status": status ,
	})
}

func Next( context *fiber.Ctx ) ( error ) {
	result := "failed"
	var status types.VLCCommonStatus
	try.This( func() {
		status = localtvshow.Next()
		result = "success"
	}).Catch( func ( e try.E ) {
		fmt.Println( e )
	})
	logger.WithFields( logrus.Fields {
		"command": "route_local_tv_show_next_result" ,
		"route_local_tv_show_next_result": result ,
		"status": status ,
	}).Info( "Route === LocalTVShow === Next()" )
	return context.JSON( fiber.Map{
		"route": "/states/localtvshow/next" ,
		"result": "success" ,
		"status": status ,
	})
}

func Previous( context *fiber.Ctx ) ( error ) {
	result := "failed"
	var status types.VLCCommonStatus
	try.This( func() {
		status = localtvshow.Previous()
		result = "success"
	}).Catch( func ( e try.E ) {
		fmt.Println( e )
	})
	logger.WithFields( logrus.Fields {
		"command": "route_local_tv_show_previous_result" ,
		"route_local_tv_show_previous_result": result ,
		"status": status ,
	}).Info( "Route === LocalTVShow === Next()" )
	return context.JSON( fiber.Map{
		"route": "/states/localtvshow/previous" ,
		"result": "success" ,
		"status": status ,
	})
}

func Status( context *fiber.Ctx ) ( error ) {
	fmt.Println( "/states/localtvshow/status" )
	result := "failed"
	var status types.VLCCommonStatus
	try.This( func() {
		status = localtvshow.Status()
		result = "success"
	}).Catch( func ( e try.E ) {
		fmt.Println( e )
	})
	logger.WithFields( logrus.Fields {
		"command": "route_local_tv_show_status_result" ,
		"route_local_tv_show_status_result": result ,
		"status": status ,
	}).Info( "Route === LocalTVShow === Next()" )
	return context.JSON( fiber.Map{
		"route": "/states/localtvshow/status" ,
		"result": result ,
		"status": status ,
	})
}