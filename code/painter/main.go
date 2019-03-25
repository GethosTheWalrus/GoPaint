package main

import (
	"fmt"
	"lexer"
)

func main() {

	program := "paint rectangle at (x=10, y=10, width=20, height=20)"

	tokens := lexer.Tokenize(program)

	for _, token := range(tokens) {

		fmt.Printf("%+v\n", token)

	}

}