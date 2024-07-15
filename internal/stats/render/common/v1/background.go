package common

import (
	"image"
	"image/color"
	"math"
	"math/rand"
	"sync"

	"github.com/disintegration/imaging"
	"github.com/fogleman/gg"
)

var globalLogoCacheMx sync.Mutex
var globalLogoCache = make(map[color.Color]image.Image)

func NewBrandedBackground(width, height int, colors []color.Color, hashSeed int) image.Image {
	var logoSize = 30
	var padding = 30

	// 2/3 of the image should be left for logos
	rows := (height - padding*2) * 2 / 3 / logoSize
	cols := (width - padding*2) * 2 / 3 / logoSize
	// the rest is gaps
	xGapsTotal := width - padding*2 - (cols)*logoSize
	yGapsTotal := height - padding*2 - (rows)*logoSize
	xGap := xGapsTotal / (cols - 1)
	yGap := yGapsTotal / (rows - 1)

	localLogoCache := make(map[color.Color]image.Image)
	getLogo := func(color color.Color) image.Image {
		if img := localLogoCache[color]; img != nil {
			return img
		}

		logo, ok := globalLogoCache[color]
		if !ok {
			logo = AftermathLogo(color, SmallLogoOptions())
			globalLogoCacheMx.Lock()
			globalLogoCache[color] = logo
			globalLogoCacheMx.Unlock()
		}
		logo = imaging.Fill(logo, logoSize, logoSize, imaging.Center, imaging.Lanczos)
		localLogoCache[color] = logo // mx is already locked inside the caller
		return logo
	}

	ctx := gg.NewContext(width, height)
	var makeFn []func()
	var mx sync.Mutex
	var wg sync.WaitGroup

	for c := range cols {
		for r := range rows {
			wg.Add(1)
			makeFn = append(makeFn, func() {
				defer wg.Done()

				clr := pickColor(c*logoSize, r*logoSize, colors, hashSeed)
				scale := pickScaleFactor(c*logoSize, r*logoSize, hashSeed)
				rotation := pickRotationRad(c*logoSize, r*logoSize, hashSeed)

				mx.Lock()
				logo := getLogo(clr)
				mx.Unlock()

				logoAdjusted := imaging.New(logoSize, logoSize, color.Transparent)
				logoAdjusted = imaging.PasteCenter(logoAdjusted, logo)
				logoAdjusted = imaging.Rotate(logoAdjusted, rotation, color.Transparent)
				logoAdjusted = imaging.Resize(logoAdjusted, int(float64(logoSize)*scale), int(float64(logoSize)*scale), imaging.Linear)

				posX := padding + c*(logoSize+xGap)
				posY := padding + r*(logoSize+yGap)

				mx.Lock()
				ctx.DrawImage(logoAdjusted, posX, posY)
				mx.Unlock()
			})
		}
	}
	for _, fn := range makeFn {
		go fn() // wg.Done is called inside already, wg.Add was called already
	}
	wg.Wait()

	return ctx.Image()
}

// pickColor function that includes hashSeed in the hash calculation
func pickColor(x, y int, colors []color.Color, hashSeed int) color.Color {
	if len(colors) < 1 {
		return color.White
	}

	// Create a new rand source with the hashSeed and pixel position
	source := rand.NewSource(int64(hashSeed) + int64(x)*31 + int64(y)*37)
	r := rand.New(source)

	index := r.Intn(len(colors))
	return colors[index]
}

// pickScaleFactor function that generates a scale factor clamped between 0.8 and 1.2, influenced by image size
func pickScaleFactor(x, y, hashSeed int) float64 {
	// Create a new rand source with the hashSeed and pixel position
	source := rand.NewSource(int64(hashSeed) + int64(x)*31 + int64(y)*37)
	r := rand.New(source)

	// Generate a pseudo-random float between 0 and 1
	randomValue := r.Float64()
	// Clamp the scale factor between 0.8 and 1.2
	scaleFactor := 0.8 + (randomValue * 0.4)
	return scaleFactor
}

func pickRotationRad(x, y, hashSeed int) float64 {
	// Create a new rand source with the hashSeed and pixel position
	source := rand.NewSource(int64(hashSeed) + int64(x)*31 + int64(y)*37)
	r := rand.New(source)

	// Generate a pseudo-random float between 0 and 1
	randomValue := r.Float64()
	// Clamp the rotation angle between -4π and 4π radians
	rotationRad := -4*math.Pi + (randomValue * (8 * math.Pi))
	return rotationRad
}
