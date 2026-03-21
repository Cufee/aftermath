//go:build ignore

package main

import (
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	"math"
	"os"
	"path/filepath"
	"strings"

	"github.com/nao1215/imaging"
)

const (
	baseSize   = 14
	padding    = 10
	blurLength = 10
	gaussSigma = 2.0
	rotations  = 8
)

var scales = []float64{0.5, 0.75, 1.0, 1.5}

type tintPreset struct {
	color   color.NRGBA
	opacity float64
}

var tints = []tintPreset{
	{color.NRGBA{255, 180, 200, 255}, 0.35}, // light pink
	{color.NRGBA{200, 120, 150, 255}, 0.40}, // muted rose
	{color.NRGBA{140, 70, 100, 255}, 0.45},  // dark plum
}

func main() {
	generatePetals()
	generateBlurredBackground()
}

func generatePetals() {
	sourceDir := filepath.Join("petals", "source")
	outDir := filepath.Join("petals", "processed")

	os.RemoveAll(outDir)
	if err := os.MkdirAll(outDir, 0o755); err != nil {
		panic(err)
	}

	entries, err := os.ReadDir(sourceDir)
	if err != nil {
		panic(err)
	}

	idx := 0
	for _, entry := range entries {
		name := entry.Name()
		if !strings.HasSuffix(strings.ToLower(name), ".png") {
			continue
		}
		src, err := imaging.Open(filepath.Join(sourceDir, name))
		if err != nil {
			panic(fmt.Sprintf("failed to open %s: %v", name, err))
		}

		small := imaging.Resize(src, baseSize, 0, imaging.Lanczos)

		for _, tint := range tints {
			tinted := tintImage(small, tint.color, tint.opacity)
			canvasSize := baseSize + padding*2

			for v := range rotations {
				rotAngle := float64(v) * (360.0 / float64(rotations))
				blurAngle := rotAngle + 30

				canvas := imaging.New(canvasSize, canvasSize, color.Transparent)
				rotated := imaging.Rotate(tinted, rotAngle, color.Transparent)
				canvas = imaging.PasteCenter(canvas, rotated)
				canvas = motionBlur(canvas, blurAngle, blurLength)
				canvas = imaging.Blur(canvas, gaussSigma)
				canvas = trimAlpha(canvas, 2)

				for _, scale := range scales {
					var scaled image.Image = canvas
					if scale != 1.0 {
						sw := max(int(float64(canvas.Bounds().Dx())*scale), 1)
						scaled = imaging.Resize(canvas, sw, 0, imaging.Linear)
					}

					outPath := filepath.Join(outDir, fmt.Sprintf("petal_%03d.png", idx))
					f, err := os.Create(outPath)
					if err != nil {
						panic(err)
					}
					if err := png.Encode(f, scaled); err != nil {
						f.Close()
						panic(err)
					}
					f.Close()
					idx++
				}
			}
		}
	}
	fmt.Printf("generated %d processed petal images\n", idx)
}

