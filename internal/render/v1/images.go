package render

import (
	"image"
	"math"

	"github.com/cufee/aftermath/internal/log"
	"github.com/pkg/errors"

	"github.com/fogleman/gg"
	"github.com/nao1215/imaging"
)

func AddGlassMaskedBackground(content, background image.Image, mask *image.Alpha, style Style) image.Image {
	if background == nil || mask == nil {
		return content
	}

	background = imaging.Fill(background, content.Bounds().Dx(), content.Bounds().Dy(), imaging.Center, imaging.Linear)
	blurred, err := BlurWithMask(background, mask, DefaultBackgroundBlur, GlassEffectBackgroundBlur)
	if err == nil {
		background = blurred
	}

	return AddBackground(content, background, style)
}

func AddBackground(content, background image.Image, style Style) image.Image {
	if background == nil {
		return content
	}

	// Fill background with black and round the corners
	frameCtx := gg.NewContext(content.Bounds().Dx(), content.Bounds().Dy())
	if style.BackgroundColor != nil {
		frameCtx.SetColor(style.BackgroundColor)
		frameCtx.Clear()
	}

	clipRoundedRect(frameCtx, style.BorderRadius)

	// Resize the background image to fit the cards
	if background.Bounds().Dx() != frameCtx.Width() || background.Bounds().Dy() != frameCtx.Height() {
		background = imaging.Fill(background, frameCtx.Width(), frameCtx.Height(), imaging.Center, imaging.NearestNeighbor)
	}
	if style.Blur > 0 {
		background = imaging.Blur(background, style.Blur)
	}

	frameCtx.DrawImage(background, 0, 0)
	frameCtx.DrawImage(content, 0, 0)

	return frameCtx.Image()
}

func renderImages(images []image.Image, style Style) (image.Image, error) {
	if len(images) < 1 {
		return nil, errors.New("no images to render")
	}

	imageSize := getDetailedSize(images, style)

	var lastX, lastY float64 = style.PaddingX, style.PaddingY
	var justifyOffsetX, justifyOffsetY float64
	var elementWidth, elementHeight float64 // Not used at the moment, will enforce each element to be the same size if set

	/*
		TODO: Some math here under some configurations is certainly broken and needs to be fixed or removed
	*/

	// Set correct gaps and offsets based on justify content
	switch style.JustifyContent {
	case JustifyContentCenter:
		lastX += float64(imageSize.extraSpacingX / 2)
		lastY += float64(imageSize.extraSpacingY / 2)
	case JustifyContentEnd:
		lastX += float64(imageSize.extraSpacingX)
		lastY += float64(imageSize.extraSpacingY)
	case JustifyContentSpaceBetween:
		if len(images) > 0 {
			justifyOffsetX = float64(imageSize.extraSpacingX / float64(len(images)-1))
			justifyOffsetY = float64(imageSize.extraSpacingY / float64(len(images)-1))
		}
	case JustifyContentSpaceAround:
		spacingX := float64(imageSize.extraSpacingX / float64(len(images)+1))
		spacingY := float64(imageSize.extraSpacingY / float64(len(images)+1))
		justifyOffsetX = spacingX
		justifyOffsetY = spacingY
		lastX += spacingX
		lastY += spacingY
	default: // JustifyContentStart
		// 0,0
	}

	ctx := gg.NewContext(int(math.Ceil(imageSize.width)), int(math.Ceil(imageSize.height)))

	clipRoundedRect(ctx, style.BorderRadius)

	if style.BackgroundColor != nil {
		ctx.DrawRectangle(0, 0, imageSize.width, imageSize.height)
		ctx.SetColor(style.BackgroundColor)
		ctx.Fill()
	}

	for i, img := range images {
		if img == nil {
			log.Error().Stack().Msg("a nil image was passed to render")
			return nil, errors.New("one of the images is nil")
		}

		posX, posY := lastX, lastY

		imgWidth := float64(img.Bounds().Dx())
		imgHeight := float64(img.Bounds().Dy())

		targetWidth := max(imgWidth, elementWidth)
		targetHeight := max(imgHeight, elementHeight)

		switch style.Direction {
		case DirectionVertical:
			if i > 0 {
				posY += justifyOffsetY + style.Gap
			}
			lastY = posY + targetHeight

			switch style.AlignItems {
			case AlignItemsCenter:
				posX = (imageSize.width - imgWidth) / 2
			case AlignItemsEnd:
				posX = imageSize.width - imgWidth - style.PaddingX
			default: // AlignItemsStart
				posX = style.PaddingX
			}
		default: // DirectionHorizontal
			if i > 0 {
				posX += justifyOffsetX + style.Gap
			}
			lastX = posX + targetWidth

			switch style.AlignItems {
			case AlignItemsCenter:
				posY = (imageSize.height - imgHeight) / 2
			case AlignItemsEnd:
				posY = imageSize.height - imgHeight - style.PaddingY
			default: // AlignItemsStart
				posY = style.PaddingY
			}

		}

		if style.Debug {
			ctx.SetColor(getDebugColor())
			ctx.DrawRectangle(posX, posY, targetWidth, targetHeight)
			ctx.Stroke()
		}

		ctx.DrawImage(img, int(math.Ceil(posX+(targetWidth-imgWidth)/2)), int(math.Ceil(posY+(targetHeight-imgHeight)/2)))
	}

	return ctx.Image(), nil
}

