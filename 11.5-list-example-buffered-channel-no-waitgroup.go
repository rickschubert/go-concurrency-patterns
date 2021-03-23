package main

import (
	"fmt"
	"time"
)

func processItem(item int, channel chan string) {
	time.Sleep(1000)
	channel <- fmt.Sprintf("The item I processed: %d", item)
}

func printProcessedItem(channel chan string) {
	time.Sleep(1000)
	fmt.Println("The item coming back from channel ---", <-channel)
}

func main() {
	items := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	channel := make(chan string, 10)
	for _, item := range items {
		go processItem(item, channel)
	}

	for i := 0; i < len(items); i++ {
		printProcessedItem(channel)
	}
}
