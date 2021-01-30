package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"math"
	"os"
)

func main() {
	fn := "gopher.png"
	file, err := os.Open(fn)

	if err != nil {
		log.Printf("failed opening %s: %s", fn, err)
		panic(err.Error())
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		panic(err.Error())
	}

	bounds := img.Bounds()
	w, h := bounds.Max.X, bounds.Max.Y
	grayScale := image.NewGray(image.Rectangle{image.Point{0, 0}, image.Point{w, h}})
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			imageColor := img.At(x, y)
			rr, gg, bb, _ := imageColor.RGBA()
			r := math.Pow(float64(rr), 2.2)
			g := math.Pow(float64(gg), 2.2)
			b := math.Pow(float64(bb), 2.2)
			m := math.Pow(0.3*r+0.59*g+0.11*b, 1/2.2)
			Y := uint16(m + 0.5)
			grayColor := color.Gray{uint8(Y >> 8)}
			grayScale.Set(x, y, grayColor)
		}
	}

	nfn := "grayscaled.png"
	nf, err := os.Create(nfn)
	if err != nil {
		log.Printf("failed creating %s: %s", nfn, err)
		panic(err.Error())
	}
	defer nf.Close()
	png.Encode(nf, grayScale)
}
