package main

import (
	"time"
	utils "c2server/utils"
	//spotify "c2server/states/spotify"
	local_tv_show "c2server/states/local_tv_show"
)


func Test() {
	logger := utils.BuildLogger( "C2Server" )
	logger.Info( "main()" )

	// Spotify Test
	// spotify.Start()
	// time.Sleep( 1 * time.Second )
	// spotify.Next()
	// time.Sleep( 1 * time.Second )

	// Local TV Show Test
	local_tv_show.Start()
	time.Sleep( 3 * time.Second )
	local_tv_show.Pause()
	time.Sleep( 3 * time.Second )
	local_tv_show.Play()

}