package asciiart

import (
	"strings"
)

func IsNotPrintable(str string) bool {
	NonPrintableChars := []string{"\\a", "\\b", "\\f", "\\r", "\\t", "\\v", "\f", "\t", "\a,", "\b", "\f", "\r", "\v"}
	for _, char := range NonPrintableChars {
		if contains := strings.Contains(str, char); contains {
			return true
		}
	}
	return false
}
