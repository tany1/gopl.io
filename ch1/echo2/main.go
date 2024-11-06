// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 6.
//!+

// Echo2 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	s, sep := "", ""
	start := time.Now()
	for i, arg := range os.Args[1:] {
		s += sep + strconv.Itoa(i) + ":" + arg
		sep = "\n"
	}
	fmt.Println(s, time.Since(start))
	fmt.Printf("%vus\n", time.Since(start).Microseconds())
}

//!-
