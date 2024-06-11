package asciiart

import (
	"fmt"
	"strconv"
)

// Converts Hex color code input to nearest ANSI color
func ConvertHexToAnsi(hex string) string {
	if len(hex) != 7 || hex[0] != '#' {
		fmt.Println("Invalid hex format: use 6 #RRGGBB values from 0-F")
		return ""
	}
	// validate each character of the hex string
	var (
		r, g, b int64
	)
	for _, char := range hex[1:] {
		if !((char >= '0' && char <= '9') || (char >= 'a' && char <= 'f') || (char >= 'A' && char <= 'F')) {
			fmt.Println("Invalid hex color code use values from 0-9 and A-F")
			return ""
		}

		r, _ = strconv.ParseInt(hex[1:3], 16, 64)
		g, _ = strconv.ParseInt(hex[3:5], 16, 64)
		b, _ = strconv.ParseInt(hex[5:7], 16, 64)
	}
	return ConvertRGBToAnsi(int(r), int(g), int(b))
}