func generateBlurredBackground() {
	const (
		minSigma = 0.5
		maxSigma = 8.0
		levels   = 8
		quality  = 95
	)

	bg, err := imaging.Open("background.jpg")
	if err != nil {
		panic(fmt.Sprintf("failed to open background.jpg: %v", err))
	}

	dm, err := imaging.Open("depthmap.png")
	if err != nil {
		panic(fmt.Sprintf("failed to open depthmap.png: %v", err))
	}

	w, h := bg.Bounds().Dx(), bg.Bounds().Dy()
	dm = imaging.Resize(dm, w, h, imaging.Linear)

	stack := make([]*image.NRGBA, levels)
	for i := range levels {
		sigma := minSigma + (maxSigma-minSigma)*float64(i)/float64(levels-1)
		stack[i] = imaging.Blur(bg, sigma)
		fmt.Printf("  blur level %d/%d (sigma=%.1f)\n", i+1, levels, sigma)
	}

	out := image.NewNRGBA(image.Rect(0, 0, w, h))
	for y := range h {
		for x := range w {
			r, _, _, _ := dm.At(x, y).RGBA()
			depth := float64(r>>8) / 255.0 // 1.0 = close (bright), 0.0 = far (dark)

			// close → low index (less blur), far → high index (more blur)
			idx := (1.0 - depth) * float64(levels-1)
			lo := int(math.Floor(idx))
			hi := int(math.Ceil(idx))
			lo = max(lo, 0)
			hi = min(hi, levels-1)
			frac := idx - float64(lo)

			c1 := stack[lo].NRGBAAt(x, y)
			c2 := stack[hi].NRGBAAt(x, y)
			out.SetNRGBA(x, y, color.NRGBA{
				R: uint8(lerp(float64(c1.R), float64(c2.R), frac)),
				G: uint8(lerp(float64(c1.G), float64(c2.G), frac)),
				B: uint8(lerp(float64(c1.B), float64(c2.B), frac)),
				A: 255,
			})
		}
	}

	f, err := os.Create("background_blurred.jpg")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	if err := jpeg.Encode(f, out, &jpeg.Options{Quality: quality}); err != nil {
		panic(err)
	}
	fmt.Println("generated depth-blurred background")
}

// trimAlpha crops transparent borders, keeping at least minPad pixels of padding.
func trimAlpha(img *image.NRGBA, minPad int) *image.NRGBA {
	b := img.Bounds()
	minX, minY, maxX, maxY := b.Max.X, b.Max.Y, b.Min.X, b.Min.Y

	for y := b.Min.Y; y < b.Max.Y; y++ {
		for x := b.Min.X; x < b.Max.X; x++ {
			if img.NRGBAAt(x, y).A > 0 {
				minX = min(minX, x)
				minY = min(minY, y)
				maxX = max(maxX, x+1)
				maxY = max(maxY, y+1)
			}
		}
	}
	if maxX <= minX || maxY <= minY {
		return img
	}

	minX = max(minX-minPad, b.Min.X)
	minY = max(minY-minPad, b.Min.Y)
	maxX = min(maxX+minPad, b.Max.X)
	maxY = min(maxY+minPad, b.Max.Y)

	return imaging.Crop(img, image.Rect(minX, minY, maxX, maxY))
}

func tintImage(img image.Image, tint color.NRGBA, opacity float64) *image.NRGBA {
	b := img.Bounds()
	out := imaging.New(b.Dx(), b.Dy(), color.Transparent)
	for y := b.Min.Y; y < b.Max.Y; y++ {
		for x := b.Min.X; x < b.Max.X; x++ {
			r0, g0, b0, a0 := img.At(x, y).RGBA()
			if a0 == 0 {
				continue
			}
			r := lerp(float64(r0>>8), float64(tint.R), opacity)
			g := lerp(float64(g0>>8), float64(tint.G), opacity)
			bl := lerp(float64(b0>>8), float64(tint.B), opacity)
			out.SetNRGBA(x-b.Min.X, y-b.Min.Y, color.NRGBA{uint8(r), uint8(g), uint8(bl), uint8(a0 >> 8)})
		}
	}
	return out
}

func lerp(a, b, t float64) float64 {
	return a*(1-t) + b*t
}

func motionBlur(img image.Image, angleDeg float64, length int) *image.NRGBA {
	w := img.Bounds().Dx()
	h := img.Bounds().Dy()
	angleRad := angleDeg * math.Pi / 180.0
	dx := math.Cos(angleRad)
	dy := math.Sin(angleRad)

	result := imaging.New(w, h, color.Transparent)
	for i := range length {
		t := float64(i)/float64(length-1) - 0.5
		ox := int(math.Round(t * float64(length) * dx))
		oy := int(math.Round(t * float64(length) * dy))
		result = imaging.Overlay(result, img, image.Pt(ox, oy), 1.0/float64(length))
	}
	return result
}
