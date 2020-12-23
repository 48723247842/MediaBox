package utils

import (
	pretty "github.com/gobs/pretty"
	"fmt"
	"strings"
	"reflect"
	"encoding/json"
	"encoding/base64"
	//"io"
	"io/ioutil"
	"net/http"
	log "github.com/sirupsen/logrus"
	types "c2server/types"
	redis "github.com/0187773933/RedisManagerUtils/manager"
	viziocontroller "github.com/0187773933/VizioController/controller"
)

func AddLogToRedis( input_struct *types.LoggerMain ) {
	//pretty.PrettyPrint( input_struct )
	redis := redis.Manager{}
	redis.Connect( "localhost:6379" , 3 , "" )
	json_marshal_result , json_marshal_error := json.Marshal( input_struct )
	if json_marshal_error != nil { panic( json_marshal_error ) }
	json_string := string( json_marshal_result )
	redis.ListPushRight( "LOG.ALL" , json_string )
	//var ctx = context.Background()
	//redis.Redis.Do( ctx , "PUBLISH" , "LOG.ALL" , json_string )
	redis.Publish( "LOG.ALL" , json_string )
	return
}

// func GetJSONStringFromRedis( redis_key string ) {
// 	json_get_test := redis.Get( "testmeta" )
// 	var json_get_test_struct TestStruct
// 	json_unmarshal_error := json.Unmarshal( []byte( json_get_test ) , &json_get_test_struct )
// 	if json_unmarshal_error != nil { panic( json_unmarshal_error ) }
// 	fmt.Println( json_get_test_struct )
// }

// https://godoc.org/github.com/sirupsen/logrus#Entry
// https://stackoverflow.com/a/54314594
type LoggerMain struct {}
type LoggerMainHook struct {}
func ( hook *LoggerMainHook ) Fire( entry *log.Entry ) error {
	time_stamp := fmt.Sprintf( "%d%s%d===%02d:%02d:%02d" ,
		entry.Time.Day() , strings.ToUpper( entry.Time.Month().String()[:3] ) , entry.Time.Year() ,
		entry.Time.Hour() , entry.Time.Minute() , entry.Time.Second() ,
	)
	new_log_line := types.LoggerMain{
		TimeStamp:  time_stamp ,
		NanosecondsSinceEpoch: entry.Time.UnixNano() ,
		Msg: entry.Message ,
		Author: entry.Data["author"].(string) ,
		Fields: entry.Data ,
		File: entry.Caller.File ,
		Function: entry.Caller.Function ,
		Line: entry.Caller.Line ,
		Level: entry.Level.String() ,
	}
	AddLogToRedis( &new_log_line )
	return nil
}
func ( hook *LoggerMainHook ) Levels() []log.Level {
	return []log.Level{
		log.PanicLevel,
		log.FatalLevel,
		log.ErrorLevel,
		log.WarnLevel,
		log.InfoLevel,
		log.DebugLevel,
	}
}


// func Write( input []byte ) ( n int , err error ) {
// 	fmt.Println( "here in custom writer" )
// 	fmt.Println( string( input ) )
// 	return
// }

// type CustomWriter interface {
// 	Write( input []byte ) ( n int , err error )
// }
// var writer CustomWriter

// func ( entry *log.Entry ) Writer() ( *io.PipeWriter ) {

// }

// "github.com/gobs/pretty"
func BuildLogger( author_name string ) ( logger *log.Entry ) {

	log.SetFormatter( &log.TextFormatter{
		//DisableColors: true,
		FullTimestamp: true ,
	})
	log.SetFormatter( &log.JSONFormatter{ DisableHTMLEscape: true } )
	log.SetReportCaller( true )


	logger_main_hook := LoggerMainHook{}
	log.AddHook( &logger_main_hook )

	logger = log.WithFields( log.Fields{
		"author": author_name ,
	})

	// https://github.com/sirupsen/logrus/blob/cd4bf4ef8de16b243cce0e062742feb34128648b/entry.go
	// logger.Entry = writer

	return
}


func Base64Encode( source string ) ( encoded string ) {
	encoded = base64.StdEncoding.EncodeToString( []byte( source ) )
	return
}

func Base64Decode( source string ) ( decoded string ) {
	decoded_bytes , _ := base64.StdEncoding.DecodeString( source )
	decoded = string( decoded_bytes[:] )
	return
}


func PrepareTV() ( result string ) {
	result = "failed"
	redis := redis.Manager{}
	redis.Connect( "localhost:6379" , 3 , "" )
	ip_address := redis.Get( "STATE.VIZIO_TV.IP_ADDRESS" )
	if ip_address == "" { return }
	auth_token := redis.Get( "STATE.VIZIO_TV.AUTH_TOKEN" )
	if auth_token == "" { return }

	current_power_state := viziocontroller.GetPowerState( ip_address , auth_token )
	fmt.Println( current_power_state )
	if current_power_state == 0 {
		viziocontroller.PowerOn( ip_address , auth_token )
	}
	current_volume := viziocontroller.GetVolume( ip_address , auth_token )
	fmt.Println( current_volume )
	if current_volume < 12 {
		viziocontroller.SetSettingsOption( ip_address , auth_token , "audio" , "volume" , 12 )
	}
	current_input := viziocontroller.GetCurrentInput( ip_address , auth_token )
	fmt.Println( current_input.Name )
	if current_input.Name != "hdmi1" {
		viziocontroller.SetInput( ip_address , auth_token , "HDMI-1" )
	}

	mute_value :=  viziocontroller.GetSetting( ip_address , auth_token , "audio" , "mute" )
	mute_value_string := mute_value.ITEMS[0].VALUE.(string)
	if mute_value_string != "Off" {
		viziocontroller.SetSettingsOption( ip_address , auth_token , "audio" , "mute" , "Off" )
	}
	result = "success"
}

func TeardownCurrentState() ( result string ) {
	result = "failed"
	redis := redis.Manager{}
	redis.Connect( "localhost:6379" , 3 , "" )
	state_current := redis.Get( "STATE.CURRENT" )
	var current_state types.StateMetaData
	json_unmarshal_error := json.Unmarshal( []byte( state_current ) , &current_state )
	if json_unmarshal_error != nil {
		fmt.Println( json_unmarshal_error )
	}
	var url string
	type_lowercase := strings.ToLower( current_state.GenericType )
	fmt.Println( type_lowercase )
	switch current_state.GenericType {
		case "Spotify":
			url = fmt.Sprintf( "http://localhost:9363/states/%s/teardown" , type_lowercase )
		case "LocalTVShow":
			url = fmt.Sprintf( "http://localhost:9363/states/%s/teardown" , type_lowercase )
		default:
			fmt.Println( "Unknown Current State" )
			fmt.Println( current_state.GenericType )
	}
	response , err := http.Get( url )
	if err != nil { fmt.Println( err ) }
	defer response.Body.Close()
	body , err := ioutil.ReadAll( response.Body )
	if err != nil { fmt.Println( err ) }
	result = string( body )
	return
}