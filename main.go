package main

import (
	"fmt"
    "io"
    "os"
    "log"
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
        fmt.Printf("Index: %d Char: %q\n", index, char)
    }
}
