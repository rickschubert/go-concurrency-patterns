package main

import (
	"fmt"
	"time"
)

func pingOntoChannel(c chan string) {
	for i := 0; ; i++ {
		time.Sleep(time.Millisecond * 500)
		c <- fmt.Sprintf("PING %d", i)
	}
}

func pongOntoChannel(c chan string) {
	for i := 0; ; i++ {
		time.Sleep(time.Millisecond * 500)
		c <- fmt.Sprintf("PONG %d", i)
	}
}

func printValuesFromChannel(c chan string) {
	for {
		msg := <-c
		fmt.Println(msg)
	}
}

func initiatePingAndListenForFinish() {
	var channel = make(chan string)
	go pingOntoChannel(channel)
	go pongOntoChannel(channel)
	go printValuesFromChannel(channel)
}

func main() {
	initiatePingAndListenForFinish()

	var input string
	fmt.Scanln(&input)
}
