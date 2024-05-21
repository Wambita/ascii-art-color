package asciiart

import "testing"

func TestConvertHexToAnsi(t *testing.T) {
	type args struct {
		hex string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "Black", args: args{hex: "#000000"}, want: "\033[38;5;16m"},
		{name: "White", args: args{hex: "#FFFFFF"}, want: "\033[38;5;231m"},
		{name: "Red", args: args{hex: "#FF0000"}, want: "\033[38;5;196m"},
		{name: "Green", args: args{hex: "#00FF00"}, want: "\033[38;5;46m"},
		{name: "Blue", args: args{hex: "#0000FF"}, want: "\033[38;5;21m"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConvertHexToAnsi(tt.args.hex); got != tt.want {
				t.Errorf("ConvertHexToAnsi() = %v, want %v", got, tt.want)
			}
		})
	}
}
