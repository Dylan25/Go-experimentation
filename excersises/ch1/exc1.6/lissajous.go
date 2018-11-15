// Lissoujas generates GIF animations of random Lissajous figures using multiple colors.
package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
	"time"
)

var palette = []color.Color{color.Black, green, blue, lavender, honeydew}
var green = color.RGBA{0x00, 0x80, 0x00, 0xff}
var blue = color.RGBA{0xf0, 0x80, 0x80, 0xff}
var lavender = color.RGBA{0xe6, 0xe6, 0xfa, 0xff}
var honeydew = color.RGBA{0xf0, 0xff, 0xf0, 0xff}

const (
	blackIndex    = 0 // first color in palette
	greenIndex    = 1 // next color in palette
	blueIndex     = 2
	lavenderIndex = 3
	honeydewIndex = 4
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5     // Number of complete x oscillator revolutions
		res     = 0.001 // angular resolution
		size    = 100   // image canvas covers [-size..+size]
		nframes = 64    // number of animation frames
		delay   = 8     // delay between frames in 10ms units
	)
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			currentColor := uint8((i / 16) + 1)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), currentColor)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) //ignores encoding errors
}