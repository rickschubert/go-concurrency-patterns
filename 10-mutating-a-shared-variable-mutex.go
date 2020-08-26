package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

func generateRandomString() string {
	rand.Seed(time.Now().UnixNano())
	return strconv.FormatInt(rand.Int63(), 10)
}

func getRandomStringNetworkRequest(finalString *string, waitgroup *sync.WaitGroup, mutex *sync.Mutex) {
	defer waitgroup.Done()
	time.Sleep(time.Second)
	// Comment out mutex.Lock and mutex.Unlock to test it for yourself - if you
	// leave these out, there will be race conditions
	mutex.Lock()
	*finalString = *finalString + "a"
	mutex.Unlock()
}

func main() {
	fmt.Println("Let's pretend I need to fire off lots of goroutines which all have to modify a shared variable, for example adding a character to a string. Since this shared variable is a string, I can't use the atomic package but have to use a mutex which will lock the variable down so that only one goroutine at a time is able to read and modify it.\n")

	var theFinalString string
	var waitgroup sync.WaitGroup
	var mutex sync.Mutex
	timesOfCharactersToAdd := 40

	waitgroup.Add(timesOfCharactersToAdd)
	for i := 0; i < timesOfCharactersToAdd; i++ {
		go getRandomStringNetworkRequest(&theFinalString, &waitgroup, &mutex)
	}
	waitgroup.Wait()

	if len(theFinalString) != timesOfCharactersToAdd {
		panic(fmt.Sprintf("Race condition: %d goroutines ran, all adding a character to the shared string variable, but the string only has %d characters.", timesOfCharactersToAdd, len(theFinalString)))
	} else {
		fmt.Println(fmt.Sprintf("The shared string variable has %d characters, as expected.", timesOfCharactersToAdd))
	}
}
