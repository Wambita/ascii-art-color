package asciiart

import "strconv"

// Converts Hex color code input to nearest ANSI color
func ConvertHexToAnsi(hex string) string {
	if len(hex) != 7 && hex[0] != '#' {
		return ""
	}

	r, _ := strconv.ParseInt(hex[1:3], 16, 64)
	g, _ := strconv.ParseInt(hex[3:5], 16, 64)
	b, _ := strconv.ParseInt(hex[5:7], 16, 64)

	return ConvertRGBToAnsi(int(r), int(g), int(b))
}
