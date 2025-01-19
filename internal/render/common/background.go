package common

import (
	"image"
	"image/color"
	"math"
	"math/rand"
	"sync"

	"github.com/fogleman/gg"
	"github.com/nao1215/imaging"
)

var DefaultBackgroundBlur float64 = 5
var GlassEffectBackgroundBlur float64 = DefaultBackgroundBlur * 5

var globalLogoCacheMx sync.Mutex
var globalLogoCache = make(map[color.Color]image.Image)

func AddDefaultBrandedOverlay(background image.Image, colors []color.Color, seed int, colorChance float32) image.Image {
	if len(colors) < 1 {
		colors = DefaultLogoColorOptions
	}

	source := rand.NewSource(int64(seed))
	r := rand.New(source)
	for i := range colors {
		if r.Float32() > colorChance {
			colors[i] = TextSecondary
		}
	}

	size := 15
	overlay := NewBrandedBackground(background.Bounds().Dx()*2, background.Bounds().Dy()*2, size, -size/2, colors, seed)
	return imaging.OverlayCenter(background, overlay, 0.5)
}

func NewBrandedBackground(width, height, logoSize, padding int, colors []color.Color, hashSeed int) image.Image {
	// 1/3 of the image should be left for logos
	rows := max((height-padding*2)/3/logoSize, 2)
	cols := max((width-padding*2)/3/logoSize, 2)
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

				posX := float64(padding + c*(logoSize+xGap))
				posY := float64(padding + r*(logoSize+yGap))
				source := rand.NewSource(int64(hashSeed) + int64(posX)*51 + int64(posY)*37)
				rnd := rand.New(source)

				if n := rnd.Float32(); n < 0.5 {
					return
				}

				clr := pickColor(colors, rnd)
				scale := pickScaleFactor(rnd)
				rotation := pickRotationRad(rnd)

				mx.Lock()
				logo := getLogo(clr)
				mx.Unlock()

				logoAdjusted := imaging.New(logoSize, logoSize, color.Transparent)
				logoAdjusted = imaging.PasteCenter(logoAdjusted, logo)
				logoAdjusted = imaging.Rotate(logoAdjusted, rotation, color.Transparent)
				logoAdjusted = imaging.Resize(logoAdjusted, int(float64(logoSize)*scale), int(float64(logoSize)*scale), imaging.Linear)

				xJ, yJ := pickPositionJitter(rnd)
				posX += xJ
				posY += yJ

				mx.Lock()
				ctx.DrawImage(logoAdjusted, int(posX), int(posY))
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
func pickColor(colors []color.Color, r *rand.Rand) color.Color {
	if len(colors) < 1 {
		return color.White
	}

	index := r.Intn(len(colors))
	return colors[index]
}

// pickScaleFactor function that generates a scale factor clamped between 0.8 and 1.2, influenced by image size
func pickScaleFactor(r *rand.Rand) float64 {
	// Clamp the scale factor between 0.5 and 1.5
	scaleFactor := 0.5 + (r.Float64())
	return scaleFactor
}

// pickPositionJitter function that generates an x,y  position offset based on the hash seed
func pickPositionJitter(r *rand.Rand) (float64, float64) {
	// Clamp between 0.5 and 1.5
	xJitter := -0.5 + r.Float64()
	yJitter := -0.5 + r.Float64()
	return xJitter, yJitter
}

func pickRotationRad(r *rand.Rand) float64 {
	// Clamp the rotation angle between -2.5π and 2.5π radians
	rotationRad := -2.5*math.Pi + (r.Float64() * (5 * math.Pi))
	return rotationRad
}

func BlurWithMask(content image.Image, mask *image.Alpha, blur, maskBlur float64) (image.Image, error) {
	ctx := gg.NewContext(content.Bounds().Dx(), content.Bounds().Dy())
	err := ctx.SetMask(mask)
	if err != nil {
		return nil, err
	}
	ctx.DrawImage(imaging.Blur(content, maskBlur), 0, 0)

	content = imaging.Blur(content, blur)
	content = imaging.OverlayCenter(content, ctx.Image(), 1) // paste masked content on top of content
	return content, nil
}
