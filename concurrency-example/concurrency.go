package main

import (
	"errors"
	"fmt"
	"github.com/quminzhi/go-examples/concurrency-example/pool"
	"github.com/quminzhi/go-examples/concurrency-example/runner"
	"io"
	"log"
	"math/rand"
	"strconv"
	"sync"
	"sync/atomic"
	"time"
)

// Demo0 compares buffered channel with empty element and unbuffered channel
func Demo0() {
	ch := make(chan int, 1)

	go func(ch chan int) {
		time.Sleep(10 * time.Millisecond)
		close(ch) // goready() all waiting go routines
	}(ch)

	if v, ok := <-ch; !ok {
		fmt.Println("Buffered channel closed, no block")
	} else {
		fmt.Println("Buffered channel is open, got", v)
	}
}

// ===========================================================

func createFunc() func(int) {
	return func(tid int) {
		time.Sleep(1 * time.Second)
		fmt.Println("Task #" + strconv.Itoa(tid) + " completed.")
	}
}

func Demo1() {
	r := runner.NewRunner(4 * time.Second)
	r.AddTasks(createFunc(), createFunc(), createFunc())
	err := r.Start()
	switch {
	case errors.Is(err, runner.ErrInterrupt):
		fmt.Println("Tasks interrupted.")
	case errors.Is(err, runner.ErrTimeout):
		fmt.Println("Tasks timeout.")
	default:
		fmt.Println("Tasks finished.")
	}
}

// ===========================================================

type DBConnection struct {
	id int32
}

func (d *DBConnection) Close() error {
	fmt.Println("Closing DB connection " + fmt.Sprint(d.id))
	return nil
}

var counter int32 // for resources, race condition when multiple goroutines call factory

func Factory() (io.Closer, error) {
	atomic.AddInt32(&counter, 1)
	// Return a pointer to the object because the return type is an interface
	return &DBConnection{id: counter}, nil
}

const (
	POOLSIZE = 2
	NWORKER  = 10
)

func PerformQuery(jid int, p *pool.Pool) {
	defer wg.Done()
	conn, err := p.AcquireResources()
	if err != nil {
		log.Fatalln(err)
	}
	defer func() {
		if err := p.ReleaseResources(conn); err != nil {
			log.Fatalln(err)
		}
	}()

	// Do some work here with conn
	t := rand.Int()%3 + 1
	time.Sleep(time.Duration(t) * time.Second)
	fmt.Println("Task #" + strconv.Itoa(jid) + " completed.")
}

var wg sync.WaitGroup

func Demo2() {
	p, err := pool.NewPool(Factory, POOLSIZE)
	if err != nil {
		log.Fatal(err)
	}

	wg.Add(NWORKER)
	for i := 0; i < NWORKER; i++ {
		go PerformQuery(i, p)
	}

	wg.Wait()

	if err := p.Close(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	Demo2()
}
