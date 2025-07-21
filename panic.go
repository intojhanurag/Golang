package main

import "os"

func main() {

    panic("a problem") // stop the program right now and throw an error, in emergency

	//it doesn't go to the next line and throw runtime error and exit

    _, err := os.Create("/tmp/file")
    if err != nil {
        panic(err)
    }

	// error: This means the panic happened in goroutine 1, which is your main goroutine (the main function).

	//Goroutine 1 is the default thread that runs your program unless you spawn others using go keyword.
}