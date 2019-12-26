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

// boundedWorkPooling: You are a manager and you hire a team of employees. None of
// the new employees know what they are expected to do and wait for you to
// provide work. The amount of work that needs to get done is fixed and staged
// ahead of time. Any given employee can take work and you don't care who it is
// or what they take. The amount of time you wait on the employees to finish
// all the work is unknown because you need a guarantee that all the work is
// finished.

/**
- you don't care who it is and what they take
- The amount of time you wait on teh rmployees to finish al the work is un known
- guarantee that all the work is finished

**/
func boundedWorkPooling() {
	var wg sync.WaitGroup
	works := []string{"work1", "work2", "work3", "work4", "work5"}
	workers := runtime.NumCPU()
	ch := make(chan string)
	wg.Add(len(works))

	for i := 0; i < workers; i++ {
		go func(id int) {
			defer wg.Done()
			for c := range ch {
				time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
				fmt.Printf("worker %v completed %v", id, c)
			}
		}(i)
	}

	for _, w := range works {
		ch <- w
	}
	close(ch)
	wg.Wait()

}
