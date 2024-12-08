package images

import (
	"errors"
	"image"
	"io"
	"net/http"
	"net/url"
	"os"
	"slices"
	"strconv"
	"time"

	"github.com/cufee/aftermath/internal/log"
)

var ErrInvalidImage = errors.New("invalid image")
var ErrUnsupportedImageFormat = errors.New("unsupported image format")

var remoteImageClient = http.Client{
	Timeout: time.Second * 1,
}

func SafeLoadFromURL(url *url.URL, maxSize int64) (image.Image, error) {
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
		log.Warn().Msg("Content-Length header is missing on remote file")
		return nil, ErrInvalidImage
	}
	contentLength, err := strconv.ParseInt(contentLengthStr, 10, 64)
	if err != nil {
		return nil, err
	}
	if contentLength > maxSize {
		return nil, ErrInvalidImage
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
