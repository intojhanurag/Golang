package main

import (
    "fmt"
    "time"
)

func f(from string) {
    for i := range 3 {
        fmt.Println(from, ":", i)
    }
	//from:0 
	//from :1 
	//from:2

}

func main() {

    f("direct")

    go f("goroutine") //this function run in the background
	// But the main fucntion might finish before goroutine prints anything
	// so we fix later using time.sleep()

    go func(msg string) {
        fmt.Println(msg)
    }("going") //a function without a name, called as gorouting
	// immediately prints going

    time.Sleep(time.Second) // Go does not wait for goroutines by default
	// so we use Sleep 1 second to give time for goroutines to finish
    fmt.Println("done")
}