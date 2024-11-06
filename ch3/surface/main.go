// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 58.
//!+

// Surface computes an SVG rendering of a 3-D surface function.
package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"strconv"
	// "os"
	// "sort"
	// "strconv"
)

const (
	width, height = 600, 600            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)
var readWidth, readHeight = width, height

// Red
var color = "#ff0000"

func main() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		// Read from request
		var err error
		readWidth, err = strconv.Atoi(r.URL.Query().Get("width"))
		if err != nil {
			fmt.Println(err)
			return
		}
		readHeight, err = strconv.Atoi(r.URL.Query().Get("height"))
		if err != nil {
			fmt.Println(err)
			return
		}

		// Set content type
		w.Header().Set("Content-Type", "image/svg+xml")

		writerFunc(w)
	}

	http.HandleFunc("/", handler)

	//!-http
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

var zValues []float64

func writerFunc(s io.Writer) {
	fmt.Fprintf(s, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", readWidth, readHeight)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			// Reset color
			color = "#ff0000"

			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)
			fmt.Fprintf(s, "<polygon points='%g,%g %g,%g %g,%g %g,%g' style=\"fill:%s\" />\n",
				ax, ay, bx, by, cx, cy, dx, dy, color)
		}
	}
	fmt.Fprint(s, "</svg>")
}

func corner(i, j int) (float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y)

	if math.IsNaN(z) {
		z = 1.
	}

	if z <= 0.03 {
		color = "#0000ff"
	}

	zValues = append(zValues, z)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}

//!-
