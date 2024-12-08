package common

import (
	"context"

	"image"
	"net/url"

	"github.com/cufee/aftermath/internal/constants"
	"github.com/cufee/aftermath/internal/database"
	"github.com/cufee/aftermath/internal/database/models"
	"github.com/cufee/aftermath/internal/encoding"
	"github.com/cufee/aftermath/internal/external/images"
	"github.com/nao1215/imaging"
)

func CreateImageModerationRequest(ctx context.Context, db database.ModerationClient, userID string, imageURL *url.URL) (image.Image, string, error) {
	img, err := images.SafeLoadFromURL(imageURL, constants.ImageUploadMaxSize)
	if err != nil {
		return nil, "", err
	}

	img = imaging.Fill(img, 300, 300, imaging.Center, imaging.Linear)
	data, err := encoding.EncodeGob(img)
	if err != nil {
		return nil, "", err
	}

	request := models.ModerationRequest{
		RequestContext: "image moderation request from /fancy for " + userID,
		ActionStatus:   models.ModerationStatusSubmitted,
		Data:           map[string]any{"image": data},
		ReferenceID:    "custom-image-upload",
		RequestorID:    userID,
	}

	request, err = db.CreateModerationRequest(ctx, request)
	if err != nil {
		return nil, "", err
	}

	return img, request.ID, nil
}
