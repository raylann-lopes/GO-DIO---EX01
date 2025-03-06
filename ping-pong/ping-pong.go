package main

import (
	"fmt"
	"time"
)

func ping(ping_pong chan string){
	for ping := 0; ping<100; ping++{
		ping_pong <- "ping"

	}
}

func pong(ping_pong chan string){
	for pong := 0; pong<100; pong++{
		ping_pong <- "pong"
	}
}

func timer(ping_pong chan string){
	for{
	msg := <-ping_pong
	fmt.Println(msg)
	time.Sleep(time.Second * 2)
	}
}
func main () {

	ping_pong := make(chan string)

	go ping(ping_pong)
	go pong(ping_pong)	
	go timer(ping_pong)

	var write string
	fmt.Scanln(&write)
}