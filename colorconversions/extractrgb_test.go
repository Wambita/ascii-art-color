package asciiart

import "testing"

func TestExtractRGBValues(t *testing.T) {
	type args struct {
		rgb string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		want1   int
		want2   int
		wantErr bool
	}{
		{
			name:    "Valid Input",
			args:    args{rgb: "rgb(123, 45, 67)"},
			want:    123,
			want1:   45,
			want2:   67,
			wantErr: false,
		},
		{
			name:    "Invalid Format",
			args:    args{rgb: "rgb(23-45-67)"},
			want:    0,
			want1:   0,
			want2:   0,
			wantErr: true,
		},
		{
			name:    "Out of Range",
			args:    args{rgb: "rgb(300, 45, 67)"},
			want:    0,
			want1:   0,
			want2:   0,
			wantErr: true,
		},
		{
			name:    "Non-numeric Values",
			args:    args{rgb: "rgb(abc, 45, 67)"},
			want:    0,
			want1:   0,
			want2:   0,
			wantErr: true,
		},
		{
			name:    "Empty Input",
			args:    args{rgb: "rgb()"},
			want:    0,
			want1:   0,
			want2:   0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2, err := ExtractRGBValues(tt.args.rgb)
			if (err != nil) != tt.wantErr {
				t.Errorf("ExtractRGBValues() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ExtractRGBValues() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("ExtractRGBValues() got1 = %v, want %v", got1, tt.want1)
			}
			if got2 != tt.want2 {
				t.Errorf("ExtractRGBValues() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}
