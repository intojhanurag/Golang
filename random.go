package main

import (
    "fmt"
    "math/rand/v2"
)

func main() {

    fmt.Print(rand.IntN(100), ",")
	// random integer between 0 to n-1
    fmt.Print(rand.IntN(100))
    fmt.Println()

    fmt.Println(rand.Float64())

    fmt.Print((rand.Float64()*5)+5, ",")
    fmt.Print((rand.Float64() * 5) + 5)

	//rand.Float64() → 0 to 1

	// *5 → 0 to 5

	// +5 → 5 to 10
    fmt.Println()

    s2 := rand.NewPCG(42, 1024)
    r2 := rand.New(s2)

	//You are creating a custom random generator using a fixed seed (42, 1024)

	// This allows you to get the same "random" numbers every time the program runs 
	// — useful for testing, simulations, etc.
    fmt.Print(r2.IntN(100), ",")
    fmt.Print(r2.IntN(100))

	//Since the seed is fixed, these numbers will always be the same
	//  every time you run the program, like:
    fmt.Println()

    s3 := rand.NewPCG(42, 1024)
    r3 := rand.New(s3)
    fmt.Print(r3.IntN(100), ",")
    fmt.Print(r3.IntN(100))
    fmt.Println()
}