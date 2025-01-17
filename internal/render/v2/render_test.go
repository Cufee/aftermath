package render

import (
	"image"
	"image/color"
	"image/png"
	"os"
	"path/filepath"
	"testing"

	"github.com/cufee/aftermath/internal/render/v2/style"
	"github.com/cufee/aftermath/tests/path"
	"github.com/matryer/is"
)

var _ = saveImage

var contentSize = 12.0
var contentColorValue uint32
var contentColor = color.RGBA{255, 255, 255, 255}

func init() {
	_, _, _, a := contentColor.RGBA()
	contentColorValue = a
}

func TestApplyPadding(t *testing.T) {
	is := is.New(t)

	content := NewEmptyContent(style.NewStyle(style.Parent(style.Style{Width: contentSize, Height: contentSize, BackgroundColor: contentColor})))

	t.Run("uniform", func(t *testing.T) {
		is := is.New(t)

		wrapper := NewBlocksContent(style.NewStyle(
			style.SetPadding(10),
		), content)

		d := wrapper.Dimensions()
		is.True(d.width == ceil(contentSize)+20)
		is.True(d.height == ceil(contentSize)+20)

		img, err := wrapper.Render()
		is.NoErr(err)

		{
			_, _, _, a := img.At(9, 9).RGBA()
			is.True(a == 0)
		}
		{
			_, _, _, a := img.At(10, 10).RGBA()
			is.True(a == contentColorValue)
		}
	})

	t.Run("X", func(t *testing.T) {
		wrapper := NewBlocksContent(style.NewStyle(
			style.SetPaddingX(10),
		), content)

		d := wrapper.Dimensions()
		is.True(d.width == ceil(contentSize)+20)
		is.True(d.height == ceil(contentSize))
	})

	t.Run("Y", func(t *testing.T) {
		wrapper := NewBlocksContent(style.NewStyle(
			style.SetPaddingY(10),
		), content)

		d := wrapper.Dimensions()
		is.True(d.width == ceil(contentSize))
		is.True(d.height == ceil(contentSize)+20)
	})

	t.Run("overwrite", func(t *testing.T) {
		wrapper := NewBlocksContent(style.NewStyle(
			style.SetPadding(10),
			style.SetPadding(0),
		), content)

		d := wrapper.Dimensions()
		is.True(d.width == ceil(contentSize))
		is.True(d.height == ceil(contentSize))
	})

	t.Run("left", func(t *testing.T) {
		wrapper := NewBlocksContent(style.NewStyle(
			style.Parent(style.Style{
				PaddingLeft: 10,
			}),
		), content)

		d := wrapper.Dimensions()
		is.True(d.width == ceil(contentSize)+10)
		is.True(d.height == ceil(contentSize))

		img, err := wrapper.Render()
		is.NoErr(err)

		{
			_, _, _, a := img.At(9, 0).RGBA()
			is.True(a == 0)
		}
		{
			_, _, _, a := img.At(10, 0).RGBA()
			is.True(a == contentColorValue)
		}
	})

	t.Run("top", func(t *testing.T) {
		wrapper := NewBlocksContent(style.NewStyle(
			style.Parent(style.Style{
				PaddingTop: 10,
			}),
		), content)

		d := wrapper.Dimensions()
		is.True(d.width == ceil(contentSize))
		is.True(d.height == ceil(contentSize)+10)

		img, err := wrapper.Render()
		is.NoErr(err)

		{
			_, _, _, a := img.At(0, 9).RGBA()
			is.True(a == 0)
		}
		{
			_, _, _, a := img.At(0, 10).RGBA()
			is.True(a == contentColorValue)
		}
	})
}

