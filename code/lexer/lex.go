package lexer

import (
	"fmt"
	"strings"
	"globals"
)

func Tokenize(program string) {

	// initialize lexing variables
	char_num, line_num, token_num := 0, 0, 0
	token_buffer := ""
	word_in_buffer, number_in_buffer := false, false

	// break the program into characters
	program_chars := strings.Split(program, "")

	// loop through each character and lex
	for _, char := range(program_chars) {

		// token variables
		var token_type string
		var token_value interface{}

		// increment line_num if we encounter a \n
		if char == "\n" {

			line_num += 1
			char_num = 0

			token_type = "newline"
			token_value = char

		// spaces count as tokens
		} else if char == " " {

			token_type = "space"
			token_value = char

		// operator tokens are = : + - * /
		} else if globals.IsOperator(char) {

			token_type = "operator"
			token_value = char

		// separator tokens are ( ) ,
		} else if globals.IsSeparator(char) {

			token_type = "separator"
			token_value = char

		// variables, keywords, etc
		} else if globals.IsAlpha(char)
		
			
		
		} else if globals.IsDigit(char) {


		}

		fmt.Println(token_type, token_value, token_num, char_num)

	}

}