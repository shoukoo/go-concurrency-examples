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
	drop()
}

// drop: You are a manager and you hire a new employee. Your new employee
// doesn't know immediately what they are expected to do and waits for
// you to tell them what to do. You prepare the work and send it to them. The
// amount of time they wait is unknown because you need a guarantee that the
// work your sending is received by the employee. You won't wait for the
// employee to take the work if they are not ready to receive it. In that case
// you drop the work on the floor and try again with the next piece of work.
/**
task is droped and it's gone forever.
**/
func drop() {
	const ca = 100
	ch := make(chan string, ca)

	go func() {
		for c := range ch {
			fmt.Println("emplpyee received a signal ", c)
		}
	}()

	const works = 2000
	for i := 0; i < works; i++ {
		select {
		case ch <- "paper":
			fmt.Println("mager sent a signal ", i)
		default:
			fmt.Println("xxxxxxxxxxxxDrop the signal", i)
		}
	}
	time.Sleep(time.Second)
	close(ch)

}
