package watcher

import (
	"os"
	"fmt"
	"strings"
	"net/http"
	"io/ioutil"
	try "github.com/manucorporat/try"
	evdev "github.com/gvalkov/golang-evdev"
)

func PostButtonEventToC2Server( button_number int ) {
	try.This( func() {
		url := fmt.Sprintf( "http://localhost:9363/button/%d" , button_number )
		fmt.Println( url )
		resp , err := http.Get( url )
		if err != nil { fmt.Println( err ) }
		defer resp.Body.Close()
		body , err := ioutil.ReadAll( resp.Body )
		fmt.Println( string( body ) )
	}).Catch( func ( e try.E ) {
		fmt.Println( e )
	})
}

func ListInputDevices() {
	devices, _ := evdev.ListInputDevices()
	fmt.Println( devices )
	for _ , dev := range devices {
		fmt.Println( dev.Fn )
		fmt.Println( dev.Name )
		fmt.Println( dev.Phys )
	}
}

// https://godoc.org/github.com/gvalkov/golang-evdev#InputDevice
func FindDragonRiseUSBGamePadEventPath() ( event_path string ) {
	event_path = ""
	devices, _ := evdev.ListInputDevices()
	for _ , dev := range devices {
		if strings.ContainsAny( dev.Fn , "DragonRise" ) {
			event_path = dev.Fn
			return
		}
	}
	return
}

func OpenInputDeviceNumber( event_number int ) ( device *evdev.InputDevice ) {
	var err error
	var device_open_result *evdev.InputDevice
	device_open_result , err = evdev.Open( fmt.Sprintf( "/dev/input/event%d" , event_number ) )
	if err != nil { panic( err ) }
	device = device_open_result
	return
}
func OpenInputDevicePath( input_device_path string ) ( device *evdev.InputDevice ) {
	var err error
	var device_open_result *evdev.InputDevice
	device_open_result , err = evdev.Open( input_device_path )
	if err != nil { panic( err ) }
	device = device_open_result
	return
}

type DragonRiseScanCodeButtonMapType map[uint16]int
var DragonRiseScanCodeButtonMap = DragonRiseScanCodeButtonMapType {
	297: 1 ,
	298: 2 ,
	299: 3 ,
	295: 4 ,
	296: 5 ,
	288: 6 ,
	289: 7 ,
	293: 8 ,
	294: 9 ,
	290: 10 ,
	291: 11 ,
	292: 12 ,
}

func Watch() {
	dragon_rise_event_path := FindDragonRiseUSBGamePadEventPath()
	if dragon_rise_event_path == "" { panic( "Couldn't Find DragonRise USB Gamepad" ) }
	fmt.Println( dragon_rise_event_path )
	usb_device := OpenInputDevicePath( dragon_rise_event_path )
	fmt.Println( usb_device )

	// input_event , _ := usb_device.ReadOne()
	// fmt.Println( input_event )

	// https://github.com/gvalkov/golang-evdev/issues/15
	for {
		events, err := usb_device.Read()
		if err != nil {
			fmt.Printf( "device.Read() Error: %v\n" , err )
			os.Exit( 1 )
		}
		for _ , ev := range events {
			if ev.Type != evdev.EV_KEY { continue }
			key := evdev.NewKeyEvent( &ev )
			if key.State != evdev.KeyDown { continue }
			//fmt.Println( key.Scancode )
			//fmt.Println( DragonRiseScanCodeButtonMap[ key.Scancode ] )
			_ , exists := DragonRiseScanCodeButtonMap[ key.Scancode ]
			if exists == false { continue }
			PostButtonEventToC2Server( DragonRiseScanCodeButtonMap[ key.Scancode ] )
		}
	}

}