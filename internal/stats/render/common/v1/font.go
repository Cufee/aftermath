package common

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

func (f *Font) Face() font.Face {
	ttf, _ := truetype.Parse(f.data)
	return truetype.NewFace(ttf, &truetype.Options{
		Size: f.size,
	})
}