type imageSize struct {
	width  float64
	height float64

	totalGap float64
	// The amount of extra spacing added to the image, used for alignment
	extraSpacingX float64
	extraSpacingY float64

	maxElementWidth  float64
	maxElementHeight float64
}

func getDetailedSize(images []image.Image, style Style) imageSize {
	imageWidth, imageHeight := style.Width, style.Height

	var totalWidth float64
	var totalHeight float64
	maxWidth, maxHeight := 0.0, 0.0
	for _, img := range images {
		if img == nil {
			continue
		}

		imgX := float64(img.Bounds().Dx())
		maxWidth = max(maxWidth, imgX)

		imgY := float64(img.Bounds().Dy())
		maxHeight = max(maxHeight, imgY)

		if style.Direction == DirectionHorizontal {
			totalWidth += float64(img.Bounds().Dx())
			totalHeight = max(totalHeight, imgY)
		} else {
			totalHeight += float64(img.Bounds().Dy())
			totalWidth = max(totalWidth, imgX)
		}
	}

	totalWidth += style.PaddingX * 2
	totalHeight += style.PaddingY * 2
	totalGap := float64(len(images)-1) * style.Gap

	switch style.Direction {
	case DirectionVertical:
		totalHeight += totalGap
	case DirectionHorizontal:
		totalWidth += totalGap
	}

	if imageWidth < 1 {
		imageWidth = totalWidth
	}
	if imageHeight < 1 {
		imageHeight = totalHeight
	}

	extraSpacingX := max(0, imageWidth-totalWidth)
	extraSpacingY := max(0, imageHeight-totalHeight)

	return imageSize{
		width:            imageWidth,
		height:           imageHeight,
		totalGap:         totalGap,
		extraSpacingX:    extraSpacingX,
		extraSpacingY:    extraSpacingY,
		maxElementWidth:  maxWidth,
		maxElementHeight: maxHeight,
	}
}

func clipRoundedRect(ctx *gg.Context, borderRadius float64) {
	if borderRadius <= 0 {
		return
	}

	width, height := float64(ctx.Width()), float64(ctx.Height())
	// top left
	ctx.DrawEllipticalArc(borderRadius, borderRadius, borderRadius, borderRadius, gg.Radians(270), gg.Radians(180))
	// bottom left
	ctx.LineTo(0, height-borderRadius)
	ctx.DrawEllipticalArc(borderRadius, height-borderRadius, borderRadius, borderRadius, gg.Radians(180), gg.Radians(90))
	// bottom right
	ctx.LineTo(width-borderRadius, height)
	ctx.DrawEllipticalArc(width-borderRadius, height-borderRadius, borderRadius, borderRadius, gg.Radians(90), gg.Radians(0))
	// top right
	ctx.LineTo(width, borderRadius)
	ctx.DrawEllipticalArc(width-borderRadius, borderRadius, borderRadius, borderRadius, gg.Radians(0), gg.Radians(-90))
	// clip final path
	ctx.Clip()
}
