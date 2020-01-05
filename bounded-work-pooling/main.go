package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	boundedWorkPooling()
}

// boundedWorkPooling: You are a developer and you hire a team of robots. None of
// the new robots know what they are expected to do and wait for you to
// provide work. The amount of work that needs to get done is fixed and staged
// ahead of time. Any given robot can take work and you don't care who it is
// or what they take. The amount of time you wait on the robots to finish
// all the work is unknown because you need a guarantee that all the work is
// finished.
func boundedWorkPooling() {
	ch := make(chan string)
	works := []string{"task1", "task2", "task3", "task4", 200: "task2000"}
	workers := runtime.NumCPU()
	var wg sync.WaitGroup
	wg.Add(workers)

	for i := 0; i < workers; i++ {
		go func(id int) {
			defer wg.Done()
			for c := range ch {
				time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
				fmt.Printf("Robot %v finished task %+v\n", id, c)
			}
		}(i)
	}

	for i, _ := range works {
		ch <- fmt.Sprintf("%v", i)
	}

	close(ch)
	wg.Wait()
}
