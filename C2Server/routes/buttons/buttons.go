package buttonsroutehandler

import (
	"fmt"
	//"reflect"
	logrus "github.com/sirupsen/logrus"
	try "github.com/manucorporat/try"
	async "github.com/rafaeldias/async"
	fiber "github.com/gofiber/fiber/v2"
	types "c2server/types"
	utils "c2server/utils"
	spotify "c2server/states/spotify"
	localtvshow "c2server/states/local_tv_show"
)

var logger *logrus.Entry = utils.BuildLogger( "Buttons" )


// Spotify Random Currated Playlist
func Button1( context *fiber.Ctx ) ( error ) {
	result := "failed"
	status := types.SpotifyStatus{}
	teardown_result := "failed"
	tv_preparation_result := "failed"
	try.This( func() {
		_ , async_error := async.Parallel( async.MapTasks{
			"teardown": func() {
				teardown_result = utils.TeardownCurrentState()
			} ,
			"tv_prep": func() {
				tv_preparation_result = utils.PrepareTV()
			} ,
		})
		if async_error == nil {
			status = spotify.Start()
			result = "success"
		}
	}).Catch( func ( e try.E ) {
		fmt.Println( e )
	})
	button_1_result := struct {
		Status types.SpotifyStatus
		Teardown string
		TVPrep string
		Result string
	} {
		Status: status ,
		Teardown: teardown_result ,
		TVPrep: tv_preparation_result ,
		Result: result ,
	}
	logger.WithFields( logrus.Fields {
		"command": "button_1_result" ,
		"button_1_result": button_1_result ,
	}).Info( "State === Button 1 === Status")
	return context.JSON( fiber.Map{
		"route": "/button/1" ,
		"previous_state_teardown_result": teardown_result ,
		"tv_preparation_result": tv_preparation_result ,
		"result": result ,
		"status": status ,
	})
}


// Local TV Show
func Button2( context *fiber.Ctx ) ( error ) {
	result := "failed"
	status := types.VLCCommonStatus{}
	teardown_result := "failed"
	tv_preparation_result := "failed"
	try.This( func() {
		_ , async_error := async.Parallel( async.MapTasks{
			"teardown": func() {
				teardown_result = utils.TeardownCurrentState()
			} ,
			"tv_prep": func() {
				tv_preparation_result = utils.PrepareTV()
			} ,
		})
		if async_error == nil {
			status = localtvshow.Start()
			result = "success"
		}
	}).Catch( func ( e try.E ) {
		fmt.Println( e )
	})
	button_2_result := struct {
		Status types.VLCCommonStatus
		Teardown string
		TVPrep string
		Result string
	} {
		Status: status ,
		Teardown: teardown_result ,
		TVPrep: tv_preparation_result ,
		Result: result ,
	}
	logger.WithFields( logrus.Fields {
		"command": "button_2_result" ,
		"button_2_result": button_2_result ,
	}).Info( "State === Button 2 === Status")
	return context.JSON( fiber.Map{
		"route": "/button/2" ,
		"previous_state_teardown_result": teardown_result ,
		"tv_preparation_result": tv_preparation_result ,
		"result": result ,
		"status": status ,
	})
}


// Generic Pause
func Button6( context *fiber.Ctx ) ( error ) {
	result := "failed"
	try.This( func() {
		result = utils.PauseCurrentState()
	}).Catch( func ( e try.E ) {
		fmt.Println( e )
	})
	button_6_result := struct {
		Result string
	} {
		Result: result ,
	}
	logger.WithFields( logrus.Fields {
		"command": "button_6_result" ,
		"button_6_result": button_6_result ,
	}).Info( "State === Button 6 === Result")
	return context.JSON( fiber.Map{
		"route": "/states/button/6" ,
		"result": result ,
	})
}

// Generic Previous
func Button7( context *fiber.Ctx ) ( error ) {
	result := "failed"
	try.This( func() {
		result = utils.PreviousCurrentState()
	}).Catch( func ( e try.E ) {
		fmt.Println( e )
	})
	button_7_result := struct {
		Result string
	} {
		Result: result ,
	}
	logger.WithFields( logrus.Fields {
		"command": "button_7_result" ,
		"button_7_result": button_7_result ,
	}).Info( "State === Button 7 === Result")
	return context.JSON( fiber.Map{
		"route": "/states/button/7" ,
		"result": result ,
	})
}

// Generic Stop
func Button8( context *fiber.Ctx ) ( error ) {
	result := "failed"
	try.This( func() {
		result = utils.StopCurrentState()
	}).Catch( func ( e try.E ) {
		fmt.Println( e )
	})
	button_8_result := struct {
		Result string
	} {
		Result: result ,
	}
	logger.WithFields( logrus.Fields {
		"command": "button_8_result" ,
		"button_8_result": button_8_result ,
	}).Info( "State === Button 8 === Result")
	return context.JSON( fiber.Map{
		"route": "/states/button/8" ,
		"result": result ,
	})
}

// Generic Next
func Button9( context *fiber.Ctx ) ( error ) {
	result := "failed"
	try.This( func() {
		result = utils.NextCurrentState()
	}).Catch( func ( e try.E ) {
		fmt.Println( e )
	})
	button_9_result := struct {
		Result string
	} {
		Result: result ,
	}
	logger.WithFields( logrus.Fields {
		"command": "button_9_result" ,
		"button_9_result": button_9_result ,
	}).Info( "State === Button 9 === Result")
	return context.JSON( fiber.Map{
		"route": "/states/button/9" ,
		"result": result ,
	})
}


func ButtonPlay( context *fiber.Ctx ) ( error ) {
	result := "failed"
	try.This( func() {
		result = utils.PlayCurrentState()
	}).Catch( func ( e try.E ) {
		fmt.Println( e )
	})
	button_play_result := struct {
		Result string
	} {
		Result: result ,
	}
	logger.WithFields( logrus.Fields {
		"command": "button_play_result" ,
		"button_play_result": button_play_result ,
	}).Info( "State === Play === Result")
	return context.JSON( fiber.Map{
		"route": "/play" ,
		"result": result ,
	})
}