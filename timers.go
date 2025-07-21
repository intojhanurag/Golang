package main

import (
    "fmt"
    "time"
)

func main() {

    timer1 := time.NewTimer(2 * time.Second)
	// this timer will fire after 2 second

    <-timer1.C
    fmt.Println("Timer 1 fired") // after 2 second it will print

    timer2 := time.NewTimer(time.Second) // fired after 1 second
    go func() {
        <-timer2.C
        fmt.Println("Timer 2 fired")
    }()
    stop2 := timer2.Stop()
    if stop2 {
        fmt.Println("Timer 2 stopped")
    }

    time.Sleep(2 * time.Second)
}