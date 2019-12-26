package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	var wg sync.WaitGroup

	wg.Add(2)
	ch := make(chan int)

	go func() {
		tennis("Andy", ch)
		wg.Done()
	}()

	go func() {
		tennis("Hiroko", ch)
		wg.Done()
	}()

	go func() {
		tennis("Mei", ch)
		wg.Done()
	}()

	ch <- 0
	wg.Wait()
}

func tennis(name string, ch chan int) {

	for {
		ball, alive := <-ch

		if !alive {
			fmt.Printf("Player won %+v\n", name)
			return
		}

		hit := rand.Intn(100)
		if hit%13 == 0 {
			fmt.Printf("%v Misses the ball\n", name)
			close(ch)
			return
		}

		fmt.Printf("%v Hits the ball back\n", name)
		ball++
		ch <- ball
	}
}
