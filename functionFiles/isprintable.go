package asciiart

import (
	"fmt"
	"strings"
)

func IsNotPrintable(str string) bool {
	NonPrintableChars := []string{"\\a", "\\b", "\\f", "\\n", "\\r", "\\t", "\\v", "\\f"}

	for _, char := range NonPrintableChars {
		if contains := strings.Contains(str, char); contains {
			fmt.Println("Error: input contains non-Printable character: ", char)
			return true
		}
	}
	return false
}
