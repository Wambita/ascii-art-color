package asciiart

import (
	"testing"
)

const (
	standardHash   = 0x9ffd59bc
	shadowHash    = 0x2f465361
	thinkertoyHash = 0x6ee86a07
	testFileHash  = 0x597778CF
)

func TestIsAltered(t *testing.T) {
	tests := []struct {
		name          string
		generatedHash uint32
		want          bool
	}{
		{
			name:          "Valid standard file hash",
			generatedHash: standardHash,
			want:          false,
		},
		{
			name:          "Valid shadow file hash",
			generatedHash: shadowHash,
			want:          false,
		},
		{
			name:          "Valid thinkertoy file hash",
			generatedHash: thinkertoyHash,
			want:          false,
		},
		{
			name:          "Valid test file hash",
			generatedHash: testFileHash,
			want:          false,
		},
		{
			name:          "Altered data - invalid checksum",
			generatedHash: 0xFFFFFFFF, 
			want:          true,
		},
		{
			name:          "Random hash - invalid checksum",
			generatedHash: 0x12345678, // Random hash not matching any predefined values
			want:          true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsAltered(tt.generatedHash); got != tt.want {
				t.Errorf("IsAltered() = %v, want %v", got, tt.want)
			}
		})
	}
}
