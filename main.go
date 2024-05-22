package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"unicode"
)

func main() {
	fmt.Println("Hello JSON.")

	stdin, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	data := string(stdin[:])

	token := ""

	var tokens []string

	for _, char := range data {

		if unicode.IsLetter(char) {
			token += string(char)
			continue
		}

		if unicode.IsDigit(char) {
			token += string(char)
			continue
		}

		tokens = append(tokens, token)
		token = ""
		tokens = append(tokens, string(char))

	}

	fmt.Println(tokens)
}
