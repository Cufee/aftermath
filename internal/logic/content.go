package logic

import (
	"errors"
	"image"

	"github.com/cufee/aftermath/internal/database/models"
	"github.com/cufee/aftermath/internal/encoding"
)

func UserContentToImage(content models.UserContent) (image.Image, error) {
	if content.Value == "" {
		return nil, errors.New("value is empty")
	}
	var img image.NRGBA
	err := encoding.DecodeGob([]byte(content.Value), &img)
	if err != nil {
		return nil, err
	}
	return &img, nil
}

func ImageToUserContentValue(img image.Image) ([]byte, error) {
	if img == nil {
		return nil, errors.New("image is nil")
	}
	encoded, err := encoding.EncodeGob(img)
	if err != nil {
		return nil, err
	}
	return encoded, nil
}
