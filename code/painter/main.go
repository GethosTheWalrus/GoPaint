package main

import (
	// "fmt"
	"Lexer"
	"Parser"
)

func main() {

	program := "paint rectangle at (x=10, y=10, width=20, height=20)"

	tokens := Lexer.Tokenize(program)

	Parser.Parse(tokens)
	// for _, token := range(tokens) {

	// 	fmt.Printf("%+v\n", token)

	// }

}