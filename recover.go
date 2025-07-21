package main

import (
	"fmt"
	"time"
)

func mayPanic() {
    panic("a problem")
}

func main() {

    defer func() {
        if r := recover(); r != nil {

            fmt.Println("Recovered. Error:\n", r)
        }
    }()

    mayPanic()

    fmt.Println("After mayPanic()")
	time.Sleep(time.Second * 1)
}