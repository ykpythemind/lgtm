package generator

import (
	"fmt"
	"image"
	"image/draw"
	"image/jpeg"
	_ "image/jpeg"
	"io"

	pkgFont "golang.org/x/image/font"
	"golang.org/x/image/math/fixed"

	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font/gofont/gobold"
)

type Option struct {
	FontSize float64
}

type Generator struct {
	Option Option
}

func NewGenerator(opt Option) *Generator {
	return &Generator{Option: opt}
}

func (g *Generator) Generate(source io.Reader, dest io.Writer) error {
	font, err := truetype.Parse(gobold.TTF)
	if err != nil {
		return err
	}

	fontOpt := truetype.Options{
		Size:              g.Option.FontSize,
		DPI:               0,
		Hinting:           0,
		GlyphCacheEntries: 0,
		SubPixelsX:        0,
		SubPixelsY:        0,
	}
	face := truetype.NewFace(font, &fontOpt)
	defer face.Close()

	srcImg, _, err := image.Decode(source)
	if err != nil {
		return err
	}

	dstImg := image.NewRGBA(image.Rect(0, 0, srcImg.Bounds().Dx(), srcImg.Bounds().Dy()))
	draw.Draw(dstImg, dstImg.Rect, srcImg, image.Point{}, draw.Src)
	dot := fixed.P(0, 0)
	fontDrawer := pkgFont.Drawer{Dst: dstImg, Src: image.White, Face: face, Dot: dot}

	text := "LGTM"

	bounds, _ := fontDrawer.BoundString(text)
	fmt.Printf("%+v\n", bounds)

	dx := int((bounds.Max.X - bounds.Min.X) / 60)
	dy := int((bounds.Max.Y - bounds.Min.Y) / 60)

	fmt.Printf("%v, %v\n", dx, dy)

	x := (srcImg.Bounds().Dx() - dx) / 2
	y := (srcImg.Bounds().Dy() - dy) / 2

	fontDrawer.Dot = fixed.P(x, y)
	fmt.Printf("%+v", fontDrawer.Dot)
	fontDrawer.DrawString(text)

	o := jpeg.Options{Quality: 100}
	err = jpeg.Encode(dest, dstImg, &o)
	if err != nil {
		return err
	}

	return nil
}
