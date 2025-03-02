package dicebear

import (
	"testing"
)

func TestNewColor(t *testing.T) {
	tests := []struct {
		name     string
		hexCode  string
		expected Color
		wantErr  bool
	}{
		{"Valid Hex Code", "#ffffff", Color("ffffff"), false},
		{"Valid Hex Code Without #", "ffffff", Color("ffffff"), false},
		{"Transparent", "transparent", Color("transparent"), false},
		{"Invalid Hex Code", "zzzzzz", "", true},
		{"Invalid Length", "ffff", "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			color, err := NewColor(tt.hexCode)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewColor() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if color != tt.expected {
				t.Errorf("NewColor() = %v, expected %v", color, tt.expected)
			}
		})
	}
}

func TestRandomString(t *testing.T) {
	length := 10
	str := RandomString(length)
	if len(str) != length {
		t.Errorf("RandomString() length = %v, expected %v", len(str), length)
	}
}

func TestIsValidStyle(t *testing.T) {
	tests := []struct {
		style    Style
		expected bool
	}{
		{Adventurer, true},
		{"invalid-style", false},
	}

	for _, tt := range tests {
		t.Run(string(tt.style), func(t *testing.T) {
			if IsValidStyle(tt.style) != tt.expected {
				t.Errorf("IsValidStyle() = %v, expected %v", IsValidStyle(tt.style), tt.expected)
			}
		})
	}
}
