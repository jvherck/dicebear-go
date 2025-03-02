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

const (
	BaseURL = "https://api.dicebear.com/9.x" // Base URL for DiceBear API
	Timeout = 30                             // Default timeout in seconds
)

var baseURL = BaseURL

func resetURL() {
	baseURL = BaseURL
}

// Style represents the available avatar styles.
type Style string

const (
	Adventurer        Style = "adventurer"         // Friendly adventurer style
	AdventurerNeutral Style = "adventurer-neutral" // Neutral adventurer style
	Avataaars         Style = "avataaars"          // Avatar style inspired by avataaars
	AvataaarsNeutral  Style = "avataaars-neutral"  // Neutral avataaars style
	BigEars           Style = "big-ears"           // Big ears style
	BigEarsNeutral    Style = "big-ears-neutral"   // Neutral big ears style
	BigSmile          Style = "big-smile"          // Big smile style
	Bottts            Style = "bottts"             // Bottts style
	BotttsNeutral     Style = "bottts-neutral"     // Neutral bottts style
	Croodles          Style = "croodles"           // Croodles style
	CroodlesNeutral   Style = "croodles-neutral"   // Neutral croodles style
	Dylan             Style = "dylan"              // Dylan style
	FunEmoji          Style = "fun-emoji"          // Fun emoji style
	Glass             Style = "glass"              // Glass style
	Icons             Style = "icons"              // Icons style
	Identicon         Style = "identicon"          // Identicon style
	Initials          Style = "initials"           // Initials style
	Lorelei           Style = "lorelei"            // Lorelei style
	LoreleiNeutral    Style = "lorelei-neutral"    // Neutral lorelei style
	Micah             Style = "micah"              // Micah style
	Miniavs           Style = "miniavs"            // Miniavs style
	Notionists        Style = "notionists"         // Notionists style
	NotionistsNeutral Style = "notionists-neutral" // Neutral notionists style
	OpenPeeps         Style = "open-peeps"         // Open peeps style
	Personas          Style = "personas"           // Personas style
	PixelArt          Style = "pixel-art"          // Pixel art style
	PixelArtNeutral   Style = "pixel-art-neutral"  // Neutral pixel art style
	Rings             Style = "rings"              // Rings style
	Shapes            Style = "shapes"             // Shapes style
	Thumbs            Style = "thumbs"             // Thumbs style
)

// AllStyles is a list of all available avatar styles.
var AllStyles = []Style{
	Adventurer, AdventurerNeutral, Avataaars, AvataaarsNeutral,
	BigEars, BigEarsNeutral, BigSmile, Bottts, BotttsNeutral,
	Croodles, CroodlesNeutral, Dylan, FunEmoji, Glass, Icons,
	Identicon, Initials, Lorelei, LoreleiNeutral, Micah, Miniavs,
	Notionists, NotionistsNeutral, OpenPeeps, Personas,
	PixelArt, PixelArtNeutral, Rings, Shapes, Thumbs,
}

// Format represents the available image formats.
type Format string

const (
	SVG  Format = "svg"  // Scalable Vector Graphics
	WEBP Format = "webp" // WebP format
	AVIF Format = "avif" // AVIF format
	PNG  Format = "png"  // Portable Network Graphics
	JPG  Format = "jpg"  // JPEG format
	JPEG Format = "jpeg" // JPEG format (alias)
	JSON Format = "json" // JSON format
)

// AllFormats is a list of all available image formats.
var AllFormats = []Format{SVG, WEBP, AVIF, PNG, JPG, JPEG, JSON}

// IsValidFormat checks if a format is valid.
func IsValidFormat(format Format) bool {
	for _, f := range AllFormats {
		if f == format {
			return true
		}
	}
	return false
}
