package render

import (
	"sync"

	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
)

type Font struct {
	size float64
	face font.Face
	mx   *sync.Mutex
}

func (f *Font) Size() float64 {
	return f.size
}

func (f *Font) Valid() bool {
	return f.face != nil
}

func (f *Font) Face() (font.Face, func() error) {
	f.mx.Lock()
	return f.face, func() error { f.mx.Unlock(); return nil }
}

func NewFont(data []byte, size float64) *Font {
	ttf, _ := truetype.Parse(data)
	face := truetype.NewFace(ttf, &truetype.Options{
		Size: size,
	})
	return &Font{size, face, &sync.Mutex{}}
}
