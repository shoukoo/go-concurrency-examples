package main

import (
	"context"
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
	//waitForResult()
	//fanOut()
	//waitForTask()
	//pooling()
	//fanOutSem()
	//boundedWorkPooling()
	//cancellation()
	//tennis()

	// Write a program that uses goroutines to generate up to 100 random numbers.
	// Do not send values that are divisible by 2. Have the main goroutine receive
	// values and add them to a slice.
	//genNum()

	// Write a program that creates a fixed set of workers to generate random
	// numbers. Discard any number divisible by 2. Continue receiving until 100
	// numbers are received. Tell the workers to shut down before terminating.
	genNumWithLimit()
}

func genNumWithLimit() {

	workers := runtime.NumCPU()
	ch := make(chan int, workers)
	shutdown := make(chan struct{})
	var wg sync.WaitGroup
	wg.Add(workers)

	for i := 0; i < workers; i++ {
		go func() {
			for {
				r := rand.Intn(1000)
				select {
				case ch <- r:
				case <-shutdown:
					wg.Done()
					return
				}
			}
		}()
	}

	var total []int
	for c := range ch {

		if len(total) == 100 {
			close(shutdown)
			break
		}

		if c%2 == 0 {
			continue
		}

		total = append(total, c)
	}

	wg.Wait()
	fmt.Printf("total = %+v\n", total)

}

func genNum() {
	goroutines := 300
	ch := make(chan int, goroutines)
	var wg sync.WaitGroup
	wg.Add(goroutines)

	for i := 0; i < goroutines; i++ {
		go func() {
			r := rand.Intn(1000)
			if r%2 != 0 {

				ch <- r
			}
			wg.Done()
		}()
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	var total []int
	for c := range ch {
		total = append(total, c)
	}
	fmt.Printf("total = %+v\n", total)

}

func waitForResult() {

	ch := make(chan string)

	go func() {
		ch <- "done"
	}()

	task := <-ch

	time.Sleep(time.Second)

	fmt.Printf("task = %+v\n", task)

}

func fanOut() {

	tasks := 100

	ch := make(chan string)

	for i := 0; i < tasks; i++ {
		go func(id int) {
			fmt.Printf("working on task %+v\n", i)
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
			ch <- fmt.Sprintf("completed task %v", id)
		}(i)
	}

	for tasks > 0 {
		tasks--
		task := <-ch
		fmt.Printf("Completed task %+v\n", task)
	}

}

func waitForTask() {

	ch := make(chan string)

	go func() {
		task := <-ch
		fmt.Printf("Recv this task = %+v\n", task)
	}()

	ch <- "send an email"

}

func pooling() {
	workers := runtime.NumCPU()
	ch := make(chan int)
	tasks := 100

	//create workers
	for i := 0; i < workers; i++ {
		go func(id int) {
			for c := range ch {
				time.Sleep(time.Duration(100) * time.Millisecond)
				fmt.Printf(" worker %v working on %v \n", id, c)
			}
		}(i)
	}

	for i := 0; i < tasks; i++ {
		ch <- i
	}
	time.Sleep(time.Duration(1000) * time.Millisecond)
}

func fanOutSem() {
	ch := make(chan string)
	limit := make(chan bool, 5)
	tasks := 300
	for i := 0; i < tasks; i++ {
		go func(id int) {
			limit <- true
			{
				time.Sleep(time.Duration(100) * time.Millisecond)
				ch <- fmt.Sprintf("task %v is done", id)
			}
			<-limit
		}(i)
	}

	for tasks > 0 {
		tasks--
		fmt.Printf("task = %+v\n", <-ch)
	}
	time.Sleep(time.Second)
}

func boundedWorkPooling() {
	ch := make(chan int)
	var wg sync.WaitGroup
	workers := runtime.NumCPU()
	wg.Add(workers)
	tasks := 200

	for i := 0; i < workers; i++ {
		go func() {
			defer wg.Done()
			for c := range ch {
				time.Sleep(time.Duration(100) * time.Millisecond)
				fmt.Printf("working on task id %v\n", c)
			}
		}()
	}

	for i := 0; i < tasks; i++ {
		ch <- i
	}

	close(ch)
	wg.Wait()

}

func cancellation() {

	ch := make(chan string)

	c, cancel := context.WithTimeout(context.Background(), time.Duration(2000)*time.Millisecond)
	defer cancel()

	go func() {
		time.Sleep(time.Duration(1000) * time.Millisecond)
		ch <- "hello"
	}()

	select {
	case <-c.Done():
		fmt.Println("Expired")
	case <-ch:
		fmt.Println("Done")
	}
}

//-------------------------------------------------------------------------------------------------

func tennis() {

	var wg sync.WaitGroup
	ch := make(chan int)
	wg.Add(2)

	go func() {
		ball("Andy", ch)
		wg.Done()
	}()

	go func() {
		ball("Chris", ch)
		wg.Done()
	}()

	ch <- 0

	wg.Wait()
}

func ball(name string, ch chan int) {
	for {

		ball, alive := <-ch
		if !alive {
			fmt.Printf("%s won the match\n", name)
			return
		}

		flip := rand.Intn(100)
		if flip%13 == 0 {
			fmt.Printf("%v misses the ball\n", name)
			close(ch)
			return
		}

		fmt.Printf("%v hits back the ball\n", name)
		ch <- (ball + 1)

	}
}

//--------------------------------------------------------------------------------------------------

func test1() {
	// Write a program that uses goroutines to generate up to 100 random numbers.
	// Do not send values that are divisible by 2. Have the main goroutine receive
	// values and add them to a slice.
	goroutines := 100

	// Create the channel for sharing results.
	ch := make(chan int)
	var wg sync.WaitGroup
	wg.Add(goroutines)

	for i := 0; i < goroutines; i++ {
		go func() {
			defer wg.Done()
			r := rand.Intn(100)
			if r%2 == 0 {
				return
			}
			ch <- r
		}()
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	var t []int
	for c := range ch {
		t = append(t, c)
	}

	fmt.Printf("t = %+v\n", len(t))

}

func test2() {

	workers := runtime.NumCPU()
	var wg sync.WaitGroup
	wg.Add(workers)
	ch := make(chan int)
	shutdown := make(chan struct{})
	for i := 0; i < workers; i++ {
		go func() {
			for {
				r := rand.Intn(1000)
				select {
				case ch <- r:
					fmt.Println("sending channel")
				case <-shutdown:
					fmt.Println("shutdown")
					wg.Done()
					return
				}
			}
		}()
	}

	var t []int
	for c := range ch {

		if len(t) == 100 {
			break
		}

		if c%2 != 0 {
			t = append(t, c)
		}

	}
	close(shutdown)

	fmt.Printf("len(t) = %+v\n", len(t))

	wg.Wait()

}
