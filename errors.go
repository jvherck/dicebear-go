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
	"errors"
	"fmt"
)

// Error types with context
type InvalidStyleError struct{ Style Style }

func (e *InvalidStyleError) Error() string { return fmt.Sprintf("invalid style: %s", e.Style) }

type InvalidColorError struct{ Code string }

func (e *InvalidColorError) Error() string { return fmt.Sprintf("invalid color: %s", e.Code) }

type ImageFormatError struct{ Format Format }

func (e *ImageFormatError) Error() string { return fmt.Sprintf("invalid format: %s", e.Format) }

type HTTPResponseError struct {
	StatusCode int
	Body       string
}

func (e *HTTPResponseError) Error() string {
	return fmt.Sprintf("HTTP error: status %d, body: %s", e.StatusCode, e.Body)
}

// Common errors
var (
	ErrInvalidOption       = errors.New("invalid option")
	ErrNetworkRequest      = errors.New("network request failed")
	ErrImageSave           = errors.New("failed to save image")
	ErrFileOperationFailed = errors.New("file operation failed")
)
