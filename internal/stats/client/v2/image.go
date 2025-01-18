package client

import (
	"image"
	"image/png"
	"io"

	"github.com/pkg/errors"
)

type imageImp struct {
	image.Image
}

func (i *imageImp) PNG(w io.Writer) error {
	if i.Image == nil {
		return errors.New("image cannot be nil")
	}
	return png.Encode(w, i.Image)
}
