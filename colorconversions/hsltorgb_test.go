package asciiart

import "testing"

func TestConvertHSLToRGB(t *testing.T) {
	type args struct {
		h float64
		s float64
		l float64
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 int
		want2 int
	}{
		{name: "Red", args: args{h: 0, s: 100, l: 50}, want: 255, want1: 0, want2: 0},
		{name: "Green", args: args{h: 120, s: 100, l: 50}, want: 0, want1: 255, want2: 0},
		{name: "Blue", args: args{h: 240, s: 100, l: 50}, want: 0, want1: 0, want2: 255},
		{name: "White", args: args{h: 0, s: 0, l: 100}, want: 255, want1: 255, want2: 255},
		{name: "Black", args: args{h: 0, s: 0, l: 0}, want: 0, want1: 0, want2: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2 := ConvertHSLToRGB(tt.args.h, tt.args.s, tt.args.l)
			if got != tt.want {
				t.Errorf("ConvertHSLToRGB() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("ConvertHSLToRGB() got1 = %v, want %v", got1, tt.want1)
			}
			if got2 != tt.want2 {
				t.Errorf("ConvertHSLToRGB() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}
