package main

import (
	"fmt"
	"sync"
	"time"
)

func processItem(item int, wg *sync.WaitGroup, channel chan string) {
	defer wg.Done()
	time.Sleep(1000)
	channel <- fmt.Sprintf("The item I processed: %d", item)
}

func printProcessedItem(item string, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(1000)
	fmt.Println("The item coming back from channel ---", item)
}

func main() {
	items := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	channel := make(chan string, 10)
	var wg sync.WaitGroup
	wg.Add(len(items))
	for _, item := range items {
		go processItem(item, &wg, channel)
	}
	wg.Wait()
	close(channel)

	var wgTwo sync.WaitGroup
	for i := 1; i <= len(items); i++ {
		wgTwo.Add(1)
		go printProcessedItem(<-channel, &wgTwo)
	}
	wgTwo.Wait()
}
