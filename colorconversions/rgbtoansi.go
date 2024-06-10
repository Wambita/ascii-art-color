package asciiart

import "fmt"

// convert the normalised rgb values to the nearest ansi 256 color code
func ConvertRGBToAnsi(r, g, b int) string {
	rnorm := r / 51
	gnorm := g / 51
	bnorm := b / 51
	color := 16 + ((36 * rnorm) + (6 * gnorm) + bnorm)
	return fmt.Sprintf("\033[38;5;%dm", color)
}
