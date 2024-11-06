// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 61.
//!+

// Mandelbrot emits a PNG image of the Mandelbrot fractal.
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"math/big"
	"math/cmplx"
	"net/http"
	// "os"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

const (
	xmin, ymin, xmax, ymax = -2, -2, +2, +2
	width, height          = 512, 512
	subpixelx, subpixelsy  = 2, 2
)

func mandelbrotHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var _w, _h int
	var err error
	_w, _h = width, height

	// Override with params
	_w, err = strconv.Atoi(ps.ByName("width"))
	if err != nil {
		fmt.Fprintf(w, "invalid width")
	}

	_h, err = strconv.Atoi(ps.ByName("height"))
	if err != nil {
		fmt.Fprintf(w, "invalid height")
	}

	img := image.NewRGBA(image.Rect(0, 0, _w, _h))
	for py := 0; py < _w; py++ {
		y := float64(py)/float64(_w)*(ymax-ymin) + ymin
		for px := 0; px < _h; px++ {
			x := float64(px)/float64(_h)*(xmax-xmin) + xmin

			// Compute supersampling value
			// var value float64

			// for i := 0; i < subpixelx; i++ {
			// 	for j := 0; j < subpixelsy; j++ {
			// 		x = float64(px+i)/width*(xmax-xmin) + xmin
			// 		y = float64(py+i)/width*(ymax-ymin) + ymin

			// 		value += mandelbrotVal(complex(x, y))
			// 	}
			// }
			// value /= subpixelx * subpixelsy

			z := complex(x, y)
			// zBig := Complex{big.NewFloat(x), big.NewFloat(y)}
			// Image point (px, py) represents complex value z.
			img.SetRGBA(px, py, mandelbrot(z))
			// img.Set(px, py, mandelbrotBig(zBig))
		}
	}
	png.Encode(w, img) // NOTE: ignoring errors
}

func main() {
	router := httprouter.New()
	router.GET("/render/:width/:height", mandelbrotHandler)

	log.Fatal(http.ListenAndServe(":8080", router))
}

func mandelbrot(z complex128) color.RGBA {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.RGBA{0x00, 0x00, 255 - contrast*n, 0xff}
		}
	}
	return color.RGBA{0x00, 0x00, 0x00, 0xff}
}

type Complex struct {
	Real, Imag *big.Float
}

// Multiply 2 numbers in big.Float space
func complexMulBig(x, y Complex) Complex {
	result := Complex{}
	result.Real = big.NewFloat(0)
	result.Imag = big.NewFloat(0)

	result.Real.Mul(x.Real, y.Real).Sub(result.Real, new(big.Float).Mul(x.Imag, y.Imag))
	result.Imag.Mul(x.Real, y.Imag).Add(result.Imag, new(big.Float).Mul(x.Imag, y.Real))

	return result
}

// Extract the absolute value of a complex
// number in big.Float space
func complexAbsBig(x Complex) big.Float {
	result := big.NewFloat(0)

	result.Sqrt(result.Mul(x.Real, x.Real).Add(result, x.Imag.Mul(x.Imag, x.Imag)))

	return *result
}

func mandelbrotBig(z Complex) color.RGBA {
	const iterations = 225
	const contrast = 15

	var v = Complex{new(big.Float), new(big.Float)}
	for n := uint8(0); n < iterations; n++ {
		result := complexMulBig(v, v)

		result.Real.Add(result.Real, z.Real)
		result.Imag.Add(result.Imag, z.Imag)

		v = result

		resultAbs := complexAbsBig(v)

		// fmt.Println(resultAbs.String(), resultAbs.Cmp(big.NewFloat(1.94)))

		if resultAbs.Cmp(big.NewFloat(2)) == 1 {
			return color.RGBA{0x00, 0x00, uint8(255 * n), 0xff}
		}
	}

	return color.RGBA{0x00, 0x00, 0x00, 0xff}
}

func mandelbrotVal(z complex128) float64 {
	const iterations = 200

	var v complex128

	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return float64(n)
		}
	}

	return 255
}

func rgbaColorFromValue(value uint8) color.RGBA {
	const contrast = 15

	return color.RGBA{0x00, 0x00, 255 - contrast*value, 0xff}
}

//!-

// Some other interesting functions:

func acos(z complex128) color.Color {
	v := cmplx.Acos(z)
	blue := uint8(real(v)*128) + 127
	red := uint8(imag(v)*128) + 127
	return color.YCbCr{192, blue, red}
}

func sqrt(z complex128) color.Color {
	v := cmplx.Sqrt(z)
	blue := uint8(real(v)*128) + 127
	red := uint8(imag(v)*128) + 127
	return color.YCbCr{128, blue, red}
}

// f(x) = x^4 - 1
//
// z' = z - f(z)/f'(z)
//
//	= z - (z^4 - 1) / (4 * z^3)
//	= z - (z - 1/z^3) / 4
func newton(z complex128) color.RGBA {
	const iterations = 37
	const contrast = 7
	for i := uint8(0); i < iterations; i++ {
		z -= (z - 1/(z*z*z)) / 4
		if cmplx.Abs(z*z*z*z-1) < 1e-6 {
			// return color.Gray{255 - contrast*i}
			return color.RGBA{0x00, 0x00, 255 - contrast*i, 0xff}
		}
	}
	return color.RGBA{A: 0xff}
}
