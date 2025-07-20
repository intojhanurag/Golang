package main

import "fmt"

type ServerState int

const (
    StateIdle ServerState = iota //0 // we write serverstate one time and go remember from next line
    StateConnected //1
    StateError //2
    StateRetrying //3
)

var stateName = map[ServerState]string{
    StateIdle:      "idle",
    StateConnected: "connected",
    StateError:     "error",
    StateRetrying:  "retrying",
}

func (ss ServerState) String() string {
    return stateName[ss]
	//String interface
}

//"We are implementing the String() method from the fmt.Stringer interface to customize how our type is printed."

func main() {
    ns := transition(StateIdle)
    fmt.Println(ns)

    ns2 := transition(ns)
    fmt.Println(ns2)


}

func transition(s ServerState) ServerState {
    switch s {
    case StateIdle:
        return StateConnected
    case StateConnected, StateRetrying:

        return StateIdle
    case StateError:
        return StateError
    default:
        panic(fmt.Errorf("unknown state: %s", s))
    }
}