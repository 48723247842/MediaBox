package spotify

import (
	"fmt"
	utils "c2server/utils"
	"encoding/json"
	//"reflect"
	types "c2server/types"
	redis "github.com/0187773933/RedisManagerUtils/manager"
	spotify_dbus "github.com/0187773933/SpotifyDBUSController/controller"
	logrus "github.com/sirupsen/logrus"
)

var logger *logrus.Entry = utils.BuildLogger( "State-Spotify" )

func Status() ( result types.SpotifyStatus ) {
	logger.Info( "State === Spotify === Status()" )
	spotify := spotify_dbus.Controller{}
	spotify.Connect()
	spotify.UpdateStatus()
	logger.WithFields( logrus.Fields{
		"command": "spotify_status" ,
		"spotify_status": spotify.Status ,
	}).Info( "State === Spotify === Status() === Spotify Status" )
	result = spotify.Status
	return
}

func PlaybackStatus() ( result string ) {
	logger.Info( "State === Spotify === PlaybackStatus()" )
	result = "failed"
	spotify := spotify_dbus.Controller{}
	spotify.Connect()
	result = spotify.PlaybackStatus()
	logger.WithFields( logrus.Fields{
		"command": "spotify_playback_status" ,
		"spotify_playback_status": result ,
	}).Info( "State === Spotify === PlaybackStatus () === Spotify Status" )
	return
}

func Stop() ( result types.SpotifyStatus ) {
	logger.Info( "State === Spotify === Stop()" )
	spotify := spotify_dbus.Controller{}
	spotify.Connect()
	spotify.Stop()
	fmt.Println( spotify.Status )
	logger.WithFields( logrus.Fields{
		"command": "spotify_status" ,
		"spotify_status": spotify.Status ,
	}).Info( "State === Spotify === Stop() === Spotify Status")
	result = spotify.Status
	return
}

func Play() ( result types.SpotifyStatus ) {
	logger.Info( "State === Spotify === Play()" )
	spotify := spotify_dbus.Controller{}
	spotify.Connect()
	spotify.Play()
	logger.WithFields( logrus.Fields{
		"command": "spotify_status" ,
		"spotify_status": spotify.Status ,
	}).Info( "State == Spotify === Play() === Spotify Status")
	result = spotify.Status
	return
}

func Pause() ( result types.SpotifyStatus ) {
	logger.Info( "State === Spotify === Pause()" )
	spotify := spotify_dbus.Controller{}
	spotify.Connect()
	spotify.Pause()
	logger.WithFields( logrus.Fields{
		"command": "spotify_status" ,
		"spotify_status": spotify.Status ,
	}).Info( "State === Spotify === StartNextInCircularListOfMiscGenrePlaylists() === Spotify Status")
	result = spotify.Status
	return
}

func Previous() ( result types.SpotifyStatus ) {
	logger.Info( "State === Spotify === Previous()" )
	spotify := spotify_dbus.Controller{}
	spotify.Connect()
	spotify.Previous()
	logger.WithFields( logrus.Fields{
		"command": "spotify_status" ,
		"spotify_status": spotify.Status ,
	}).Info( "State === Spotify === Previous() === Spotify Status")
	result = spotify.Status
	return
}

func Next() ( result types.SpotifyStatus ) {
	logger.Info( "State === Spotify === Next()" )
	spotify := spotify_dbus.Controller{}
	spotify.Connect()
	spotify.Next()
	logger.WithFields( logrus.Fields{
		"command": "spotify_status" ,
		"spotify_status": spotify.Status ,
	}).Info( "State === Spotify === Next() === Spotify Status" )
	result = spotify.Status
	return
}

func StartNextInCircularListOfMiscGenrePlaylists() ( result types.SpotifyStatus ) {
	logger.Info( "State === Spotify === StartNextInCircularListOfMiscGenrePlaylists()" )
	redis := redis.Manager{}
	redis.Connect( "localhost:6379" , 3 , "" )
	next_playlist_uri := redis.CircleNext( "CONFIG.SPOTIFY.PLAYLISTS.GENERES.MISC" )
	logger.WithFields( logrus.Fields{
		"command": "next_playlist_uri" ,
		"next_playlist_uri": next_playlist_uri ,
	}).Info( "State === Spotify === StartNextInCircularListOfMiscGenrePlaylists() === Next Playlist URI" )
	spotify := spotify_dbus.Controller{}
	spotify.Connect()
	spotify.OpenURI( next_playlist_uri )
	logger.WithFields( logrus.Fields{
		"command": "" ,
		"spotify_status": spotify.Status ,
	}).Info( "State === Spotify === StartNextInCircularListOfMiscGenrePlaylists() === Spotify Status" )
	result = spotify.Status
	return
}

func build_state_meta_data( state_name string ) ( json_string string ) {
	state_data := types.StateMetaData {
		Name: state_name ,
		GenericType: "Spotify" ,
		RestartOnFail: true ,
		NowPlaying: types.NowPlayingMeta{} ,
	}
	json_marshal_result , json_marshal_error := json.Marshal( state_data )
	if json_marshal_error != nil { panic( json_marshal_error ) }
	json_string = string( json_marshal_result )
	return
}

func swap_current_and_previous_state_info( state_name string ) {
	redis := redis.Manager{}
	redis.Connect( "localhost:6379" , 3 , "" )
	state_current := redis.Get( "STATE.CURRENT" )
	logger.WithFields( logrus.Fields{
		"command": "state_current" ,
		"state_current": state_current ,
	}).Info( "State === Spotify === swap_current_and_previous_state_info() === STATE CURRENT" )
	redis.Set( "STATE.PREVIOUS" , state_current )
	state_meta_data := build_state_meta_data( state_name )
	logger.WithFields( logrus.Fields{
		"command": "new_state" ,
		"new_state": state_meta_data ,
	}).Info( "State === Spotify === swap_current_and_previous_state_info() === NEW STATE" )
	redis.Set( "STATE.CURRENT" , state_meta_data )
}

func Start() ( result types.SpotifyStatus ) {
	logger.Info( "State === Spotify === Start()" )
	swap_current_and_previous_state_info( "SpotifyStartNextInCircularListOfMiscGenrePlaylists" )
	result = StartNextInCircularListOfMiscGenrePlaylists()
	return
}