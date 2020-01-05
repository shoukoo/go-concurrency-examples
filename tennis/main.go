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
	ch := make(chan int)
	wg.Add(2)

	go func() {
		player("Andy", ch)
		wg.Done()
	}()

	go func() {
		player("Danny", ch)
		wg.Done()
	}()

	ch <- 1
	wg.Wait()

}

func player(name string, ch chan int) {
	for {

		ball, alive := <-ch

		if !alive {
			fmt.Printf("%+v has won the match\n", name)
			return
		}

		randomness := rand.Intn(1000)

		if randomness%13 == 0 {
			fmt.Printf("%+v misses the ball\n", name)
			close(ch)
			return
		}

		fmt.Printf("%v hits the ball back\n", name)
		ch <- ball + 1
	}
}
