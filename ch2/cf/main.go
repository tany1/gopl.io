// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 43.
//!+

// Cf converts its numeric argument to Celsius and Fahrenheit.
package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"

	"gopl.io/ch2/lengthconv"
	"gopl.io/ch2/tempconv"
)

var unit = flag.String("u", "temperature", "unit type")

func main() {
	flag.Parse()

	nonFlagArgs := flag.Args()

	if len(nonFlagArgs) > 0 {

		for _, arg := range nonFlagArgs {
			fmt.Println(arg)
			value, err := strconv.ParseFloat(arg, 64)
			if err != nil {
				fmt.Fprintf(os.Stderr, "cf: %v\n", err)
				os.Exit(1)
			}

			switch *unit {
			case "temperature":
				convertTemperature(value)
			case "length":
				convertLength(value)
			}
		}
	} else {
		scanner := bufio.NewScanner(os.Stdin)

		for scanner.Scan() {
			text := scanner.Text()
			
			value, err := strconv.ParseFloat(text, 64)
			if err != nil {
				fmt.Fprintf(os.Stderr, "cf: %v\n", err)
				os.Exit(1)
			}

			switch *unit {
			case "temperature":
				convertTemperature(value)
			case "length":
				convertLength(value)
			}
		}

		if err := scanner.Err(); err != nil {
			fmt.Printf("scanner error: %v", err)
		}
	}
}

//!-

func convertTemperature(t float64) {
	f := tempconv.Fahrenheit(t)
	c := tempconv.Celsius(t)

	fmt.Printf("%s = %s, %s = %s\n",
		f, tempconv.FToC(f),
		c, tempconv.CToF(c),
	)
}

func convertLength(l float64) {
	f := lengthconv.Foot(l)
	m := lengthconv.Meter(l)

	fmt.Printf("%s = %s, %s = %s\n",
		f, lengthconv.FToM(f),
		m, lengthconv.MToF(m),
	)
}
