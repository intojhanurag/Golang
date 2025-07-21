package main

import (
    "fmt"
    "time"
)

func main() {

    requests := make(chan int, 5) // capacity=5
    for i := 1; i <= 5; i++ {
        requests <- i
    }
	//You've created 5 incoming requests numbered 1 to 5.
    close(requests)


    limiter := time.Tick(200 * time.Millisecond)
	// time.Tick(...) returns a channel that sends current time every 200ms

    for req := range requests {
        <-limiter
        fmt.Println("request", req, time.Now())
		//Wait until next tick (200ms) , Then process the request
		// 1 request every - 200second
    }

    burstyLimiter := make(chan time.Time, 3) // size=3

    for range 3 {
        burstyLimiter <- time.Now() // fill the burstyLiniter with time 
    }

    go func() {
        for t := range time.Tick(200 * time.Millisecond) {
            burstyLimiter <- t
        }
		//Every 200ms, we refill the bucket with new token
    }()

    burstyRequests := make(chan int, 5) 
    for i := 1; i <= 5; i++ {
        burstyRequests <- i
    }
    close(burstyRequests)
    for req := range burstyRequests {
        <-burstyLimiter
        fmt.Println("request", req, time.Now())
    }
}