func TestRenderJustify(t *testing.T) {
	is := is.New(t)

	content := NewEmptyContent(style.NewStyle(style.Parent(style.Style{Width: contentSize, Height: contentSize, BackgroundColor: contentColor})))

	t.Run("horizontal", func(t *testing.T) {
		t.Run("start", func(t *testing.T) {
			is := is.New(t)

			wrapper := NewBlocksContent(style.NewStyle(
				style.SetWidth(contentSize*2),
			), content)

			img, err := wrapper.Render()
			is.NoErr(err)

			{
				_, _, _, imgA := img.At(0, 0).RGBA()
				is.True(imgA == contentColorValue)
			}
			{
				_, _, _, imgA := img.At(int(contentSize*2-1), 0).RGBA()
				is.True(imgA == 0)
			}
		})

		t.Run("center", func(t *testing.T) {
			is := is.New(t)

			wrapper := NewBlocksContent(style.NewStyle(
				style.Parent(style.Style{
					JustifyContent: style.JustifyContentCenter,
				}),
				style.SetWidth(contentSize*2),
			), content)

			img, err := wrapper.Render()
			is.NoErr(err)

			{
				_, _, _, imgA := img.At(int(contentSize/3), 0).RGBA()
				is.True(imgA == 0)
			}
			{
				_, _, _, imgA := img.At(int(contentSize), 0).RGBA()
				is.True(imgA == contentColorValue)
			}
			{
				_, _, _, imgA := img.At(int(contentSize*2-contentSize/3), 0).RGBA()
				is.True(imgA == 0)
			}
		})

		t.Run("end", func(t *testing.T) {
			is := is.New(t)

			wrapper := NewBlocksContent(style.NewStyle(
				style.Parent(style.Style{
					JustifyContent: style.JustifyContentEnd,
				}),
				style.SetWidth(contentSize*2),
			), content)

			img, err := wrapper.Render()
			is.NoErr(err)

			{
				_, _, _, imgA := img.At(int(contentSize-1), 0).RGBA()
				is.True(imgA == 0)
			}
			{
				_, _, _, imgA := img.At(int(contentSize*2-1), 0).RGBA()
				is.True(imgA == contentColorValue)
			}
		})
	})

	t.Run("vertical", func(t *testing.T) {
		t.Run("start", func(t *testing.T) {
			is := is.New(t)

			wrapper := NewBlocksContent(style.NewStyle(
				style.Parent(style.Style{
					Direction: style.DirectionVertical,
				}),
				style.SetHeight(contentSize*2),
			), content)

			img, err := wrapper.Render()
			is.NoErr(err)

			{
				_, _, _, imgA := img.At(0, 0).RGBA()
				is.True(imgA == contentColorValue)
			}
			{
				_, _, _, imgA := img.At(0, int(contentSize*2-1)).RGBA()
				is.True(imgA == 0)
			}
		})

		t.Run("center", func(t *testing.T) {
			is := is.New(t)

			wrapper := NewBlocksContent(style.NewStyle(
				style.Parent(style.Style{
					JustifyContent: style.JustifyContentCenter,
					Direction:      style.DirectionVertical,
				}),
				style.SetHeight(contentSize*2),
			), content)

			img, err := wrapper.Render()
			is.NoErr(err)

			{
				_, _, _, imgA := img.At(0, int(contentSize/4)).RGBA()
				is.True(imgA == 0)
			}
			{
				_, _, _, imgA := img.At(0, int(contentSize)).RGBA()
				is.True(imgA == contentColorValue)
			}
			{
				_, _, _, imgA := img.At(0, int(contentSize*2-contentSize/4)).RGBA()
				is.True(imgA == 0)
			}
		})

		t.Run("end", func(t *testing.T) {
			is := is.New(t)

			wrapper := NewBlocksContent(style.NewStyle(
				style.Parent(style.Style{
					JustifyContent: style.JustifyContentEnd,
					Direction:      style.DirectionVertical,
				}),
				style.SetHeight(contentSize*2),
			), content)

			img, err := wrapper.Render()
			is.NoErr(err)

			{
				_, _, _, imgA := img.At(0, int(contentSize-1)).RGBA()
				is.True(imgA == 0)
			}
			{
				_, _, _, imgA := img.At(0, int(contentSize*2-1)).RGBA()
				is.True(imgA == contentColorValue)
			}
		})
	})
}

func saveImage(is *is.I, img image.Image) {
	f, err := os.Create(filepath.Join(path.Root(), "tmp", "test_render_blocks.png"))
	is.NoErr(err)

	err = png.Encode(f, img)
	is.NoErr(err)
}
