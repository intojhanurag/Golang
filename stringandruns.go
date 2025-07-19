package main

import (
    "fmt"
    "unicode/utf8"
)

//Old systems used ASCII, which only supports 128 characters:

//English letters (A-Z, a-z)

//Digits (0–9)

//Some symbols (!@#$%^&*) , ASCII don't contain hindi, Emojis, Chinese, Arbaic, Thai
//So we use unicode - it uses 1 to 4 bytes per char

func main() {

    const s = "สวัสดี"

    fmt.Println("Len:", len(s))
	// length=18, because each Thai char take 3 bytes in UTF-8=> 6*3=18

    for i := 0; i < len(s); i++ {
        fmt.Printf("%x ", s[i])
    }
    fmt.Println()

    fmt.Println("Rune count:", utf8.RuneCountInString(s))

    for idx, runeValue := range s {
        fmt.Printf("%#U starts at %d\n", runeValue, idx)
    }

    fmt.Println("\nUsing DecodeRuneInString")
    for i, w := 0, 0; i < len(s); i += w {
        runeValue, width := utf8.DecodeRuneInString(s[i:])
        fmt.Printf("%#U starts at %d\n", runeValue, i)
        w = width

        examineRune(runeValue)
    }
}

func examineRune(r rune) {

    if r == 't' {
        fmt.Println("found tee")
    } else if r == 'ส' {
        fmt.Println("found so sua")
    }
}