package main

import (
	"fmt"
	"io"
	"log"
	"os"
	//"unicode"
)

func main() {
	fmt.Println("Hello JSON.")

	stdin, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	data := string(stdin[:])

	var tokens []string

    tokens = lexer(data)

    os.Exit(parser(tokens))

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
    //lastchar := ""
    //token := ""

    for _, c := range s {
        char := string(c)
        if char == "{" {
            tokens = append(tokens, string(char))
            continue
        }

        if char == "}" {
            tokens = append(tokens, string(char))
            continue
        }

        if char == "[" {
            tokens = append(tokens, string(char))
            continue
        }
        
        if char == "]" {
            tokens = append(tokens, string(char))
            continue
        }

        if char == ":" {
            tokens = append(tokens, string(char))
            continue
        }

        if char == "," {
            tokens = append(tokens, string(char))
            continue
        }

    } // end for loop

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
