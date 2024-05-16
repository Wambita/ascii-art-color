package asciiart

import (
	"testing"
)

func TestIsPrintable(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		expeted bool
	}{
		{name: "printable string", input: "heLLo (world)!", expeted: true},
		{name: "Non-printable string", input: "heLLo\t(world)!", expeted: false},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := IsPrintable(test.input)
			if result != test.expeted{
				t.Errorf("IsPrintable(%s) expected %v, got %v ", test.input, test.expeted, result)
			}
		})
	}
}
