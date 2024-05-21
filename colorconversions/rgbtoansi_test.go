package asciiart

import "testing"

func TestConvertRGBToAnsi(t *testing.T) {
	type args struct {
		r int
		g int
		b int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "Black", args: args{r: 0, g: 0, b: 0}, want: "\033[38;5;16m"},
		{name: "White", args: args{r: 255, g: 255, b: 255}, want: "\033[38;5;231m"},
		{name: "Red", args: args{r: 255, g: 0, b: 0}, want: "\033[38;5;196m"},
		{name: "Green", args: args{r: 0, g: 255, b: 0}, want: "\033[38;5;46m"},
		{name: "Blue", args: args{r: 0, g: 0, b: 255}, want: "\033[38;5;21m"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConvertRGBToAnsi(tt.args.r, tt.args.g, tt.args.b); got != tt.want {
				t.Errorf("ConvertRGBToAnsi() = %v, want %v", got, tt.want)
			}
		})
	}
}
