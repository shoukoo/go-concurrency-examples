package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	waitForResult()
}

// waitForResult: You are a manager and you hire a new employee. Your new
// employee knows immediately what they are expected to do and starts their
// work. You sit waiting for the result of the employee's work. The amount
// of time you wait on the employee is unknown because you need a
// guarantee that the result sent by the employee is received by you.
func waitForResult() {
	task := make(chan string, 1)
	go func() {
		time.Sleep(time.Duration(rand.Intn(5000)) * time.Millisecond)
		task <- "done"
	}()

	manager := <-task

	fmt.Println(manager)
}
