package Parser

import (
	"fmt"
	"models"
)

/*
PROGRAM -> STATEMENT | PROGRAM NEWLINE STATEMENT
STATEMENT -> DRAW_COMMAND FUNCTION LOCATION
DRAW_COMMAND -> INSTRUCTION OBJECT
LOCATION -> CONTAINER_LEFT COORD_VAL SEPARATOR COORD_VAL SEPARATOR DIM_VAL SEPARATOR DIM_VAL CONTAINER_RIGHT
COORD_VAL -> COORDINATE ASSIGNER NUMBER
DIM_VAL -> DIMENSION ASSIGNER NUMBER
NUMBER = DIGIT | NUMBER DIGIT
DIGIT -> 0 | 1 | ... | 9
ASSIGNER -> =
DIMENSION -> width | height
COORDINATE -> x | y
SEPARATOR -> ,
CONTAINER_LEFT -> (
CONTAINER_RIGHT -> )
FUNCTION -> at
INSTRUCTION -> paint
OBJECT -> rectangle
NEWLINE -> \n
*/

func Parse(tokens []models.Token) {

	// check for at least one token
	if tokens == nil || len(tokens) < 1 {

		fmt.Println("Error: program cannot be null and must contain at least one instruction")

	}

	fmt.Println("Validating your program with", len(tokens), "instructions...")

	// current token number
	current_token_number := 0

	if program(tokens, &current_token_number) {

		fmt.Println("Your program has been successfully validated!")

	}

}

// validate a program
func program(tokens []models.Token, current_token_number *int) bool {

	// there are no more tokens to parse, and we 
	// are successfully finished
	if isNoMoreTokens(tokens, current_token_number) {

		return true

	}

	// if we don't encounter a statement, then we 
	// throw a compile error
	if !statement(tokens, current_token_number) {

		fmt.Println("There is an error in your program on line", tokens[*current_token_number].Line_num, "character number", tokens[*current_token_number].Char_num, "token number", tokens[*current_token_number].Token_num, ":\nUnexpected", tokens[*current_token_number].Token_value)
		return false

	}

	// if we encounter a newline after a statement, check
	// for another statement
	if newline(tokens, current_token_number) {

		return program(tokens, current_token_number)

	}

	return true

}

func statement(tokens []models.Token, current_token_number *int) bool {

	// if a statement doesn't begin with a draw command,
	// throw a compile error
	if !draw_command(tokens, current_token_number) {

		return false

	}

	// if a draw command is not followed by a function,
	// throw a compile error
	if !function(tokens, current_token_number) {

		return false

	}

	// if a function is not followed by a location, 
	// throw a compile error
	if !location(tokens, current_token_number) {

		return false

	}

	return true

}

func newline(tokens []models.Token, current_token_number *int) bool {

	// if the token is null, then this is not a newline
	if isNoMoreTokens(tokens, current_token_number) {

		return false

	}

	if tokens[*current_token_number].Token_type != "newline" {

		return false

	}

	// move to the next token
	*current_token_number += 1
	return true

}

func draw_command(tokens []models.Token, current_token_number *int) bool {

	// if a draw_command does not start with an instruction,
	// throw a compile error
	if !instruction(tokens, current_token_number) {

		return false

	}

	// if an instruction is not followed by an object,
	// throw a compile error
	if !object(tokens, current_token_number) {

		return false

	}

	return true

}

func instruction(tokens []models.Token, current_token_number *int) bool {

	// if the token isn't an instruction, return false
	if tokens[*current_token_number].Token_type != "instruction" {

		return false

	}

	*current_token_number += 1
	return true

}

func object(tokens []models.Token, current_token_number *int) bool {

	if tokens[*current_token_number].Token_type != "object" {

		return false

	}

	*current_token_number += 1
	return true

}

func function(tokens []models.Token, current_token_number *int) bool {

	if tokens[*current_token_number].Token_type != "function" {

		return false

	}

	*current_token_number += 1
	return true

}

func location(tokens []models.Token, current_token_number *int) bool {

	if !container_left(tokens, current_token_number) {

		return false

	}

	if !coord_val(tokens, current_token_number) {

		return false

	}

	if !separator(tokens, current_token_number) {

		return false

	}

	if !coord_val(tokens, current_token_number) {

		return false

	}

	if !separator(tokens, current_token_number) {

		return false

	}

	if !dim_val(tokens, current_token_number) {

		return false

	}

	if !separator(tokens, current_token_number) {

		return false

	}

	if !dim_val(tokens, current_token_number) {

		return false

	}

	if !container_right(tokens, current_token_number) {

		return false

	}

	return true

}

func container_left(tokens []models.Token, current_token_number *int) bool {

	if tokens[*current_token_number].Token_type != "container_left" {

		return false

	}

	*current_token_number += 1
	return true

}

func container_right(tokens []models.Token, current_token_number *int) bool {

	if tokens[*current_token_number].Token_type != "container_right" {

		return false

	}

	*current_token_number += 1
	return true

}

func coord_val(tokens []models.Token, current_token_number *int) bool {

	// fmt.Println(tokens[*current_token_number].Token_type, tokens[*current_token_number].Token_value, *current_token_number)

	if !coordinate(tokens, current_token_number) {

		return false

	}

	// fmt.Println(tokens[*current_token_number].Token_type, tokens[*current_token_number].Token_value, *current_token_number)

	if !assigner(tokens, current_token_number) {

		return false

	}

	// fmt.Println(tokens[*current_token_number].Token_type, tokens[*current_token_number].Token_value, *current_token_number)

	if !number(tokens, current_token_number) {

		return false

	}

	return true

}

func dim_val(tokens []models.Token, current_token_number *int) bool {

	if !dimension(tokens, current_token_number) {

		return false

	}

	if !assigner(tokens, current_token_number) {

		return false

	}

	if !number(tokens, current_token_number) {

		return false

	}

	return true

}

func separator(tokens []models.Token, current_token_number *int) bool {

	if tokens[*current_token_number].Token_type != "separator" {

		return false

	}

	*current_token_number += 1
	return true

}

func coordinate(tokens []models.Token, current_token_number *int) bool {

	if tokens[*current_token_number].Token_type != "coordinate" {

		return false

	}

	*current_token_number += 1
	return true

}

func assigner(tokens []models.Token, current_token_number *int) bool {

	if tokens[*current_token_number].Token_type != "assigner" {

		return false

	}

	*current_token_number += 1
	return true

}

func number(tokens []models.Token, current_token_number *int) bool {

	if tokens[*current_token_number].Token_type != "number" {

		return false

	}

	*current_token_number += 1
	return true

}

func dimension(tokens []models.Token, current_token_number *int) bool {

	if tokens[*current_token_number].Token_type != "dimension" {

		return false

	}

	*current_token_number += 1
	return true

}

func isNoMoreTokens(tokens []models.Token, current_token_number *int) bool {

	if *current_token_number > len(tokens) - 1 {

		return true

	}

	return false

}