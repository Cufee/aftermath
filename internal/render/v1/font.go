package render

import (
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
)

type Font struct {
	size float64
	data []byte
}

func (f *Font) Size() float64 {
	return f.size
}

func (f *Font) Valid() bool {
	return f.data != nil && f.size > 0
}

func (f *Font) Face() (font.Face, func() error) {
	ttf, _ := truetype.Parse(f.data)
	face := truetype.NewFace(ttf, &truetype.Options{
		Size: f.size,
	})
	return face, face.Close
}
