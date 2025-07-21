package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func main() {

    scanner := bufio.NewScanner(os.Stdin)
	// bufio.NewScanner() reads input from a source.

	// Here the source is os.Stdin → standard input (keyboard).

    for scanner.Scan() {

        ucl := strings.ToUpper(scanner.Text())

        fmt.Println(ucl)

		// scanner.Scan() reads one line at a time from input.

		// .Text() gets the line as a string.

		// strings.ToUpper(...) converts it to UPPERCASE.

		// fmt.Println(...) prints it.
    }

    if err := scanner.Err(); err != nil {
        fmt.Fprintln(os.Stderr, "error:", err)
        os.Exit(1)
    }
}