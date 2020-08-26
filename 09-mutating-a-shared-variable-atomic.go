package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

type expense struct {
	outgoing bool
	money    int64
}

func performFakeNetworkRequest() {
	time.Sleep(time.Millisecond * 1300)
}

func accumulateExpenseToTotal(expense expense, total *int64, waitgroup *sync.WaitGroup) {
	defer waitgroup.Done()
	performFakeNetworkRequest()
	if expense.outgoing {
		negativeDelta := expense.money - expense.money - expense.money
		atomic.AddInt64(total, negativeDelta)
	} else {
		atomic.AddInt64(total, expense.money)
	}
}

func main() {
	start := time.Now()
	fmt.Println("Let's pretend I have a variable which tells me how much money I have. And I have a list of expenses, showing both incoming and outgoing money. I need to perform some network requests over each expense and after that I want to mutate the shared variable which tracks all the money I have. I am achieving this with the atomic package.\n")

	var waitgroup sync.WaitGroup
	var allMoneyIHave int64

	expenses := []expense{
		{
			outgoing: true,
			money:    1045,
		},
		{
			outgoing: false,
			money:    240032,
		},
		{
			outgoing: true,
			money:    2574,
		},
		{
			outgoing: true,
			money:    423,
		},
		{
			outgoing: true,
			money:    999,
		},
		{
			outgoing: false,
			money:    20,
		},
	}

	waitgroup.Add(len(expenses))
	for i := 0; i < len(expenses); i++ {
		go accumulateExpenseToTotal(expenses[i], &allMoneyIHave, &waitgroup)
	}
	waitgroup.Wait()

	fmt.Println(fmt.Sprintf("All the money I have is £%v\n", float64(allMoneyIHave)/100))

	elapsed := time.Since(start)
	fmt.Println(fmt.Sprintf("This program took %s", elapsed))
}
