package globals

import(
	"regexp"
)

func IsOperator(char string) bool {

	return char == ":" || char == "=" || char == "+" || char == "-" || char == "*" || char == "/"

}

func IsSeparator(char string) bool {

	return char == ")" || char == "(" || char == ","

}

func ClearBuffer(buf *string) {

	*buf = ""

}

var IsAlpha = regexp.MustCompile(`^[a-zA-z]+$`).MatchString
var IsDigit = regexp.MustCompile(`^[0-9]+$`).MatchString