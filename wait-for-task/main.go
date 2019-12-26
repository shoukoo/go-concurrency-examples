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
	waitForTask()
}

// waitForTask: You are a manager and you hire a new employee. Your new
// employee doesn't know immediately what they are expected to do and waits for
// you to tell them what to do. You prepare the work and send it to them. The
// amount of time they wait is unknown because you need a guarantee that the
// work your sending is received by the employee.
func waitForTask() {
	ch := make(chan string)

	go func() {
		fmt.Println("employee received a task")
		task := <-ch
		fmt.Println("Completed", task)
	}()

	time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
	ch <- "task12 "
	fmt.Println("Manager sent a task to employee")
	time.Sleep(time.Second)

}
