package spring2026

import (
	"bytes"
	"embed"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"path"
	"strings"

	"github.com/nao1215/imaging"
)

//go:embed assets/background.jpg
var backgroundBytes []byte

//go:embed assets/petals/processed
var petalsFS embed.FS

var (
	backgroundImage image.Image
	processedPetals []image.Image
)

func init() {
	var err error
	backgroundImage, _, err = image.Decode(bytes.NewReader(backgroundBytes))
	if err != nil {
		panic("spring2026: failed to decode background: " + err.Error())
	}
	backgroundImage = imaging.Blur(backgroundImage, 3)
	backgroundBytes = nil

	entries, err := petalsFS.ReadDir("assets/petals/processed")
	if err != nil {
		panic("spring2026: failed to read assets/petals/processed: " + err.Error())
	}
	for _, entry := range entries {
		name := entry.Name()
		if !strings.HasSuffix(strings.ToLower(name), ".png") {
			continue
		}
		data, err := petalsFS.ReadFile(path.Join("assets/petals/processed", name))
		if err != nil {
			panic("spring2026: failed to read " + name + ": " + err.Error())
		}
		img, _, err := image.Decode(bytes.NewReader(data))
		if err != nil {
			panic("spring2026: failed to decode " + name + ": " + err.Error())
		}
		processedPetals = append(processedPetals, img)
	}
}
