package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

var palette = []color.Color{color.Black}

func addColors() {
	for i := 0; i < 16; i++ {
		palette = append(palette, color.RGBA{uint8((15 - i) * 17), 0x00, uint8(i * 17), 0xff})
	}
	for i := 0; i < 16; i++ {
		palette = append(palette, color.RGBA{uint8(i * 17), 0x00, uint8((15 - i) * 17), 0xff})
	}
}

func main() {
	addColors()
	rand.Seed(time.Now().UTC().UnixNano())
	http.HandleFunc("/", handler)
	http.HandleFunc("/lissajous", lissajous)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}
}

func lissajous(w http.ResponseWriter, r *http.Request) {
	const (
		tmpCycles  = 5
		tmpRes     = 0.001
		tmpSize    = 100
		tmpNframes = 64
		tmpDelay   = 8
	)
	var cycles int
	var res float64
	var size int
	var nframes int
	var delay int
	params := map[string]string{}
	r.ParseForm()
	for k, v := range r.Form {
		params[k] = v[0]
	}
	if v, ok := params["cycles"]; ok {
		cycles, _ = strconv.Atoi(v)
	} else {
		cycles = tmpCycles
	}
	if v, ok := params["res"]; ok {
		res, _ = strconv.ParseFloat(v, 64)
	} else {
		res = tmpRes
	}
	if v, ok := params["size"]; ok {
		size, _ = strconv.Atoi(v)
	} else {
		size = tmpSize
	}
	if v, ok := params["nframes"]; ok {
		nframes, _ = strconv.Atoi(v)
	} else {
		nframes = tmpNframes
	}
	if v, ok := params["delay"]; ok {
		delay, _ = strconv.Atoi(v)
	} else {
		delay = tmpDelay
	}
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < float64(cycles)*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			foregroundColor := uint8(i%32 + 1)
			img.SetColorIndex(size+int(x*float64(size)+0.5), size+int(y*float64(size)+0.5), foregroundColor)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(w, &anim)
}
