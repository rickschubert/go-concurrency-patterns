package main

import (
	"fmt"
	"time"
)

func employee(id string, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Println("worker", id, "started job", job)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", job)
		results <- job * 2
	}
}

func main() {
	numJobs := 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// Schedule 5 jobs
	jobs <- 1
	jobs <- 2
	jobs <- 3
	jobs <- 4
	jobs <- 5
	// (Optionally) close channel
	close(jobs)

	go employee("Jeff", jobs, results)
	go employee("Bob", jobs, results)
	go employee("Thomas", jobs, results)

	// Wait to receive finished indicators for all jobs
	for a := 1; a <= numJobs; a++ {
		fmt.Println(<-results)
	}
}
