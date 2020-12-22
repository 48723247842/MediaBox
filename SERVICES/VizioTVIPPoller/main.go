package main

import (
	"os"
	"fmt"
	"context"
	localnetwork "github.com/0187773933/LocalNetworkScanner"
	"github.com/go-redis/redis/v8"
	// https://godoc.org/github.com/gomodule/redigo/redis#pkg-examples ????
)

var ctx = context.Background()
func main() {
	redis_client := redis.NewClient( &redis.Options{
		Addr: "localhost:6379" ,
		//Addr: "100.112.11.51:6379" ,
		Password: "" ,
		DB: 3 ,
	})
	mac_address , err := redis_client.Get( ctx , "CONFIG.VIZIO_TV.MAC_ADDRESS" ).Result()
	if err != nil { /*fmt.Println(err);*/ os.Exit(0) }
	fmt.Println( "CONFIG.VIZIO_TV.MAC_ADDRESS" , mac_address )
	interface_name , err := redis_client.Get( ctx , "CONFIG.INTERNET_INTERFACE_NAME" ).Result()
	if err != nil { /*fmt.Println(err);*/ os.Exit(0) }
	ip_address := localnetwork.GetIPAddressFromMacAddress( interface_name , mac_address )
	if len( ip_address ) < 1 { os.Exit(0) }
	err = redis_client.Set( ctx , "STATE.VIZIO_TV.IP_ADDRESS" , ip_address , 0 ).Err()
	if err != nil { os.Exit(0) }
	fmt.Println( ip_address )
}
