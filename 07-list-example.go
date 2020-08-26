package main

import (
	"fmt"
	"sync"
	"time"
)

type hobby struct {
	id   int
	name string
}

type hobbyDetails struct {
	id        int
	name      string
	timeSpent string
	priority  string
}

func getHobbyDetails(hobby hobby, channel chan<- hobbyDetails) {
	channel <- getHobbyDetailsNetworkRequest(hobby)
}

func getHobbyDetailsNetworkRequest(hobby hobby) hobbyDetails {
	time.Sleep(time.Millisecond * 1000)
	return hobbyDetails{
		id:        hobby.id,
		name:      hobby.name,
		priority:  "High",
		timeSpent: "A lot",
	}
}

func getDetailedHobbiesWithGoRoutines(hobbies []hobby) []hobbyDetails {
	waitgroup := &sync.WaitGroup{}
	waitgroup.Add(len(hobbies))

	hobbyDetailsChannel := make(chan hobbyDetails)
	for _, hobby := range hobbies {
		go getHobbyDetails(hobby, hobbyDetailsChannel)
	}

	var detailedHobbies []hobbyDetails
	for i := 0; i < len(hobbies); i++ {
		detailedHobbies = append(detailedHobbies, <-hobbyDetailsChannel)
		waitgroup.Done()
	}
	waitgroup.Wait()
	return detailedHobbies
}

func getDetailedHobbiesSynchronously(hobbies []hobby) []hobbyDetails {
	var detailedHobbies []hobbyDetails
	for _, hobby := range hobbies {
		detailedHobbies = append(detailedHobbies, getHobbyDetailsNetworkRequest(hobby))
	}
	return detailedHobbies
}

func main() {
	start := time.Now()

	fmt.Println("Let's pretend I have a list of items and need to get some additional data for each of the items by making a network request.\n")
	hobbies := []hobby{
		{
			id:   1,
			name: "Programming",
		},
		{
			id:   2,
			name: "Writing",
		},
		{
			id:   3,
			name: "Theatre",
		},
		{
			id:   4,
			name: "Music",
		},
	}

	hobbiesWithDetails := getDetailedHobbiesWithGoRoutines(hobbies)
	// Comment out above line and use this line to see the difference in time
	// between fetching the details with goroutines and doing it synchronously
	// hobbiesWithDetails = getDetailedHobbiesSynchronously(person.hobbies)

	fmt.Println(fmt.Sprintf("Hobbies with addtional details:\n%+v\n", hobbiesWithDetails))

	elapsed := time.Since(start)
	fmt.Println(fmt.Sprintf("This program took took %s", elapsed))
}
