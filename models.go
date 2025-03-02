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
	"math/rand"
	"strings"
	"time"
)

// Color represents a color in hex or "transparent" format.
type Color string

// NewColor validates and creates a new Color from a hex code.
func NewColor(hexCode string) (Color, error) {
	hexCode = strings.TrimPrefix(hexCode, "#")
	if hexCode == "transparent" {
		return Color("transparent"), nil
	}
	if len(hexCode) != 6 {
		return "", &InvalidColorError{Code: hexCode}
	}
	for _, c := range hexCode {
		if !strings.ContainsRune("0123456789abcdef", c) {
			return "", &InvalidColorError{Code: hexCode}
		}
	}
	return Color(hexCode), nil
}

// NewRandomColor generates a random hex color.
func NewRandomColor() Color {
	const hexChars = "0123456789abcdef"
	b := make([]byte, 6)
	for i := range b {
		b[i] = hexChars[rand.Intn(len(hexChars))]
	}
	return Color(string(b))
}

// Options defines avatar customization options.
type Options struct {
	Flip               bool   `json:"flip,omitempty"`
	Rotate             int    `json:"rotate,omitempty"`
	Scale              int    `json:"scale,omitempty"`
	Radius             int    `json:"radius,omitempty"`
	Size               int    `json:"size,omitempty"`
	BackgroundColor    Color  `json:"-"`
	BackgroundType     string `json:"backgroundType,omitempty"`
	BackgroundRotation int    `json:"backgroundRotation,omitempty"`
	TranslateX         int    `json:"translateX,omitempty"`
	TranslateY         int    `json:"translateY,omitempty"`
	RandomizeIDs       bool   `json:"randomizeIds,omitempty"`
}

// DefaultOptions returns the default avatar options.
func DefaultOptions() *Options {
	return &Options{
		Flip:               false,
		Rotate:             0,
		Scale:              100,
		Radius:             0,
		Size:               0,
		BackgroundColor:    Color("transparent"),
		BackgroundType:     "solid",
		BackgroundRotation: 0,
		TranslateX:         0,
		TranslateY:         0,
		RandomizeIDs:       false,
	}
}

// RandomString generates a random string of specified length.
func RandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyz0123456789"
	b := make([]byte, length)

	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}

	return string(b)
}

// IsValidStyle checks if a style is valid.
func IsValidStyle(style Style) bool {
	for _, s := range AllStyles {
		if s == style {
			return true
		}
	}
	return false
}

// RandomStyle returns a random avatar style.
func RandomStyle() Style {
	return AllStyles[rand.Intn(len(AllStyles))]
}

func init() {
	rand.Seed(time.Now().UnixNano())
}
