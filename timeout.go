package main

import (
    "fmt"
    "time"
)

func main() {

    c1 := make(chan string, 1)
    go func() {
        time.Sleep(2 * time.Second) // wait for 2 second
        c1 <- "result 1"
    }()

    select {
    case res := <-c1:
        fmt.Println(res)
    case <-time.After(1 * time.Second): // wait for 1 second
        fmt.Println("timeout 1")
    }
	//1 < 2 , so timeout 1 print first

    c2 := make(chan string, 1)
    go func() {
        time.Sleep(2 * time.Second)// wait for 2 second
        c2 <- "result 2"
    }()
    select {
    case res := <-c2:
        fmt.Println(res)
    case <-time.After(3 * time.Second):// wait for 3 second
        fmt.Println("timeout 2")
    }

	// 3>2 ,so res=result2 from the channel in 2 second and got print before case2
}