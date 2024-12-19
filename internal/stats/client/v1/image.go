package client

import (
	"image"
	"image/png"
	"io"

	"github.com/pkg/errors"

	"github.com/cufee/aftermath/internal/render/v1"
)

type Image interface {
	AddBackground(image.Image, render.Style) error
	PNG(io.Writer) error
}

type imageImp struct {
	image.Image
}

func (i *imageImp) AddBackground(bg image.Image, style render.Style) error {
	if bg == nil {
		return errors.New("background cannot be nil")
	}
	i.Image = render.AddBackground(i, bg, style)
	return nil
}

func (i *imageImp) PNG(w io.Writer) error {
	if i.Image == nil {
		return errors.New("image cannot be nil")
	}
	return png.Encode(w, i.Image)
}
