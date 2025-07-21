package main

import (
    "fmt"
    "os"
)

func main() {

    f := createFile("/tmp/defer.txt")
    defer closeFile(f)

	
	// Defer is used to ensure that a function call
	//  is performed later in a program’s execution, 
	//  usually for purposes of cleanup. defer is often 
	//  used where e.g. ensure and finally would be used 
	//  in other languages.
    writeFile(f) // after writing file remember to close the file , and we will ensure that file close due to defer keyword
}

func createFile(p string) *os.File {
    fmt.Println("creating")
    f, err := os.Create(p)
    if err != nil {
        panic(err)
    }
    return f
}

func writeFile(f *os.File) {
    fmt.Println("writing")
    fmt.Fprintln(f, "data")
}

func closeFile(f *os.File) {
    fmt.Println("closing")
    err := f.Close()

    if err != nil {
        panic(err)
    }
}