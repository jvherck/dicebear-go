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
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// Avatar represents a DiceBear avatar.
type Avatar struct {
	style        Style
	seed         string
	options      *Options
	customParams map[string]string
	client       *http.Client
	url          string
}

// NewAvatar creates a new Avatar with the given style and seed.
func NewAvatar(style Style, seed string, options *Options, customParams map[string]string) (*Avatar, error) {
	if !IsValidStyle(style) {
		return nil, &InvalidStyleError{Style: style}
	}
	if seed == "" {
		seed = RandomString(20)
	}
	if options == nil {
		options = DefaultOptions()
	}
	if customParams == nil {
		customParams = make(map[string]string)
	}

	client := &http.Client{Timeout: 30 * time.Second}
	avatar := &Avatar{
		style:        style,
		seed:         seed,
		options:      options,
		customParams: customParams,
		client:       client,
	}

	if err := avatar.update(); err != nil {
		return nil, err
	}

	return avatar, nil
}

// GetSchema returns the schema for the avatar style.
func (a *Avatar) GetSchema() (map[string]interface{}, error) {
	resp, err := a.client.Get(fmt.Sprintf("%s/%s/schema.json", baseURL, a.style))
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrNetworkRequest, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%s: status code %d", "HTTPResponseError", resp.StatusCode)
	}

	var schema map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&schema)
	if err != nil {
		return nil, fmt.Errorf("failed to decode JSON response: %v", err)
	}

	return schema, nil
}

// Style returns the style of the avatar.
func (a *Avatar) Style() Style {
	return a.style
}

// Seed returns the seed of the avatar.
func (a *Avatar) Seed() string {
	return a.seed
}

// Options returns the options of the avatar.
func (a *Avatar) Options() *Options {
	return a.options
}

// CustomParams returns the custom parameters of the avatar.
func (a *Avatar) CustomParams() map[string]string {
	return a.customParams
}

// URL returns the URL of the avatar.
func (a *Avatar) URL() string {
	return a.url
}

// URLWithFormat returns the URL of the avatar with the specified format.
func (a *Avatar) URLWithFormat(format Format) (string, error) {
	if !IsValidFormat(format) {
		return "", &ImageFormatError{Format: format}
	}

	return strings.Replace(a.url, "/svg?", fmt.Sprintf("/%s?", format), 1), nil
}

// update updates the avatar URL.
func (a *Avatar) update() error {
	params := url.Values{}

	defaultOptions := DefaultOptions()

	// Add options - comparing against default values
	if a.options.Flip != defaultOptions.Flip {
		params.Add("flip", "true")
	}
	if a.options.Rotate != defaultOptions.Rotate {
		params.Add("rotate", fmt.Sprintf("%d", a.options.Rotate))
	}
	if a.options.Scale != defaultOptions.Scale {
		params.Add("scale", fmt.Sprintf("%d", a.options.Scale))
	}
	if a.options.Radius != defaultOptions.Radius {
		params.Add("radius", fmt.Sprintf("%d", a.options.Radius))
	}
	if a.options.Size != defaultOptions.Size {
		params.Add("size", fmt.Sprintf("%d", a.options.Size))
	}
	if a.options.BackgroundColor != defaultOptions.BackgroundColor {
		params.Add("backgroundColor", string(a.options.BackgroundColor))
	}
	if a.options.BackgroundType != defaultOptions.BackgroundType {
		params.Add("backgroundType", a.options.BackgroundType)
	}
	if a.options.BackgroundRotation != defaultOptions.BackgroundRotation {
		params.Add("backgroundRotation", fmt.Sprintf("%d", a.options.BackgroundRotation))
	}
	if a.options.TranslateX != defaultOptions.TranslateX {
		params.Add("translateX", fmt.Sprintf("%d", a.options.TranslateX))
	}
	if a.options.TranslateY != defaultOptions.TranslateY {
		params.Add("translateY", fmt.Sprintf("%d", a.options.TranslateY))
	}
	if a.options.RandomizeIDs != defaultOptions.RandomizeIDs {
		params.Add("randomizeIds", "true")
	}

	// Add custom parameters
	for k, v := range a.customParams {
		params.Add(k, v)
	}

	// Build URL
	requestURL := fmt.Sprintf("%s/%s/svg?seed=%s", baseURL, a.style, url.QueryEscape(a.seed))
	if len(params) > 0 {
		requestURL += "&" + params.Encode()
	}

	// Make a GET request to check if the URL is valid
	resp, err := a.client.Get(requestURL)
	if err != nil {
		return fmt.Errorf("%w: %w", ErrNetworkRequest, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		// Try to read error response
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("%s: status code %d, body: %s", "HTTPResponseError", resp.StatusCode, body)
	}

	a.url = requestURL

	return nil
}

// Edit modifies the avatar parameters and updates the URL.
func (a *Avatar) Edit(style Style, seed string, options *Options) error {
	changed := false

	if style != "" && style != a.style {
		if !IsValidStyle(style) {
			return &InvalidStyleError{Style: style}
		}
		a.style = style
		changed = true
	}

	if seed != "" && seed != a.seed {
		a.seed = seed
		changed = true
	}

	if options != nil {
		a.options = options
		changed = true
	}

	if changed {
		return a.update()
	}

	return nil
}

// Customize sets custom parameters for the avatar.
func (a *Avatar) Customize(params map[string]string, replace bool) error {
	if replace {
		a.customParams = params
	} else {
		for k, v := range params {
			a.customParams[k] = v
		}
	}

	return a.update()
}
