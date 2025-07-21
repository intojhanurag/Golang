package main

import (
    "fmt"         
    "os"           
    "os/signal"    // to catch OS signals like Ctrl+C
    "syscall"      // for specific system call constants like SIGINT
)

func main() {

    sigs := make(chan os.Signal, 1)
    // Create a channel named 'sigs' to receive signals (buffer size 1)

    signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
    // Tell Go to notify 'sigs' channel if SIGINT (Ctrl+C) or SIGTERM is received

    done := make(chan bool, 1)
    // Create a 'done' channel to signal the main function when to exit

    go func() {
        // Start a background goroutine to handle incoming signals

        sig := <-sigs
        // Wait until a signal is received on the 'sigs' channel

        fmt.Println()
        // Print a newline for clean output

        fmt.Println(sig)
        // Print the received signal (e.g., "interrupt")

        done <- true
        // Send true to the 'done' channel to unblock the main function
    }()

    fmt.Println("awaiting signal")
    // Inform the user that the program is waiting for a signal

    <-done
    // Block here until we receive a value on the 'done' channel

    fmt.Println("exiting")
    // Print exit message when done
}
