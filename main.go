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

	//var tokens []string

    token_stream := lexer(data)

    fmt.Println(token_stream)

    os.Exit(parser(token_stream))

}

func lexer(s string) []string {
    // Takes in a text string and returns a slice of tokens.

    // Tokens you need to recognize:
    //      [ begin-array
    //      ] end-array
    //      { begin-object
    //      } end-object
    //      : name-separator
    //      , value-separator
    //      false
    //      true
    //      null
    //      value false/null/true/object/array/number/string
    //
    // All strings start and end with quotes(").

    tokens := []string{}
    lastchar := ""
    identifier := ""
    inString := false

    for _, c := range s {
        char := string(c)
        if char == "{" {
            tokens = append(tokens, string(char))
            lastchar = char
            continue
        }

        if char == "}" {
            tokens = append(tokens, string(char))
            lastchar = char
            continue
        }

        if char == "[" {
            tokens = append(tokens, string(char))
            lastchar = char
            continue
        }
        
        if char == "]" {
            tokens = append(tokens, string(char))
            lastchar = char
            continue
        }

        if char == ":" {
            tokens = append(tokens, string(char))
            lastchar = char
            continue
        }

        if char == "," {
            tokens = append(tokens, string(char))
            lastchar = char
            continue
        }

        if char == "\"" && ! inString {
            inString = true
            continue
        }

        if char == "\"" && inString {
            inString = false
            tokens = append(tokens, string(identifier))
            continue
        }

        if inString && unicode.IsLetter(c) {
            identifier += char
            continue
        }

    } // end for loop

    fmt.Println(lastchar)

    return tokens

} // end lexer()

func parser(s []string) int {
    if len(s) < 2 {
        return 1
    }

    if s[0] == "{" && s[1] == "}" {
        return 0
    }

    return 2
}
