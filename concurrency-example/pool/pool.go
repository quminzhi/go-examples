package pool

import (
	"errors"
	"fmt"
	"io"
	"sync"
)

var (
	ErrPoolClosed      = errors.New("pool is closed")
	ErrCloseResource   = errors.New("close resource error")
	ErrCloseClosedPool = errors.New("close closed pool error")
	ErrCreateResource  = errors.New("create resource error")
)

// Pool provides shared resources for users (go-routines) to use.
// Use buffered channel to cache shared resources.
type Pool struct {
	factory   func() (io.Closer, error) // the method to create resources
	resources chan io.Closer

	mtx    sync.Mutex // mutex on resources channel
	closed bool
}

func NewPool(factory func() (io.Closer, error), size int) (*Pool, error) {
	if size <= 0 {
		return nil, errors.New("size must be greater than zero")
	}

	p := &Pool{
		factory:   factory,
		resources: make(chan io.Closer, size),
		closed:    false,
	}

	// Prepare some resources in advance
	for i := 0; i < size; i++ {
		res, err := factory()
		if err != nil {
			return nil, ErrCreateResource
		}
		p.resources <- res
	}

	return p, nil
}

func (p *Pool) AcquireResources() (io.Closer, error) {
	select {
	case res, ok := <-p.resources:
		// Closed or read the resource from buffered channel
		if !ok {
			// Pool closed
			return nil, ErrPoolClosed
		}
		fmt.Println("resource acquired from the pool")
		return res, nil
	default:
		// Buffered channel is empty
		// Create resource with factory method
		fmt.Println("pool is empty, create a new resource")
		return p.factory()
	}
}

// ReleaseResources put the resource back to the buffered channel
func (p *Pool) ReleaseResources(res io.Closer) error {
	// Pool closed
	if p.closed {
		if err := res.Close(); err != nil {
			return ErrCloseResource
		}
		return nil
	}

	p.mtx.Lock()
	defer p.mtx.Unlock()

	select {
	case p.resources <- res:
		fmt.Println("resource released back into the pool")
	default:
		fmt.Println("pool is full, just close the resource")
		if err := res.Close(); err != nil {
			return ErrCloseResource
		}
	}

	return nil
}

func (p *Pool) Close() error {
	p.mtx.Lock()
	defer p.mtx.Unlock()

	if p.closed {
		return ErrCloseClosedPool
	}
	p.closed = true

	// Close channel before read from buffered channel.
	// Range will hang on and deadlock.
	close(p.resources)
	for res := range p.resources {
		if err := res.Close(); err != nil {
			return ErrCloseResource
		}
	}
	return nil
}
