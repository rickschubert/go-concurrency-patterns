package main

import (
	"fmt"
	// "time"
)

func keepProgrammingRunningUntilUserInput() {
	var input string
	fmt.Scanln(&input)
}

func main() {
	channel := make(chan int, 2)

	// go func() {
	// 	for {
	// 		channel <- 42
	// 		channel <- 13
	// 		channel <- 87
	// 	}
	// }()

	channel <- 42
	channel <- 13
	// channel <- 87
	fmt.Println(<-channel)
	fmt.Println(<-channel)
	fmt.Println(<-channel)

	// go func() {
	// 	for {
	// 		time.Sleep(time.Millisecond * 250)
	// 		fmt.Println(<- channel)
	// 	}
	// }()

	keepProgrammingRunningUntilUserInput()
}
