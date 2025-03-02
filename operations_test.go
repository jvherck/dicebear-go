package dicebear

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestAvatar_GetBytes(t *testing.T) {
	// Mock HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("test-bytes"))
	}))
	defer server.Close()

	// Override BaseURL for testing
	baseURL = server.URL

	avatar, err := NewAvatar(Adventurer, "test-seed", nil, nil)
	if err != nil {
		t.Fatalf("Failed to create avatar: %v", err)
	}

	bytes, err := avatar.GetBytes(SVG)
	if err != nil {
		t.Errorf("GetBytes() error = %v", err)
	}
	if string(bytes) != "test-bytes" {
		t.Errorf("GetBytes() = %v, expected %v", string(bytes), "test-bytes")
	}

	resetURL()
}

func TestAvatar_Save(t *testing.T) {
	// Mock HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("test-bytes"))
	}))
	defer server.Close()

	// Override BaseURL for testing
	baseURL = server.URL

	avatar, err := NewAvatar(Adventurer, "test-seed", nil, nil)
	if err != nil {
		t.Fatalf("Failed to create avatar: %v", err)
	}

	outputPath := "test-output.svg"
	defer os.Remove(outputPath) // Clean up after test

	_, err = avatar.Save(SVG, outputPath, false)
	if err != nil {
		t.Errorf("Save() error = %v", err)
	}

	// Check if file exists
	if _, err := os.Stat(outputPath); os.IsNotExist(err) {
		t.Errorf("Save() failed to create file %v", outputPath)
	}

	resetURL()
}
