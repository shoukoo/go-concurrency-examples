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

// fanOutSem: You are a manager and you hire one new employee for the exact amount
// of work you have to get done. Each new employee knows immediately what they
// are expected to do and starts their work. However, you don't want all the
// employees working at once. You want to limit how many of them are working at
// any given time. You sit waiting for all the results of the employees work.
// The amount of time you wait on the employees is unknown because you need a
// guarantee that all the results sent by employees are received by you. No
// given employee needs an immediate guarantee that you received their result.
func fanOutSem() {
	employees := 200
	ch := make(chan string, employees)

	cpu := runtime.NumCPU()
	sem := make(chan bool, cpu)

	// those goroutines are deposible.
	for i := 0; i < employees; i++ {
		go func(id int) {
			sem <- true
			{
				time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
				ch <- fmt.Sprintf("employee %v completed task\n", id)
			}
			<-sem
		}(i)
	}

	for employees > 0 {
		task := <-ch
		employees--
		fmt.Printf("%v", task)

	}
	time.Sleep(time.Second)
	close(ch)
	fmt.Println("COmpleted")
}
