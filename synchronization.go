package main

import (
    "fmt"
    "time"
)

func worker(done chan bool) {
    fmt.Print("working...") 
    time.Sleep(time.Second) //simulates some work, wait 1 second
    fmt.Println("done")

    done <- true // signal that work is done
}

func main() {

    done := make(chan bool, 1)
	//a channel that will carry one bool value
	// used for signaling completion
	// Buffer size 1 mean you can send 1 message without waiting
    go worker(done)
	// start go routing , worker function runs in parallel to main()

    <-done
	// wait here and receive a value from the channel named
	// once received continue
}