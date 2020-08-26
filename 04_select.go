package main

import (
	"time"
	"fmt"
)

func sendToChannelEveryTwoSeconds(channel chan string) {
	for {
		channel <- "from 1"
		time.Sleep(time.Second * 2)
	}
}

func sendToChannelEveryThreeSeconds(channel chan string) {
	for {
		channel <- "from 2"
		time.Sleep(time.Second * 3)
	}
}

func printValuesFromChannels(channelOne chan string, channelTwo chan string) {
	for {
		select {
		case msg1 := <- channelOne:
			fmt.Println(msg1)
		case msg2 := <- channelTwo:
			fmt.Println(msg2)
		case <- time.After(time.Second):
			fmt.Println("Waited 1 second. Nothing received on any of the channels.")
		// default:
		// 	fmt.Println("Nothing ready")
		}
	}
}

func keepProgrammingRunningUntilUserInput() {
	var input string
	fmt.Scanln(&input)
}

func main() {
	channelOne := make(chan string)
	channelTwo := make(chan string)

	go sendToChannelEveryTwoSeconds(channelOne)
	go sendToChannelEveryThreeSeconds(channelTwo)
	go printValuesFromChannels(channelOne, channelTwo)

	keepProgrammingRunningUntilUserInput()
  }
