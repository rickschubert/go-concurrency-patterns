package main

import (
	"fmt"
	"sync"
)

func sendHobbyToChannel(channel chan<- string, hobbyNumber int) {
	channel <- fmt.Sprintf("Hobby nr. %d", hobbyNumber+1)
	// This doesn't work because the channel in this function is send-only:
	// fmt.Println(<- channel)
}

func readHobbiesFromChannel(channel <-chan string, waitgroup *sync.WaitGroup) {
	for {
		hobby := <-channel
		fmt.Println(hobby)
		waitgroup.Done()
		// This doens't work because the channel in this function is read-only:
		// sendHobbyToChannel(channel, 42)
		// This doesn't work either because the channel is read-only:
		// channel <- "Another hobby"
	}
}

func main() {
	fmt.Println("Let's pretend I need to fill up a channel with multiple values and another goroutine should then read from this channel. I can prevent the first goroutine from reading off of the channel and I can prevent the second goroutine from writing onto the channel to seperate out responsibilities.")

	hobbiesChannel := make(chan string)
	waitgroup := &sync.WaitGroup{}
	amountOfHobbiesToSend := 10
	waitgroup.Add(amountOfHobbiesToSend)

	for i := 0; i < amountOfHobbiesToSend; i++ {
		go sendHobbyToChannel(hobbiesChannel, i)
	}
	go readHobbiesFromChannel(hobbiesChannel, waitgroup)

	waitgroup.Wait()
	fmt.Println("Tada, all finished.")
}
