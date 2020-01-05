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

// You are an engineer and you a build a robot. You pre-programed the new robot
// to know what needs to be done. You sit and waiting for the result from the robot.
// The amount of time you wait in unknown but you gurantee that the result sent by
// the robot is received by you
func waitForResult() {
	work := make(chan string)

	go func() {
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		work <- "Done"
	}()

	time.Sleep(time.Second)
	manager := <-work

	fmt.Printf("%+v\n", manager)
}
