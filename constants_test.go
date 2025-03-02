package dicebear

import (
	"testing"
)

func TestIsValidFormat(t *testing.T) {
	tests := []struct {
		format   Format
		expected bool
	}{
		{SVG, true},
		{"invalid-format", false},
	}

	for _, tt := range tests {
		t.Run(string(tt.format), func(t *testing.T) {
			if IsValidFormat(tt.format) != tt.expected {
				t.Errorf("IsValidFormat() = %v, expected %v", IsValidFormat(tt.format), tt.expected)
			}
		})
	}
}
