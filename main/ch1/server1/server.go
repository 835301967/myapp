package main

import (
	"fmt"
	_ "git.myscrm.cn/component-center/cc-file-box-sdk"
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
)
const (
	whiteIndex = 0 // first color in palette
	blackIndex = 1 // next color in palette
)

var palette = []color.Color{color.White, color.Black}

func main()  {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		get := r.URL.Query().Get("c")
		float, _ := strconv.ParseFloat(get, 10)
		lissajous(w,float)
	})
	log.Fatal(http.ListenAndServe("localhost:8000",nil))

}

func handler (w http.ResponseWriter,r *http.Request)  {
	fmt.Fprintf(w,"UPL.PATH=%q\n",r.URL.Path)
}


func lissajous(out io.Writer,cycles float64) {
	const (
		res     = 0.001 // angular resolution
		size    = 100   // image canvas covers [-size..+size]
		nframes = 64    // number of animation frames
		delay   = 8     // delay between frames in 10ms units
	)
	if cycles==0{
		cycles = 5
	}
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
				blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}
