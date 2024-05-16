package asciiart

// import (
// 	"bytes"
// 	"io"
// 	"os"
// 	"testing"
// )

// func TestDisplayAsciiArt(t *testing.T) {
// 	testMap := map[rune][]string{
// 		'A': {"1", "2", "3", "4", "5", "6", "7", "8"},
// 		'B': {"9", "10", "11", "13", "14", "15", "16", "17"},
// 	}
// 	tests := []struct {
// 		name   string
// 		input  string
// 		output string
// 	}{
// 		{
// 			name:   "Test two chracters",
// 			input:  "AB",
// 			output: "1 9 \n2 10 \n3 11 \n4 13 \n5 14 \n6 15 \n7 16 \n8 17 \n",
// 		},
// 	}
// 	for _, tc := range tests {
// 		t.Run(tc.name, func(t *testing.T) {
// 			// Capture standard output
// 			old := os.Stdout
// 			r, w, _ := os.Pipe()
// 			os.Stdout = w
// 			DisplayAsciiArt(testMap, tc.input)
// 			w.Close()
// 			var buf bytes.Buffer
// 			io.Copy(&buf, r)
// 			os.Stdout = old
// 			output := buf.String()
// 			if output != tc.output {
// 				t.Errorf("Expected output:\n%s\nBut got:\n%s", tc.output, output)
// 			}
// 		})
// 	}
// }
