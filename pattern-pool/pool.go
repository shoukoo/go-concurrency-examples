package pool

import (
	"errors"
	"fmt"
	"io"
	"sync"
)

type Pool struct {
	ch      chan io.Closer
	factory func() (io.Closer, error)
	closed  bool
	mu      sync.Mutex
}

var ErrPoolClosed = errors.New("Pool has been closed")

func New(size int, f func() (io.Closer, error)) (*Pool, error) {
	if size == 0 {
		return nil, errors.New("Size values too small")
	}

	return &Pool{
		ch:      make(chan io.Closer, size),
		factory: f,
	}, nil
}

func (p *Pool) Acquire() (io.Closer, error) {

	select {
	case v, alive := <-p.ch:
		fmt.Println("Acquired Shared Resource")
		if !alive {
			return nil, ErrPoolClosed
		}
		return v, nil
	default:
		fmt.Println("Acquired New Resource")
		r, _ := p.factory()
		return r, nil
	}
}

func (p *Pool) Release(r io.Closer) {

	p.mu.Lock()
	defer p.mu.Unlock()

	if p.closed == true {
		r.Close()
		return
	}

	select {
	case p.ch <- r:
		fmt.Println("Put back the resources")
	default:
		fmt.Println("Drop the Resource")
		r.Close()
	}
}

func (p *Pool) Close() error {

	p.mu.Lock()
	defer p.mu.Unlock()

	if p.closed {
		return ErrPoolClosed
	}

	p.closed = true

	close(p.ch)

	for c := range p.ch {
		c.Close()
	}
	return nil
}
