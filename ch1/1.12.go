// Lissajous generates GIF animations of random Lissajous figures.
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"
)

/*
 * Modify the lissajous server to read parameter values from the URL.
 * For example, you might arrange it so that a URL like http://localhost:8000/?cycles=20
 * sets the number of cycles to 20 instead of the default 5. Use strconv.Atoi function
 * to convert the string parameter into an integer.
 */

type lissajousParamters struct {
	Cycles  int
	Res     float64
	Size    int
	Nframes int
	Delay   int
}

var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0 // first color in palette
	blackIndex = 1 // next color in palette
)

func main() {
	// The sequence of images is deterministic unless we seed
	// the pseudo-random number generator using the current time.
	// Thanks to Randall McPherson for pointing out the omission.
	rand.Seed(time.Now().UTC().UnixNano())

	handler := func(w http.ResponseWriter, r *http.Request) {
		params := &lissajousParamters{}
		setLissajousParams(params, r)
		lissajous(w, params)
	}
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
	return
}

func setLissajousParams(params *lissajousParamters, r *http.Request) {
	// set the defaults
	params.Cycles = 5
	params.Res = 0.001
	params.Size = 100
	params.Nframes = 64
	params.Delay = 8
	query := r.URL.Query()

	if cyclesParam, ok := query["cycles"]; ok {
		if cycles, err := strconv.Atoi(cyclesParam[0]); err != nil {
			fmt.Fprintf(os.Stderr, "parsing cycles: %v", err)
		} else {
			params.Cycles = cycles
		}
	}
	if resParam, ok := query["res"]; ok {
		if res, err := strconv.ParseFloat(resParam[0], 64); err != nil {
			fmt.Fprintf(os.Stderr, "parsing res: %v", err)
		} else {
			params.Res = res
		}

	}
	if sizeParam, ok := query["size"]; ok {
		if size, err := strconv.Atoi(sizeParam[0]); err != nil {
			fmt.Fprintf(os.Stderr, "parsing size: %v", err)
		} else {
			params.Size = size
		}
	}
	if nframesParam, ok := query["nframes"]; ok {
		if nframes, err := strconv.Atoi(nframesParam[0]); err != nil {
			fmt.Fprintf(os.Stderr, "parsing nframes: %v", err)
		} else {
			params.Nframes = nframes
		}
	}
	if delayParam, ok := query["delay"]; ok {
		if delay, err := strconv.Atoi(delayParam[0]); err != nil {
			fmt.Fprintf(os.Stderr, "parsing delay: %v", err)
		} else {
			params.Delay = delay
		}
	}
}

func lissajous(out io.Writer, p *lissajousParamters) {
	cycles := float64(p.Cycles)  // number of complete x oscillator revolutions
	res := p.Res                 // angular resolution
	size := p.Size               // image canvas covers [-size..+size]
	nframes := p.Nframes         // number of animation frames
	delay := p.Delay             // delay between frames in 10ms units
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference

	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*float64(size)+0.5), size+int(y*float64(size)+0.5),
				blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}
