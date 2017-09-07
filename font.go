package blueprint

import (
	"io/ioutil"
	"os"

	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
)

// NewFont will return a new Font
// TODO: Implement a font caching manager
func NewFont(src string, size float64) (fp *Font, err error) {
	var f Font
	if f.font, err = getFont(src); err != nil {
		return
	}

	f.face = getFontFace(f.font, size)
	f.size = size
	fp = &f
	return
}

// Font represents a font
type Font struct {
	size float64
	font *truetype.Font
	face font.Face
}

// Size will return the font size
func (f *Font) Size() float64 {
	return f.size
}

// Font will return the internal font.Font
func (f *Font) Font() *truetype.Font {
	return f.font
}

// Face will return the font face
func (f *Font) Face() font.Face {
	return f.face
}

func getFont(src string) (fp *truetype.Font, err error) {
	var f *os.File
	if f, err = os.Open(src); err != nil {
		return
	}
	defer f.Close()

	var b []byte
	if b, err = ioutil.ReadAll(f); err != nil {
		return
	}

	return truetype.Parse(b)
}

func getFontFace(f *truetype.Font, size float64) font.Face {
	var opts truetype.Options
	opts.Size = size
	opts.GlyphCacheEntries = 1
	return truetype.NewFace(f, &opts)
}
