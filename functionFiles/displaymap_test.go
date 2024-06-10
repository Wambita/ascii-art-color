package asciiart

import "testing"

func TestDisplayAsciiArtWithPartialColor(t *testing.T) {
	type args struct {
		characterMap   map[rune][]string
		input          string
		lettersToColor string
		color          string
	}
	tests := []struct {
		name string
		args args
	}{
		
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			DisplayAsciiArtWithPartialColor(tt.args.characterMap, tt.args.input, tt.args.lettersToColor, tt.args.color)
		})
	}
}
