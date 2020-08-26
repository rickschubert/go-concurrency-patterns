package main

import "fmt"

func putOntoChannel(channel chan<- int) {
	for i := 0; i < 10; i++ {
		channel <- i
	}
}

func readFromChannel(channel <-chan int) {
	for {
		number := <- channel
		fmt.Println(number)
	}
}

func functionWithGoRoutines() {
	sendOnlyChannel := make(chan<- int)
	readOnlyChannel := make(<-chan int)
	go putOntoChannel(sendOnlyChannel)
	go readFromChannel(readOnlyChannel)
}

func main() {
	functionWithGoRoutines()

	var input string
	fmt.Scanln(&input)
}
