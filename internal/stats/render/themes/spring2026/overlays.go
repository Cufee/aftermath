package spring2026

import (
	"image"
	"image/color"
	"image/draw"
	"math"
	"math/rand"

	"github.com/nao1215/imaging"
)

func makeForegroundOverlay(petals []image.Image) func(image.Image, image.Rectangle, int) image.Image {
	nrgbaPetals := make([]*image.NRGBA, len(petals))
	for i, p := range petals {
		if n, ok := p.(*image.NRGBA); ok {
			nrgbaPetals[i] = n
		} else {
			b := p.Bounds()
			n := image.NewNRGBA(image.Rect(0, 0, b.Dx(), b.Dy()))
			draw.Draw(n, n.Bounds(), p, b.Min, draw.Src)
			nrgbaPetals[i] = n
		}
	}

	return func(rendered image.Image, frame image.Rectangle, seed int) image.Image {
		if len(nrgbaPetals) == 0 {
			return rendered
		}

		rng := rand.New(rand.NewSource(int64(seed)))

		w := frame.Dx()
		h := frame.Dy()
		overlay := image.NewNRGBA(image.Rect(0, 0, w, h))

		band := min(w, h) / 4
		perimeter := 2 * (w + h)
		basePetals := float64(perimeter) * 0.020
		petalCount := max(10, int(basePetals*(0.85+rng.Float64()*0.30)))

		for range petalCount {
			petal := nrgbaPetals[rng.Intn(len(nrgbaPetals))]
			opacity := 0.35 + rng.Float64()*0.45

			sw := petal.Bounds().Dx()
			sh := petal.Bounds().Dy()

			depth := int(float64(band) * rng.Float64() * rng.Float64())

			var x, y int
			edge := rng.Intn(4)
			switch edge {
			case 0: // top
				x = rng.Intn(w) - sw/2
				y = depth
			case 1: // right
				x = w - depth - sw
				y = rng.Intn(h) - sh/2
			case 2: // bottom
				x = rng.Intn(w) - sw/2
				y = h - depth - sh
			case 3: // left
				x = depth
				y = rng.Intn(h) - sh/2
			}

			x = clamp(x, 0, w-sw)
			y = clamp(y, 0, h-sh)

			blitNRGBA(overlay, petal, x, y, opacity)
		}

		dst := toNRGBA(rendered)
		compositeOver(dst, overlay)
		return dst
	}
}

// blitNRGBA alpha-blends src onto dst at (ox, oy) with an extra opacity multiplier.
func blitNRGBA(dst *image.NRGBA, src *image.NRGBA, ox, oy int, opacity float64) {
	sb := src.Bounds()
	dstStride := dst.Stride
	srcStride := src.Stride

	for sy := sb.Min.Y; sy < sb.Max.Y; sy++ {
		dy := oy + sy - sb.Min.Y
		if dy < 0 || dy >= dst.Bounds().Dy() {
			continue
		}
		srcRow := sy * srcStride
		dstRow := dy * dstStride
		for sx := sb.Min.X; sx < sb.Max.X; sx++ {
			dx := ox + sx - sb.Min.X
			if dx < 0 || dx >= dst.Bounds().Dx() {
				continue
			}
			si := srcRow + sx*4
			sa := float64(src.Pix[si+3]) / 255.0 * opacity
			if sa < 1.0/255.0 {
				continue
			}
			di := dstRow + dx*4
			sr := float64(src.Pix[si])
			sg := float64(src.Pix[si+1])
			sb := float64(src.Pix[si+2])

			da := float64(dst.Pix[di+3]) / 255.0
			outA := sa + da*(1-sa)
			if outA > 0 {
				dst.Pix[di] = uint8((sr*sa + float64(dst.Pix[di])*(da*(1-sa))) / outA)
				dst.Pix[di+1] = uint8((sg*sa + float64(dst.Pix[di+1])*(da*(1-sa))) / outA)
				dst.Pix[di+2] = uint8((sb*sa + float64(dst.Pix[di+2])*(da*(1-sa))) / outA)
				dst.Pix[di+3] = uint8(outA * 255)
			}
		}
	}
}

// compositeOver alpha-composites src over dst, mutating dst in place.
func compositeOver(dst *image.NRGBA, src *image.NRGBA) {
	db := dst.Bounds()
	sb := src.Bounds()
	h := min(db.Dy(), sb.Dy())
	w := min(db.Dx(), sb.Dx())

	for y := range h {
		dstOff := y * dst.Stride
		srcOff := y * src.Stride
		for x := range w {
			si := srcOff + x*4
			sa := src.Pix[si+3]
			if sa == 0 {
				continue
			}
			di := dstOff + x*4
			if sa == 255 {
				dst.Pix[di] = src.Pix[si]
				dst.Pix[di+1] = src.Pix[si+1]
				dst.Pix[di+2] = src.Pix[si+2]
				dst.Pix[di+3] = 255
				continue
			}
			srcA := float64(sa) / 255.0
			dstA := float64(dst.Pix[di+3]) / 255.0
			outA := srcA + dstA*(1-srcA)
			if outA > 0 {
				inv := dstA * (1 - srcA)
				dst.Pix[di] = uint8((float64(src.Pix[si])*srcA + float64(dst.Pix[di])*inv) / outA)
				dst.Pix[di+1] = uint8((float64(src.Pix[si+1])*srcA + float64(dst.Pix[di+1])*inv) / outA)
				dst.Pix[di+2] = uint8((float64(src.Pix[si+2])*srcA + float64(dst.Pix[di+2])*inv) / outA)
				dst.Pix[di+3] = uint8(outA * 255)
			}
		}
	}
}

func toNRGBA(img image.Image) *image.NRGBA {
	if n, ok := img.(*image.NRGBA); ok {
		cp := image.NewNRGBA(n.Bounds())
		copy(cp.Pix, n.Pix)
		return cp
	}
	b := img.Bounds()
	dst := image.NewNRGBA(image.Rect(0, 0, b.Dx(), b.Dy()))
	draw.Draw(dst, dst.Bounds(), img, b.Min, draw.Src)
	return dst
}

func makeBackgroundOverlay() func(image.Rectangle, int) image.Image {
	return func(bounds image.Rectangle, _ int) image.Image {
		w := bounds.Dx()
		h := bounds.Dy()

		// quarter-res vignette, upscaled for smoothness
		sw, sh := max(w/4, 1), max(h/4, 1)
		img := image.NewNRGBA(image.Rect(0, 0, sw, sh))
		cx := float64(sw) / 2
		cy := float64(sh) / 2
		maxDist := math.Sqrt(cx*cx + cy*cy)

		for y := range sh {
			for x := range sw {
				dx := float64(x) - cx
				dy := float64(y) - cy
				dist := math.Sqrt(dx*dx+dy*dy) / maxDist

				if dist > 0.35 {
					alpha := (dist - 0.35) / 0.65
					alpha = alpha * alpha
					alpha = min(alpha*0.55, 1.0)
					img.SetNRGBA(x, y, color.NRGBA{0, 0, 0, uint8(alpha * 255)})
				}
			}
		}

		return imaging.Resize(img, w, h, imaging.Linear)
	}
}

func clamp(v, lo, hi int) int {
	if v < lo {
		return lo
	}
	if v > hi {
		return hi
	}
	return v
}
