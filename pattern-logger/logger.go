package logger

import (
	"fmt"
	"io"
	"sync"
)

type Logger struct {
	ch chan string
	wg sync.WaitGroup
}

func New(w io.Writer, capacity int) *Logger {

	l := Logger{
		ch: make(chan string, capacity),
	}

	l.wg.Add(1)

	go func() {
		for c := range l.ch {
			fmt.Fprintln(w, c)
		}
		l.wg.Done()
	}()

	return &l

}

func (l *Logger) ShutDown() {
	close(l.ch)
	l.wg.Wait()
}

func (l *Logger) Write(data string) {

	select {
	case l.ch <- data:
		fmt.Println("Write to the channel")
	default:
		fmt.Println("Drop the write")
	}

}
