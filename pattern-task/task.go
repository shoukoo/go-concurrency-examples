package task

import "sync"

// Worker needs to implement Do function
type Worker interface {
	Do()
}

// Task has a worker channel and a waitgroup
// worker channel is to send works to goroutines
// that run in the background and waitgroup is to ensure
// all goroutines are shutdown properly.
type Task struct {
	ch chan Worker
	wg sync.WaitGroup
}

// New create a new work pool
func New(goroutines int) *Task {

	t := Task{
		ch: make(chan Worker),
	}

	// goroutines are the pool, we could adjust this size
	// later on
	t.wg.Add(goroutines)
	for i := 0; i < goroutines; i++ {
		go func() {
			defer t.wg.Done()
			for c := range t.ch {
				c.Do()
			}
		}()

	}

	return &t
}

// Shutdown waits for all the goroutines to shutdown
func (t *Task) ShutDown() {
	close(t.ch)
	t.wg.Wait()
}

//Do submits the work to the pool
func (t *Task) Do(w Worker) {
	t.ch <- w
}
