package asciiart

import (
	"unicode"
)

func IsPrintable(str string) bool {
	// Iterates through each character in the input string
	/* Checks if the current character is a backslash \ and if there's another character after it.
       If so, checks if the next character is one of the escape characters: \t, \a, \b, \v, \f, or \r */
	for i,value := range str {
		if i < len(str)-1 && str[i] == '\\' {
			// Allow backslash but check next character
			if str[i+1] == 't' || str[i+1] == 'a' || str[i+1] == 'b' || str[i+1] == 'v' || str[i+1] == 'f' || str[i+1] == 'r' {
				return false
			}
		}
		if !unicode.IsPrint(value) {
			return false
		}
	}
	return true
}
