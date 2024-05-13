package main

import (
	"fmt"
    "io"
    "os"
    "log"
    "unicode"
)

func main() {
	fmt.Println("Hello JSON.")

    stdin, err := io.ReadAll(os.Stdin)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("%s\n", stdin)

    data := string(stdin[:])

    for index, char := range data {
        fmt.Printf("Index: %d Char: %q IsLetter: %t\n", index, char, unicode.IsLetter(char))
    }
}
