package main

import (
	"fmt"
	"time"
)

type person struct {
	name    string
	age     int
	address string
	hobbies []string
}

func putAddressOntoChannel(channel chan string) {
	channel <- getAddressNetworkRequest()
}

func getAddressNetworkRequest() string {
	time.Sleep(time.Millisecond * 1000)
	return "LaLaLand"
}

func putAgeOntoChannel(channel chan int) {
	channel <- getAgeNetworkRequest()
}

func getAgeNetworkRequest() int {
	time.Sleep(time.Millisecond * 400)
	return 132
}

func putNameOntoChannel(channel chan string) {
	channel <- getNameNetworkRequest()
}

func getNameNetworkRequest() string {
	time.Sleep(time.Millisecond * 1000)
	return "Harold"
}

func putHobbiesOntoChannel(channel chan []string) {
	channel <- getHobbiesNetworkRequest()
}

func getHobbiesNetworkRequest() []string {
	time.Sleep(time.Millisecond * 900)
	return []string{"Climbing", "Cooking"}
}

func main() {
	start := time.Now()

	fmt.Println("Let's pretend I need to make 4 different API calls to retrieve information about a person. Each request takes a while.\n")
	var p person
	addressChannel := make(chan string)
	ageChannel := make(chan int)
	nameChannel := make(chan string)
	hobbiesChannel := make(chan []string)
	go putAddressOntoChannel(addressChannel)
	go putAgeOntoChannel(ageChannel)
	go putNameOntoChannel(nameChannel)
	go putHobbiesOntoChannel(hobbiesChannel)

	// Using concurrency with channels, retrieving all the values takes around 1 second
	p.address = <-addressChannel
	p.age = <-ageChannel
	p.name = <-nameChannel
	p.hobbies = <-hobbiesChannel

	// Without concurrency, it takes around 3.3 seconds to get the values sequentially.
	// Comment these lines in and above out to see how it would behave without concurrency:
	// p.address = getAddressNetworkRequest()
	// p.age = getAgeNetworkRequest()
	// p.name = getNameNetworkRequest()
	// p.hobbies = getHobbiesNetworkRequest()

	fmt.Println(fmt.Sprintf("All person attributes put together:\n%+v\n", p))

	elapsed := time.Since(start)
	fmt.Println(fmt.Sprintf("This program took took %s", elapsed))
}
