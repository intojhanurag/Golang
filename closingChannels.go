package main

import "fmt"

func main() {
    jobs := make(chan int, 5) //capacity=5
    done := make(chan bool)

    go func() {
        for {
            j, more := <-jobs
            if more {
                fmt.Println("received job", j)
            } else {
                fmt.Println("received all jobs")
                done <- true
                return
            }
        }
    }()

    for j := 1; j <= 3; j++ {
        jobs <- j
        fmt.Println("sent job", j)
    }
    close(jobs) // we are closing the job channel , so that more becomes false 
	// and can exit the loop 
	
    fmt.Println("sent all jobs")

    <-done

    _, ok := <-jobs // because now channel is empty 
    fmt.Println("received more jobs:", ok)
}