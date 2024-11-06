// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 41.

//!+

package lengthconv

// MToF converts a Meter length to Feet.
func MToF(m Meter) Foot { return Foot(m * MeterToFootRate) }

// FToM converts a Foot length to Meter.
func FToM(f Foot) Meter { return Meter(f / MeterToFootRate) }

//!-
