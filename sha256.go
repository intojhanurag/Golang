package main

import (
    "crypto/sha256"
    "fmt"
)

func main() {
    s := "sha256 this string"

    h := sha256.New()
	// crypto/sha256: Provides SHA-256 hash function.

    h.Write([]byte(s))

	// SHA-256 requires a byte slice as input.
	// []byte(s) converts the string "sha256 this string" to bytes.


    bs := h.Sum(nil)
	// nil means : “Start fresh. Just return me the hash result, not appended to anything.”

	// h.Sum(nil) finalizes the hash and returns the result as a byte slice ([32]byte).

    fmt.Println(s)
    fmt.Printf("%x\n", bs)
}