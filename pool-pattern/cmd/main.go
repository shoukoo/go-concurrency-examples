package main

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"pool"
	"sync"
	"sync/atomic"
	"time"
)

const (
	maxGoroutines = 25 // the number of routines to use.
	numPooled     = 2  // number of resources in the pool
)

type db struct {
	Id int32
}

func (d *db) Close() error {
	fmt.Printf("Closed connection id %v \n", d.Id)
	return nil
}

var idCounter int32

func createConnection() (io.Closer, error) {
	id := atomic.AddInt32(&idCounter, 1)
	fmt.Println("Create new connection ", id)
	return &db{id}, nil
}

// performQueries tests the resource pool of connections.
func performQueries(query int, p *pool.Pool) {

	// Acquire a connection from the pool.
	conn, err := p.Acquire()
	if err != nil {
		log.Println(err)
		return
	}

	// Release the connection back to the pool.
	defer p.Release(conn)

	// Wait to simulate a query response.
	log.Printf("Query: QID[%d] CID[%d]\n", query, conn.(*db).Id)
}

func main() {
	var wg sync.WaitGroup
	wg.Add(maxGoroutines)

	// Create the pool to manage our connections.
	p, err := pool.New(numPooled, createConnection)
	if err != nil {
		log.Println(err)
		return
	}

	// Perform queries using connections from the pool.
	for query := 0; query < maxGoroutines; query++ {

		// Each goroutine needs its own copy of the query
		// value else they will all be sharing the same query
		// variable.
		go func(q int) {
			time.Sleep(time.Duration(rand.Intn(10000)) * time.Millisecond)
			performQueries(q, p)
			wg.Done()
		}(query)
	}

	// Wait for the goroutines to finish.
	wg.Wait()

	// Close the pool.
	log.Println("Shutdown Program.")
	p.Close()
}
