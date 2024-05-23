package asciiart

import "testing"

func TestExtractHSLValues(t *testing.T) {
	type args struct {
		hsl string
	}
	tests := []struct {
		name    string
		args    args
		want    float64
		want1   float64
		want2   float64
		wantErr bool
	}{
		{
			name:    "Valid Input",
			args:    args{hsl: "hsl(123, 45, 67)"},
			want:    123,
			want1:   45,
			want2:   67,
			wantErr: false,
		},
		{
			name:    "Invalid Format",
			args:    args{hsl: "hsl(23-45-67)"},
			want:    0,
			want1:   0,
			want2:   0,
			wantErr: true,
		},
		{
			name:    "Out of Range",
			args:    args{hsl: "hsl(370, 101, 607)"},
			want:    0,
			want1:   0,
			want2:   0,
			wantErr: true,
		},
		{
			name:    "Non-numeric Values",
			args:    args{hsl: "hsl(abc, 45, 67)"},
			want:    0,
			want1:   0,
			want2:   0,
			wantErr: true,
		},
		{
			name:    "Empty Input",
			args:    args{hsl: "hsl()"},
			want:    0,
			want1:   0,
			want2:   0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2, err := ExtractHSLValues(tt.args.hsl)
			if (err != nil) != tt.wantErr {
				t.Errorf("ExtractHSLValues() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ExtractHSLValues() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("ExtractHSLValues() got1 = %v, want %v", got1, tt.want1)
			}
			if got2 != tt.want2 {
				t.Errorf("ExtractHSLValues() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}
