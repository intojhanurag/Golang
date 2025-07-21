package main

import (
    "fmt"
    "sync"
    "sync/atomic"
)

func main() {

    var ops atomic.Uint64 // atomic counter
	// this is a counter of type uint64
	// thread safe without needing a mutex
	// No race condition

    var wg sync.WaitGroup // waitgroup 

    for range 50 {
        wg.Add(1) // tell waitgroup 1 gorouting is coming

        go func() { //Launch a go routine
            for range 1000 {

                ops.Add(1) // Each go routing adds 1000 times
				// Atomic increment (thread safe)
            }

            wg.Done() // go routine is finished
        }()
    }

	// if we will not use atomic and use simple variable then all 50 workers messed up the adding
	// they will write read same variable we will not get consistent result

	// We use waitgroup to finish all 50 work , so that we can finish the program after completion

    wg.Wait() // wait for all go routine

    fmt.Println("ops:", ops.Load()) // 50 goroutines * 10000 increment=50000
}