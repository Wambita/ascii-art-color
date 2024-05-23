package asciiart

import (
	"fmt"
	"strconv"
	"strings"
)

// ExtractHSLValues extracts the h, s, l values from an hsl string.
func ExtractHSLValues(hsl string) (float64, float64, float64, error) {
	if !strings.HasPrefix(hsl, "hsl(") || !strings.HasSuffix(hsl, ")") {
		return 0, 0, 0, fmt.Errorf("invalid hsl format")
	}
	hsl = strings.TrimPrefix(hsl, "hsl(")
	hsl = strings.TrimSuffix(hsl, ")")
	values := strings.Split(hsl, ",")
	if len(values) != 3 {
		return 0, 0, 0, fmt.Errorf("invalid hsl format")
	}
	h, err1 := strconv.ParseFloat(strings.TrimSuffix(strings.TrimSpace(values[0]), "%"), 64)
	s, err2 := strconv.ParseFloat(strings.TrimSuffix(strings.TrimSpace(values[1]), "%"), 64)
	l, err3 := strconv.ParseFloat(strings.TrimSuffix(strings.TrimSpace(values[2]), "%"), 64)
	if err1 != nil || err2 != nil || err3 != nil {
		return 0, 0, 0, fmt.Errorf("invalid hsl values")
	}
	if h < 0 || h > 360 {
		return 0, 0, 0, fmt.Errorf("hue value must be between 0 and 360")
	}
	if s < 0 || s > 100 {
		return 0, 0, 0, fmt.Errorf("saturation value must be between 0 and 100")
	}
	if l < 0 || l > 100 {
		return 0, 0, 0, fmt.Errorf("lightness value must be between 0 and 100")
	}
	return h, s, l, nil
}
