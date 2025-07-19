package main

import "fmt"

//closures in Go
// Function intSeq returns Another function = > 
// that returned function has no parameters, and return an int
func intSeq() func() int {
    i := 0
    return func() int {
        i++
        return i
    }
}

// So it is returning func () int{ i++ return i}

func main() {

    nextInt := intSeq()

    fmt.Println(nextInt()) // i=1
    fmt.Println(nextInt()) //  i=2

    fmt.Println(nextInt()) // i=3

    newInts := intSeq() // Diff closure diff memory
	
    fmt.Println(newInts())//1
	fmt.Println(newInts())//2
}