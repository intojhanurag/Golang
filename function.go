package main

import "fmt"

func plus(a int, b int) int {

    return a + b
}

func plusPlus(a, b, c int) int {
    return a + b + c
}

func vals() (int, int) {
    return 3, 7
}
func main() {

    res := plus(1, 2)
    fmt.Println("1+2 =", res)

    res = plusPlus(1, 2, 3)
    fmt.Println("1+2+3 =", res)

	a,b:=vals()
	fmt.Println(a)
	fmt.Println(b)

	_,c:=vals()
	// here the behaviour of _ not work like map , that we saw in the map file 
	// it will store just two value of c and ignore the 1st value
	fmt.Println(c)
}