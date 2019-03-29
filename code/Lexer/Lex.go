package Lexer

import (
	// "fmt"
	"strings"
	"globals"
	"models"
)

func Tokenize(program string) []models.Token {

	// initialize lexing variables
	char_num, line_num, token_num := 1, 1, 1
	tokens := make([]models.Token,0)
	
	// break the program into stream[i]acters
	stream := strings.Split(program, "")

	// loop through each stream[i]acter and lex
	for i := 0; i < len(stream); i++ {

		// reserve memory for the current token
		current_token := ""

		// token variables
		var token_type string
		var token_value string

		// increment line_num if we encounter a \n
		if stream[i] == "\n" {

			line_num += 1
			char_num = 0

			token_type = "newline"
			token_value = stream[i]

		// spaces count as tokens
		} else if stream[i] == " " {

			token_type = "space"
			token_value = stream[i]

		// operator tokens are = : + - * /
		} else if globals.IsOperator(stream[i]) {

			token_type = "operator"
			token_value = stream[i]

		// container tokens are ( )
		} else if globals.IsContainer(stream[i]) {

			token_type = "container"
			token_value = stream[i]

		// separator tokens are ,
		} else if globals.IsSeparator(stream[i]) {

			token_type = "separator"
			token_value = stream[i]

		// variables, keywords, etc
		} else if globals.IsAlpha(stream[i]) || globals.IsDigit(stream[i]) {

			was_a_word := false

			// build word tokens
			for current_token_count := 0; globals.IsAlpha(stream[i]); current_token_count++ {

				was_a_word = true
				token_type = "word"
				current_token += stream[i]
				i += 1

			}

			// build number tokens
			for current_token_count := 0; !was_a_word && globals.IsDigit(stream[i]); current_token_count++ {

				token_type = "number"
				current_token += stream[i]
				i += 1

			}

			token_value = current_token
			i -= 1;

		} else {

			for current_token_count := 0; stream[i] != " "; current_token_count++ {

				token_type = "unknown"
				current_token += stream[i]
				i += 1

			}

			token_value = current_token
			i -= 1;

		}

		// identify special tokens
		identify_special_tokens(&token_type, token_value)

		// create the  token
		token := models.Token{token_type, token_value, line_num, char_num, token_num}

		// increase the stream[i]acter counter
		char_num += len(token_value)
		token_num += 1

		// ignore spaces
		if token_type != "space" {

			tokens = append(tokens, token)

		}

	}

	return tokens

}

func identify_special_tokens(token_type *string, token_value string) {

	if *token_type == "word" {

		if token_value == "paint" {

			*token_type = "instruction"

		} else if token_value == "rectangle" {

			*token_type = "object"

		} else if token_value == "at" {

			*token_type = "function"
			
		} else if token_value == "x" ||
					token_value == "y" {

			*token_type = "coordinate"

		} else if token_value == "width" ||
					token_value == "height" {

			*token_type = "dimension"

		}
		
	} else if *token_type == "operator" {

	// check for assignment operators

		if token_value == "=" {

			*token_type = "assigner"

		}

	} else if *token_type == "container" {

	// check direction of containers

		if token_value == "(" {

			*token_type = "container_left"

		} else if token_value == ")" {

			*token_type = "container_right"

		}

	}

}