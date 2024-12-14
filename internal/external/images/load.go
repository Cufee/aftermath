package images

import (
	"context"

	"image"
	"io"
	"net/http"
	"net/url"
	"os"
	"slices"
	"strconv"

	"github.com/cufee/aftermath/internal/log"
	"github.com/pkg/errors"
)

var ErrInvalidImage = errors.New("invalid image")
var ErrUnsupportedImageFormat = errors.New("unsupported image format")

var remoteImageClient = http.DefaultClient

func SafeLoadFromURL(ctx context.Context, url *url.URL, maxSize int64) (image.Image, error) {
	req, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		return nil, ErrInvalidImage
	}
	req = req.WithContext(ctx)

	res, err := remoteImageClient.Do(req)
	if err != nil {
		if os.IsTimeout(err) {
			return nil, errors.Wrap(ErrInvalidImage, "request failed")
		}
		return nil, err
	}
	defer res.Body.Close()

	contentLengthStr := res.Header.Get("Content-Length")
	if contentLengthStr == "" {
		log.Warn().Msg("Content-Length header is missing on remote file")
		return nil, errors.Wrap(ErrInvalidImage, "missing headers")
	}
	contentLength, err := strconv.ParseInt(contentLengthStr, 10, 64)
	if err != nil {
		return nil, err
	}
	if contentLength > maxSize {
		return nil, errors.Wrap(ErrInvalidImage, "image too large")
	}

	img, format, err := image.Decode(io.LimitReader(res.Body, maxSize))
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
