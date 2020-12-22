package localtvshow

import (
	"fmt"
	logrus "github.com/sirupsen/logrus"
	//redis "github.com/0187773933/RedisManagerUtils/manager"
	utils "c2server/utils"
	substates "c2server/states/local_tv_show/substates"
	vlc "github.com/0187773933/VLCWrapper/wrapper"
)

var logger *logrus.Entry = utils.BuildLogger( "LocalTVShow" )

func Stop() ( result string ) {
	logger.Info( "State === LocalTVShow === Stop()" )
	result = "failed"
	// logger.WithFields( logrus.Fields{
	// 	"command": "spotify_status" ,
	// 	"vlc_status": spotify.Status ,
	// }).Info( "State === LocalTVShow === VLC Status" )
	return
}

func Play() ( result string ) {
	logger.Info( "State === LocalTVShow === Play()" )
	result = "failed"
	p := vlc.NewPlayer( nil )
	p.PublicExecCommand( "play" )
	status , _ := p.PublicExecCommand( "status" )
	fmt.Println( status )
	// logger.WithFields( logrus.Fields{
	// 	"command": "spotify_status" ,
	// 	"spotify_status": spotify.Status ,
	// }).Info( "State === LocalTVShow === VLC Status" )
	return
}

func Pause() ( result string ) {
	logger.Info( "State === LocalTVShow === Pause()" )
	result = "failed"
	p := vlc.NewPlayer( nil )
	p.PublicExecCommand( "pause" )
	position , _ := p.Position()
	fmt.Println( position )
	// logger.WithFields( logrus.Fields{
	// 	"command": "spotify_status" ,
	// 	"spotify_status": spotify.Status ,
	// }).Info( "State === LocalTVShow === VLC Status" )
	return
}

func Previous() ( result string ) {
	logger.Info( "State === LocalTVShow === Previous()" )
	result = "failed"
	p := vlc.NewPlayer( nil )
	p.PublicExecCommand( "prev" )
	status , _ := p.PublicExecCommand( "status" )
	fmt.Println( status )
	// logger.WithFields( logrus.Fields{
	// 	"command": "spotify_status" ,
	// 	"spotify_status": spotify.Status ,
	// }).Info( "State === LocalTVShow === VLC Status" )
	return
}

func Next() ( result string ) {
	logger.Info( "State === LocalTVShow === Next()" )
	result = "failed"
	p := vlc.NewPlayer( nil )
	p.PublicExecCommand( "next" )
	status , _ := p.PublicExecCommand( "status" )
	fmt.Println( status )
	// logger.WithFields( logrus.Fields{
	// 	"command": "spotify_status" ,
	// 	"spotify_status": spotify.Status ,
	// }).Info( "State === LocalTVShow === VLC Status" )
	return
}


func Start() ( result string ) {
	logger.Info( "State === LocalTVShow === Start()" )
	substates.StartNextShowInCircularListAndNextEpisodeInCircularList()
	return
}