package assets

import (
	"bytes"
	"embed"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"path/filepath"
	"strings"

	"github.com/cufee/aftermath/internal/files"
	"github.com/cufee/aftermath/internal/log"
	"github.com/golang/freetype/truetype"
)

//go:embed *
var assets embed.FS

var fontsMap = make(map[string][]byte, 0)
var imagesMap = make(map[string]image.Image, 0)

func Load() error {
	dirFiles, err := files.ReadDirFiles(assets, ".")
	if err != nil {
		return err
	}

	fontsMap, err = loadFonts(dirFiles)
	if err != nil {
		return err
	}

	imagesMap, err = loadImages(dirFiles)
	if err != nil {
		return err
	}

	return nil
}

func loadFonts(files map[string][]byte) (map[string][]byte, error) {
	fonts := make(map[string][]byte)
	for path, data := range files {
		if !strings.HasSuffix(path, ".ttf") {
			continue
		}

		_, err := truetype.Parse(data)
		if err != nil {
			return nil, err
		}

		fonts[strings.Split(filepath.Base(path), ".")[0]] = data
		log.Debug().Str("path", path).Msg("loaded font")
	}

	return fonts, nil
}

func loadImages(files map[string][]byte) (map[string]image.Image, error) {
	for path, file := range files {
		var decoder func(r io.Reader) (image.Image, error)
		if strings.HasSuffix(path, ".png") {
			decoder = png.Decode
		} else if strings.HasSuffix(path, ".jpg") || strings.HasSuffix(path, ".jpeg") {
			decoder = jpeg.Decode
		} else {
			continue
		}

		image, err := decoder(bytes.NewBuffer(file))
		if err != nil {
			return nil, err
		}
		key := strings.Split(filepath.Base(path), ".")[0]
		imagesMap[key] = image
		log.Debug().Str("path", path).Str("key", key).Msg("loaded image")
	}

	return imagesMap, nil
}

func FontFace(name string) ([]byte, bool) {
	f, ok := fontsMap[name]
	return f, ok
}

func Image(name string) (image.Image, bool) {
	img, ok := imagesMap[name]
	return img, ok
}
