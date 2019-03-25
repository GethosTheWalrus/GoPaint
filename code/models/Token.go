package models

type Token struct {
	Token_type string
	Token_value interface{}
	Line_num int
	Char_num int
	Token_num int
}