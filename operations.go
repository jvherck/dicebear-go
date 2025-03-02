/*
MIT License

Copyright (c) 2025 jvherck (on GitHub)

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

package dicebear

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

// GetBytes returns the avatar bytes in the specified format.
func (a *Avatar) GetBytes(format Format) ([]byte, error) {
	// Check if the format is valid
	valid := false

	for _, f := range AllFormats {
		if f == format {
			valid = true
			break
		}
	}

	if !valid {
		return nil, &ImageFormatError{Format: format}
	}

	// Get the URL with the specified format
	formatURL, err := a.URLWithFormat(format)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrNetworkRequest, err)
	}

	// Make the request
	resp, err := a.client.Get(formatURL)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrNetworkRequest, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, &HTTPResponseError{StatusCode: resp.StatusCode, Body: string(body)}
	}

	// Read the response body
	return io.ReadAll(resp.Body)
}

// GetText returns the avatar text in the specified format.
func (a *Avatar) GetText(format Format) (string, error) {
	bytes, err := a.GetBytes(format)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

// Save saves the avatar to disk.
func (a *Avatar) Save(format Format, outputFilepath string, overwrite bool) (string, error) {
	// Check if the format is valid
	valid := false

	for _, f := range AllFormats {
		if f == format {
			valid = true
			break
		}
	}

	if !valid {
		return "", &ImageFormatError{Format: format}
	}

	// Get the image bytes
	imageBytes, err := a.GetBytes(format)
	if err != nil {
		return "", err
	}

	if !strings.HasSuffix(outputFilepath, "."+strings.ToLower(string(format))) {
		outputFilepath += "." + strings.ToLower(string(format))
	}

	// Ensure the output directory exists
	dir := filepath.Dir(outputFilepath)
	if dir != "." {
		err = os.MkdirAll(dir, 0o755)
		if err != nil {
			return "", fmt.Errorf("%w: %v", ErrFileOperationFailed, err)
		}
	}

	// If not overwriting and file exists, find a unique name
	if !overwrite {
		outputFilepath = a.uniquifyPath(outputFilepath)
	}

	// Write the file
	var file *os.File
	if format == SVG || format == JSON {
		// Text-based formats
		file, err = os.Create(outputFilepath)
		if err != nil {
			return "", fmt.Errorf("%w: %v", ErrImageSave, err)
		}
		defer file.Close()

		_, err = file.WriteString(string(imageBytes))
		if err != nil {
			return "", fmt.Errorf("%w: %v", ErrImageSave, err)
		}
	} else {
		// Binary formats
		file, err = os.Create(outputFilepath)
		if err != nil {
			return "", fmt.Errorf("%w: %v", ErrImageSave, err)
		}
		defer file.Close()

		_, err = io.Copy(file, bytes.NewReader(imageBytes))
		if err != nil {
			return "", fmt.Errorf("%w: %v", ErrImageSave, err)
		}
	}

	return outputFilepath, nil
}

// uniquifyPath ensures a file path is unique by adding a counter if necessary.
func (a *Avatar) uniquifyPath(path string) string {
	dir := filepath.Dir(path)
	ext := filepath.Ext(path)
	filename := filepath.Base(path[:len(path)-len(ext)])

	counter := 0
	newPath := path

	for {
		_, err := os.Stat(newPath)
		if os.IsNotExist(err) {
			break
		}

		counter++
		newPath = filepath.Join(dir, fmt.Sprintf("%s(%d)%s", filename, counter, ext))
	}

	return newPath
}

// View opens the avatar in the default web browser.
func (a *Avatar) View(format Format) error {
	url, err := a.URLWithFormat(format)
	if err != nil {
		return err
	}

	// This is platform-dependent and might not work on all systems
	var cmd string
	var args []string

	switch {
	case os.Getenv("BROWSER") != "":
		// Use the BROWSER environment variable
		cmd = os.Getenv("BROWSER")
		args = []string{url}
	default:
		// Platform specific defaults
		switch runtime.GOOS {
		case "darwin":
			cmd = "open"
			args = []string{url}
		case "windows":
			cmd = "rundll32"
			args = []string{"url.dll,FileProtocolHandler", url}
		default: // linux, bsd, etc.
			cmd = "xdg-open"
			args = []string{url}
		}
	}

	return exec.Command(cmd, args...).Start()
}
