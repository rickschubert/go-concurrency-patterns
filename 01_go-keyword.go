package main

import (
  "fmt"
  "time"
  "math/rand"
)

func sleepRandomAmount() {
  amt := time.Duration(rand.Intn(250))
  time.Sleep(time.Millisecond * amt)
}

func looper(n int) {
  for i := 0; i < 10; i++ {
    fmt.Println(n, ":", i)
    sleepRandomAmount()
  }
}

func main() {
  for i := 0; i < 100; i++ {
    go looper(i)
  }
  var input string
  number, _ := fmt.Scanln(&input)
  fmt.Println(fmt.Sprintf("The user entered %v line(s).", number))
}
