package common

import (
	"context"
	"os"
	"slices"
	"strconv"

	"image"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/cufee/aftermath/internal/constants"
	"github.com/cufee/aftermath/internal/database"
	"github.com/cufee/aftermath/internal/database/models"
	"github.com/cufee/aftermath/internal/encoding"
	"github.com/cufee/aftermath/internal/log"
	"github.com/nao1215/imaging"
	"github.com/pkg/errors"
)

var ErrInvalidImage = errors.New("invalid image")
var ErrUnsupportedImageFormat = errors.New("unsupported image format")

func CreateImageModerationRequest(ctx context.Context, db database.ModerationClient, userID string, imageURL *url.URL) (image.Image, string, error) {
	img, err := SafeLoadRemoteImage(imageURL)
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

// https://stackoverflow.com/questions/63464736/image-maximum-file-size-based-on-image-dimension

var remoteImageClient = http.Client{
	Timeout: time.Second * 1,
}

func SafeLoadRemoteImage(url *url.URL) (image.Image, error) {
	res, err := remoteImageClient.Get(url.String())
	if err != nil {
		if os.IsTimeout(err) {
			return nil, ErrInvalidImage
		}
		return nil, err
	}
	defer res.Body.Close()

	contentLengthStr := res.Header.Get("Content-Length")
	if contentLengthStr == "" {
		log.Warn().Msg("Content-Length header is missing on remote replay file")
		return nil, ErrInvalidImage
	}
	contentLength, err := strconv.ParseInt(contentLengthStr, 10, 64)
	if err != nil {
		return nil, err
	}
	if contentLength > constants.ImageUploadMaxSize {
		return nil, ErrInvalidImage
	}

	img, format, err := image.Decode(io.LimitReader(res.Body, constants.ImageUploadMaxSize))
	if err != nil {
		if errors.Is(err, image.ErrFormat) {
			return nil, ErrUnsupportedImageFormat
		}
		return nil, err
	}

	if !slices.Contains([]string{"jpeg", "png"}, format) {
		return nil, ErrUnsupportedImageFormat
	}

	return img, nil
}
