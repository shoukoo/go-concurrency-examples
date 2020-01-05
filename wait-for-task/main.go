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

// waitForTask: You are a developer and you build a new robot. Your new
// robot doesn't know immediately what they are expected to do and waits for
// you to tell them what to do. You prepare the work and send it to them. The
// amount of time they wait is unknown because you need a guarantee that the
// work your sending is received by the robot.

func waitForTask() {
	ch := make(chan string)
	go func() {
		fmt.Println("Robot received a task")
		task := <-ch
		fmt.Printf("Robot finished the task %s\n", task)
	}()

	fmt.Println("Sending a task to the robot ")
	ch <- "103"
	time.Sleep(time.Second)
}
