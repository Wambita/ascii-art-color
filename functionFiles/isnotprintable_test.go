package asciiart

import "testing"

func TestIsNotPrintable(t *testing.T) {
	tests := []struct {
		name string
		str  string
		want bool
	}{
		{name: "Printable input", str: "Hello World", want: false},
		{name: "Non-printable input", str: "Hello\tworld", want: true},
		{name: "Non-printable input", str: "Hello\\tworld", want: true},
		{name: "Non-printable input", str: "Hello\nworld", want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsNotPrintable(tt.str); got != tt.want {
				t.Errorf("IsNotPrintable() = %v, want %v", got, tt.want)
			}
		})
	}
}
