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
	fanOut()
}

// fanOut: You are a manager and you build one new robot for the exact amount
// of work you have to get done. Each new robot pre-programmed to know what they
// are expected to do and starts their work. You sit waiting for all the results
// of the work. The amount of time you wait on the rebots is
// unknown because you need a guarantee that all the results sent by robots
// are received by you. No given employee needs an immediate guarantee that you
// received their result.

func fanOut() {
	works := 2000
	ch := make(chan string, works)
	for i := 0; i < works; i++ {
		go func(id int) {
			time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
			ch <- fmt.Sprintf("Robot %v finished working on task %v", id, id)
		}(i)
	}

	for works > 0 {
		result := <-ch
		works--
		fmt.Printf("%v\n", result)
	}

}
