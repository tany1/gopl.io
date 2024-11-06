// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 8.

// Echo3 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

// !+
func main() {
	start := time.Now()
	fmt.Println(strings.Join(os.Args[1:], " "), time.Since(start))
	fmt.Printf("%vus\n", time.Since(start).Microseconds())
}

//!-
