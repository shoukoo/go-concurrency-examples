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
	fanOutSem()
}

// fanOutSem: You are a developer and you build one new robot for the exact amount
// of work you have to get done. Each new robot knows immediately what they
// are expected to do and starts their work. However, you don't want all the
// robots working at once. You want to limit how many of them are working at
// any given time. You sit waiting for all the results of the robots work.
// The amount of time you wait on the robots is unknown because you need a
// guarantee that all the results sent by robots are received by you. No
// given robot needs an immediate guarantee that you received their result.

func fanOutSem() {
	ch := make(chan string)
	cpu := runtime.NumCPU()
	limit := make(chan bool, cpu)

	works := 1000

	for i := 0; i < works; i++ {
		go func(id int) {
			limit <- true
			{
				time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
				ch <- fmt.Sprintf("robot %v finish working on task %v\n", id, id)
			}
			<-limit
		}(i)
	}

	for works > 0 {
		result := <-ch
		works--
		fmt.Printf("You received the result: %v", result)

	}
	time.Sleep(time.Second)
	close(ch)
}
