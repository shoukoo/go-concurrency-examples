package main

import (
	"fmt"
	"task"
)

type Print struct {
	comment string
}

func New(word string) Print {
	return Print{
		comment: word,
	}
}

func (p Print) Do() {
	fmt.Println(p.comment)
}

func main() {
	p1 := New("test1")
	p2 := New("test2")
	task := task.New(2)
	task.Do(p1)
	task.Do(p2)
	task.ShutDown()
	fmt.Println("Done")
}
