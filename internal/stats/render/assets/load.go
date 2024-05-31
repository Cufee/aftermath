package assets

import (
	"bytes"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"io/fs"
	"path/filepath"
	"strings"

	"github.com/golang/freetype/truetype"
	"github.com/rs/zerolog/log"
	"golang.org/x/image/font"
)

var fontsMap = make(map[string]*truetype.Font)
var imagesMap = make(map[string]image.Image)

func LoadAssets(assets fs.FS) error {
	files, err := readAllFiles(assets, ".")
	if err != nil {
		return err
	}

	fontsMap, err = loadFonts(files)
	if err != nil {
		return err
	}

	imagesMap, err = loadImages(files)
	if err != nil {
		return err
	}

	return nil
}

func loadFonts(files map[string][]byte) (map[string]*truetype.Font, error) {
	fonts := make(map[string]*truetype.Font)
	for path, data := range files {
		if !strings.HasSuffix(path, ".ttf") {
			continue
		}

		font, err := truetype.Parse(data)
		if err != nil {
			return nil, err
		}

		fonts[strings.Split(filepath.Base(path), ".")[0]] = font
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
		imagesMap[strings.Split(filepath.Base(path), ".")[0]] = image
		log.Debug().Str("path", path).Msg("loaded image")
	}

	return imagesMap, nil
}

func readAllFiles(dir fs.FS, path string) (map[string][]byte, error) {
	if dir == nil {
		return nil, fs.ErrInvalid
	}

	files := make(map[string][]byte)

	fs.WalkDir(dir, path, func(path string, d fs.DirEntry, err error) error {
		// If we are not able to access a file/directory for some reason, continue
		if err != nil {
			log.Err(err).Str("path", path).Msg("failed to read a file")
			return nil
		}

		if d.IsDir() {
			return nil
		}

		data, err := fs.ReadFile(dir, path)
		if err != nil {
			return err
		}
		files[path] = data
		return nil
	})

	return files, nil
}

func GetLoadedFontFaces(name string, sizes ...float64) (map[float64]font.Face, bool) {
	loadedFont, ok := fontsMap[name]
	if !ok {
		return nil, false
	}
	faces := make(map[float64]font.Face)
	for _, size := range sizes {
		faces[size] = truetype.NewFace(loadedFont, &truetype.Options{
			Size: size,
		})
	}
	return faces, true
}

func GetLoadedImage(name string) (image.Image, bool) {
	img, ok := imagesMap[name]
	return img, ok
}
