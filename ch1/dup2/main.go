// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 10.
//!+

// Dup2 prints the count and text of lines that appear more than once
// in the input.  It reads from stdin or from a list of named files.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Map by the filename, then the line & count
	counts := make(map[string]map[string]int)
	fmt.Println(counts)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}

			// Init filename map
			counts[f.Name()] = make(map[string]int)

			fmt.Println(counts)
			countLines(f, counts)
			f.Close()
		}
	}
	for file, linemap := range counts {
		for line, n := range linemap {
			if n > 1 {
				fmt.Printf("%s -> %s -> %d\n", file, line, n)
			}
		}
	}
}

func countLines(f *os.File, counts map[string]map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[f.Name()][input.Text()]++
	}
	// NOTE: ignoring potential errors from input.Err()
}

//!-
