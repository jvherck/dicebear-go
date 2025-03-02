package dicebear

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewAvatar(t *testing.T) {
	tests := []struct {
		name    string
		style   Style
		seed    string
		wantErr bool
	}{
		{"Valid Style", Adventurer, "test-seed", false},
		{"Invalid Style", "invalid-style", "test-seed", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewAvatar(tt.style, tt.seed, nil, nil)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewAvatar() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAvatar_GetSchema(t *testing.T) {
	// Mock HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"test": "schema"}`))
	}))
	defer server.Close()

	// Override BaseURL for testing
	baseURL = server.URL

	avatar, err := NewAvatar(Adventurer, "test-seed", nil, nil)
	if err != nil {
		t.Fatalf("Failed to create avatar: %v", err)
	}

	schema, err := avatar.GetSchema()
	if err != nil {
		t.Errorf("GetSchema() error = %v", err)
	}
	if schema["test"] != "schema" {
		t.Errorf("GetSchema() = %v, expected %v", schema, map[string]interface{}{"test": "schema"})
	}

	resetURL()
}
