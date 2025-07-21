package main

import (
    "fmt"
    "io/fs"
    "os"
    "path/filepath"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {

    err := os.Mkdir("subdir", 0755)

	//0755 - linux file permission
    check(err)

    defer os.RemoveAll("subdir")
	// remove everthing after the program finishes

    createEmptyFile := func(name string) {
        d := []byte("")
        check(os.WriteFile(name, d, 0644))
		// quickly create empty file by giving just a name
		// 0644=read/write for owner, read only for others
    }

    createEmptyFile("subdir/file1")

    err = os.MkdirAll("subdir/parent/child", 0755)

	// MkdirAll() creates all folders in the path — like mkdir -p in Linux.
    check(err)

    createEmptyFile("subdir/parent/file2")
    createEmptyFile("subdir/parent/file3")
    createEmptyFile("subdir/parent/child/file4")

    c, err := os.ReadDir("subdir/parent")
    check(err)

    fmt.Println("Listing subdir/parent")
    for _, entry := range c {
        fmt.Println(" ", entry.Name(), entry.IsDir())
    }

    err = os.Chdir("subdir/parent/child")
    check(err)

    c, err = os.ReadDir(".")
    check(err)

    fmt.Println("Listing subdir/parent/child")
    for _, entry := range c {
        fmt.Println(" ", entry.Name(), entry.IsDir())
    }

    err = os.Chdir("../../..")
    check(err)

    fmt.Println("Visiting subdir")
    err = filepath.WalkDir("subdir", visit)
}

func visit(path string, d fs.DirEntry, err error) error {
    if err != nil {
        return err
    }
    fmt.Println(" ", path, d.IsDir())
    return nil
}