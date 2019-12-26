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

// pooling: You are a manager and you hire a team of employees. None of the new
// employees know what they are expected to do and wait for you to provide work.
// When work is provided to the group, any given employee can take it and you
// don't care who it is. The amount of time you wait for any given employee to
// take your work is unknown because you need a guarantee that the work your
// sending is received by an employee.
func pooling() {
	workers := runtime.NumCPU()
	ch := make(chan string)

	for i := 0; i < workers; i++ {
		go func(id int) {
			for v := range ch {
				time.Sleep(time.Duration(rand.Intn(1000)) * time.Microsecond)
				fmt.Printf("Worker %+v completed task %v \n", id, v)
			}
		}(i)
	}

	for i := 0; i < 100; i++ {
		ch <- fmt.Sprintf("worker %v", i)
	}
	time.Sleep(time.Second)
	close(ch)

}
