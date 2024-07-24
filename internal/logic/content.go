package logic

import (
	"context"
	"errors"
	"image"

	"github.com/cufee/aftermath/internal/database"
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

func GetBackgroundImageFromRef(ctx context.Context, db database.Client, referenceID string) (image.Image, models.UserContent, error) {
	content, err := db.GetUserContentFromRef(ctx, referenceID, models.UserContentTypePersonalBackground)
	if err != nil {
		return nil, models.UserContent{}, err
	}
	img, err := UserContentToImage(content)
	if err != nil {
		return nil, models.UserContent{}, err
	}

	return img, content, nil
}
