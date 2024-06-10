package asciiart

import (
	"fmt"
	"strconv"
	"strings"
)

// Extracting rgb values from input rgb string
func ExtractRGBValues(rgb string) (int, int, int, error) {
	if !(strings.HasPrefix(rgb, "rgb(")) && !(strings.HasSuffix(rgb, ")")) {
		return 0, 0, 0, fmt.Errorf("invalid rgb format\n\n use rgb(0, 0, 0)")
	}
	rgb = strings.TrimPrefix(rgb, "rgb(")
	rgb = strings.TrimSuffix(rgb, ")")
	rgbValues := strings.Split(rgb, ",")
	if len(rgbValues) != 3 {
		return 0, 0, 0, fmt.Errorf("invalid rgb format\n\n: Use 3 rgb values")
	}
	// convert rgbValues to int values
	r, err1 := strconv.Atoi(strings.TrimSpace(rgbValues[0]))
	g, err2 := strconv.Atoi(strings.TrimSpace(rgbValues[1]))
	b, err3 := strconv.Atoi(strings.TrimSpace(rgbValues[2]))

	if err1 != nil || err2 != nil || err3 != nil {
		return 0, 0, 0, fmt.Errorf("invalid rgb format")
	}
	// rgb values should range between 0-255
	if r < 0 || r > 255 {
		return 0, 0, 0, fmt.Errorf("red (r) must range between 0 -255")
	}
	if g < 0 || g > 255 {
		return 0, 0, 0, fmt.Errorf("green (g) must range between 0 -255")
	}
	if b < 0 || b > 255 {
		return 0, 0, 0, fmt.Errorf("blue (b) must range between 0 -255")
	}
	return r, g, b, nil
}
