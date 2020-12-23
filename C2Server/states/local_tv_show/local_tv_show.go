package localtvshow

import (
	"fmt"
	"strings"
	"strconv"
	logrus "github.com/sirupsen/logrus"
	//redis "github.com/0187773933/RedisManagerUtils/manager"
	utils "c2server/utils"
	types "c2server/types"
	substates "c2server/states/local_tv_show/substates"
	vlc_wrapper "github.com/0187773933/VLCTelnetWrapper/vlc"
)

var logger *logrus.Entry = utils.BuildLogger( "LocalTVShow" )

func get_vlc_client() ( vlc vlc_wrapper.Wrapper ) {
	vlc = vlc_wrapper.Wrapper{}
	vlc.Connect( "127.0.0.1:4212" )
	return
}

func get_status_object( vlc *vlc_wrapper.Wrapper ) ( status_object types.VLCStatus ) {
	status_string := vlc.Status()
	lines := strings.Split( status_string , "\r\n" )
	if len( lines ) == 3 { // There IS an Input File

		input_lines := strings.Split( lines[0] , "new input: " )

		// Input
		status_object.Input = strings.Split( input_lines[1] , " )" )[ 0 ]

		// Audio Volume
		audio_volume_lines := strings.Split( lines[ 1 ] , "audio volume: " )
		if len( audio_volume_lines ) < 1 { return }
		volume_string := strings.Split( audio_volume_lines[1] , " )" )[ 0 ]
		if volume_string != "" {
			volume_conversion_result , err := strconv.Atoi( volume_string )
			if err == nil {
				status_object.AudioVolume = volume_conversion_result
			}
		}

		// State
		state_lines := strings.Split( lines[ 2 ] , "state " )
		if len( state_lines ) < 2 { return }
		status_object.State = strings.Split( state_lines[1] , " )" )[ 0 ]

	} else if len( lines ) == 2 { // There Is NOT an Input File

		// Audio Volume
		audio_volume_lines := strings.Split( lines[ 0 ] , "audio volume: " )
		if len( audio_volume_lines ) >= 2 {
			volume_string := strings.Split( audio_volume_lines[ 1 ] , " )" )[ 0 ]
			if volume_string != "" {
				volume_conversion_result , err := strconv.Atoi( volume_string )
				if err == nil {
					status_object.AudioVolume = volume_conversion_result
				}
			}

		}

		// State
		state_lines := strings.Split( lines[ 1 ] , "state " )
		if len( state_lines ) < 2 { return }
		status_object.State = strings.Split( state_lines[1] , " )" )[ 0 ]

	} else if len( lines ) == 1 { // We Don't Know Why We Are Here
		fmt.Println( "1 Line" )
	}
	return
}

func get_info_object( vlc *vlc_wrapper.Wrapper ) ( info_object string ) {
	info_object = vlc.Info()
	//info_string := vlc.Info()
	// fmt.Println( info_string )
	// lines := strings.Split( info_string , "\n" )
	// if len( lines ) < 1 {
	// 	return
	// } else if len( lines ) == 2 {
	// 	fmt.Println( lines )
	// } else if len( lines ) == 1 {
	// 	fmt.Println( lines )
	// 	return
	// }
	return
}

func get_seconds( vlc *vlc_wrapper.Wrapper ) ( seconds int ) {
	seconds = 0
	seconds_string := vlc.GetTime()
	seconds_lines := strings.Split( seconds_string , "\n" )
	if len( seconds_lines ) < 1 { return }
	remove_space := strings.Split( seconds_lines[ 0 ] , " " )
	if len( remove_space ) < 2 { return }
	remove_slash_r := strings.Split( remove_space[ 1 ] , "\r" )
	if len( remove_slash_r ) < 1 { return }
	if remove_slash_r[ 0 ] == "" { return }
	conversion_result , err := strconv.Atoi( remove_slash_r[ 0 ] )
	if err != nil { fmt.Println( err ); return }
	seconds = conversion_result
	return
}

