// A Go string is a read-only slice of bytes.
// The language treat strings specially - as containers of text encoded in UTF-8.
// In Go, the concept of character is called a rune - it's an integer that represents a Unicode code point.
package main

import (
    "fmt"
    "unicode/utf8"
)

func main() {

    // s is assigned a literal value representing the word "hello" in the Thai language.
    const s = "สวัสดี"

    // Since strings are equivalent to [] byte, this will produce the length
    // of the raw bytes stored within.
    fmt.Println("Len:", len(s))

    for i := 0; i < len(s); i++ {
        fmt.Printf("%x", s[i])
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
