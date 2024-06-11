package asciiart

import "fmt"

// convert the normalised rgb values to the nearest ansi 256 color code
func ConvertRGBToAnsi(r, g, b int) string {
	rnormal := r / 51
	gnormal := g / 51
	bnormal := b / 51
	color := 16 + ((36 * rnormal) + (6 * gnormal) + bnormal)
	return fmt.Sprintf("\033[38;5;%dm", color)
}
