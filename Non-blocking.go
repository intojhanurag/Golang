package main

import "fmt"

func main() {
    messages := make(chan string) //unbuffered channel
    signals := make(chan bool)   //unbuffered channel (Noone is sending or receiving)

    select {
		//no element currently in the channel message, so default print
    case msg := <-messages:
        fmt.Println("received message", msg)
    default:
        fmt.Println("no message received")
    }

    msg := "hi"
    // select {
	// 	// we are trying to send "Hi" msg to the channel message
	// 	// but noone is ready to receive from the channel so it normally blcok
    // case messages <- msg:
    //     fmt.Println("sent message", msg)
    // default:
    //     fmt.Println("no message sent")
    // }

	go func() {messages <- msg}()

	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)

		//Now select will wait until one of the channel is ready
	}


    // select {
    // case msg := <-messages:
    //     fmt.Println("received message", msg)
    // case sig := <-signals:
    //     fmt.Println("received signal", sig)
    // default:
    //     fmt.Println("no activity")
    // }
}