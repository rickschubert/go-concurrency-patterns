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
	for stringifiedItem := range channel {
		wgTwo.Add(1)
		go printProcessedItem(stringifiedItem, &wgTwo)
	}
	wgTwo.Wait()
}
