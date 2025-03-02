package dicebear

import (
	"testing"
)

func TestCustomErrors(t *testing.T) {
	tests := []struct {
		name     string
		err      error
		expected string
	}{
		{
			name:     "InvalidStyleError",
			err:      &InvalidStyleError{Style: "invalid-style"},
			expected: "invalid style: invalid-style",
		},
		{
			name:     "InvalidColorError",
			err:      &InvalidColorError{Code: "invalid-color"},
			expected: "invalid color: invalid-color",
		},
		{
			name:     "ImageFormatError",
			err:      &ImageFormatError{Format: "invalid-format"},
			expected: "invalid format: invalid-format",
		},
		{
			name:     "HTTPResponseError",
			err:      &HTTPResponseError{StatusCode: 404, Body: "Not Found"},
			expected: "HTTP error: status 404, body: Not Found",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.err.Error() != tt.expected {
				t.Errorf("expected error message %q, got %q", tt.expected, tt.err.Error())
			}
		})
	}
}
