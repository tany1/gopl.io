// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

//!+

// Package lengthconv performs Celsius and Fahrenheit conversions.
package lengthconv

import "fmt"

type Foot float64
type Meter float64

const (
	MeterToFootRate = 3.281
)

func (f Foot) String() string  { return fmt.Sprintf("%gft", f) }
func (m Meter) String() string { return fmt.Sprintf("%gm", m) }

//!-
