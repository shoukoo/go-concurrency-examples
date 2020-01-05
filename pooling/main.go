package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	pooling()
}

// pooling: You are a developer and you hire a team of robots. None of the new
// robots know what they are expected to do and wait for you to provide work.
// When work is provided to the group, any given robot can take it and you
// don't care who it is. The amount of time you wait for any given robot to
// take your work is unknown because you need a guarantee that the work your
// sending is received by an robot.

func pooling() {
	ch := make(chan string)
	robots := runtime.NumCPU()
	works := 1000

	for i := 0; i < robots; i++ {
		go func(id int) {
			for c := range ch {
				fmt.Printf(" robot %v received task %v\n", i, c)
				time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
				fmt.Printf(" robot %v finished task %v\n", i, c)
			}
		}(i)
	}

	for i := 0; i < works; i++ {
		ch <- fmt.Sprint(i)
	}

	time.Sleep(time.Second)
	close(ch)

}
