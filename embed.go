package main

import (
    "embed"
)

// Using //go:embed tells the compiler:

// "Hey, take this file and embed it into this variable before compiling."

//go:embed folder/single_file.txt
var fileString string

//go:embed folder/single_file.txt
var fileByte []byte

//go:embed folder/single_file.txt
//go:embed folder/*.hash
var folder embed.FS

func main() {

    print(fileString)
    print(string(fileByte))

    content1, _ := folder.ReadFile("folder/file1.hash")
    print(string(content1))

    content2, _ := folder.ReadFile("folder/file2.hash")
    print(string(content2))
}