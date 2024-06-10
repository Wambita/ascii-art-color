package asciiart

import (
	"fmt"
	"strings"
)

func IsNotPrintable(str string) bool {
	NonPrintableChars := []string{"\\a", "\\b", "\\f", "\\r", "\\t", "\\v", "\f", "\t", "\a,", "\b", "\f", "\r", "\v"}
	for _, char := range NonPrintableChars {
		if contains := strings.Contains(str, char); contains {
			fmt.Println("Error: input contains non-Printable character:", char)
			return true
		}
	}
	return false
}
