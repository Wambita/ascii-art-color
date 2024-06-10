package asciiart

import "fmt"

// convert the normalised rgb values to the nearest ansi 256 color code
func ConvertRGBToAnsi(r, g, b int) string {
<<<<<<< HEAD
	rnorm := r / 51
	gnorm := g / 51
	bnorm := b / 51
	color := 16 + ((36 * rnorm) + (6 * gnorm) + bnorm)
=======
	rnormal := r / 51
	gnormal := g / 51
	bnormal := b / 51
	color := 16 + ((36 * rnormal) + (6 * gnormal) + bnormal)
>>>>>>> 55cb84d (fix: chnage variable names)
	return fmt.Sprintf("\033[38;5;%dm", color)
}
