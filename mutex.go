package main

import (
    "fmt"
    "sync"
)

type Container struct {
    mu       sync.Mutex
    counters map[string]int
}

func (c *Container) inc(name string) {

    c.mu.Lock() //Lock the shared container
    defer c.mu.Unlock() // unlock the mutex so other container can enter
    c.counters[name]++ // safe to modify

	// this locking ensures that only 1 gorouting touches the map at a time
}

func main() {
    c := Container{

        counters: map[string]int{"a": 0, "b": 0},

		// counter is a shared map across all goroutines
    }

    var wg sync.WaitGroup

    doIncrement := func(name string, n int) {
		// n=1000
        for range n {
            c.inc(name)
        }
        wg.Done()
		// each worker give notification i am done
    }

    wg.Add(3) // register 3 work are pending 
    go doIncrement("a", 10000)
    go doIncrement("a", 10000)
    go doIncrement("b", 10000)

    wg.Wait() // wait here for that 3 workers
    fmt.Println(c.counters)
	// printer after 30000 updates
}