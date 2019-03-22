package models

type token struct {
	token_type string
	token_value interface{}
	next *node
	line_num int
	char_num int
	token_num int
}