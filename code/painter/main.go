package main

import (
	// "fmt"
	"lexer"
)

func main() {

	program := "paint rectangle at (x=10, y=10, width=20, height=20)"

	lexer.Tokenize(program)

}