func get_length( vlc *vlc_wrapper.Wrapper ) ( length int ) {
	length = 0
	length_string := vlc.GetTime()
	length_lines := strings.Split( length_string , "\n" )
	if len( length_lines ) < 1 { return }
	remove_space := strings.Split( length_lines[ 0 ] , " " )
	if len( remove_space ) < 2 { return }
	remove_slash_r := strings.Split( remove_space[ 1 ] , "\r" )
	if len( remove_slash_r ) < 1 { return }
	if remove_slash_r[ 0 ] == "" { return }
	conversion_result , err := strconv.Atoi( remove_slash_r[ 0 ] )
	if err != nil { fmt.Println( err ); return }
	length = conversion_result
	return
}

func get_vlc_common_status( vlc *vlc_wrapper.Wrapper ) ( status types.VLCCommonStatus ) {
	status_object := get_status_object( vlc )
	status.Input = status_object.Input
	status.AudioVolume = status_object.AudioVolume
	status.State = status_object.State
	status.Info = get_info_object( vlc )
	status.Seconds = get_seconds( vlc )
	status.Length = get_length( vlc )
	return
}

func Stop() ( result types.VLCCommonStatus ) {
	logger.Info( "State === LocalTVShow === Stop()" )
	vlc := get_vlc_client()
	vlc.Stop()
	result = get_vlc_common_status( &vlc )
	// logger.WithFields( logrus.Fields{
	// 	"command": "spotify_status" ,
	// 	"vlc_status": spotify.Status ,
	// }).Info( "State === LocalTVShow === VLC Status" )
	return
}

func Play() ( result types.VLCCommonStatus ) {
	logger.Info( "State === LocalTVShow === Play()" )
	vlc := get_vlc_client()
	vlc.Play()
	result = get_vlc_common_status( &vlc )
	// logger.WithFields( logrus.Fields{
	// 	"command": "spotify_status" ,
	// 	"spotify_status": spotify.Status ,
	// }).Info( "State === LocalTVShow === VLC Status" )
	return
}

func Pause() ( result types.VLCCommonStatus ) {
	logger.Info( "State === LocalTVShow === Pause()" )
	vlc := get_vlc_client()
	vlc.Pause()
	result = get_vlc_common_status( &vlc )
	// logger.WithFields( logrus.Fields{
	// 	"command": "spotify_status" ,
	// 	"spotify_status": spotify.Status ,
	// }).Info( "State === LocalTVShow === VLC Status" )
	return
}

func Previous() ( result types.VLCCommonStatus ) {
	logger.Info( "State === LocalTVShow === Previous()" )
	vlc := get_vlc_client()
	vlc.Previous()
	result = get_vlc_common_status( &vlc )
	// logger.WithFields( logrus.Fields{
	// 	"command": "spotify_status" ,
	// 	"spotify_status": spotify.Status ,
	// }).Info( "State === LocalTVShow === VLC Status" )
	return
}

func Next() ( result types.VLCCommonStatus ) {
	logger.Info( "State === LocalTVShow === Next()" )
	vlc := get_vlc_client()
	vlc.Next()
	result = get_vlc_common_status( &vlc )
	// logger.WithFields( logrus.Fields{
	// 	"command": "spotify_status" ,
	// 	"spotify_status": spotify.Status ,
	// }).Info( "State === LocalTVShow === VLC Status" )
	return
}

func Status() ( result types.VLCCommonStatus ) {
	logger.Info( "State === LocalTVShow === Status()" )
	vlc := get_vlc_client()
	fmt.Println( vlc.Add( "/media/morphs/14TB/MEDIA_MANAGER/TVShows/DrakeAndJosh/001 - Drake and Josh - S01E01 - Pilot.mp4" ) )
	result = get_vlc_common_status( &vlc )
	logger.WithFields( logrus.Fields{
		"command": "local_tv_show_status" ,
		"local_tv_show_status": result ,
	}).Info( "State === LocalTVShow === Status" )
	return
}

func Start() ( result string ) {
	logger.Info( "State === LocalTVShow === Start()" )
	substates.StartNextShowInCircularListAndNextEpisodeInCircularList()
	return
}

func Teardown() ( result types.VLCCommonStatus ) {
	logger.Info( "State === LocalTVShow === Teardown()" )
	vlc := get_vlc_client()
	vlc.Stop()
	result = get_vlc_common_status( &vlc )
	// Could Run exec( "sudo pkill -9 vlc" ) and then let systemd restart
	// logger.WithFields( logrus.Fields{
	// 	"command": "spotify_status" ,
	// 	"spotify_status": spotify.Status ,
	// }).Info( "State === LocalTVShow === VLC Status" )
	return
}
