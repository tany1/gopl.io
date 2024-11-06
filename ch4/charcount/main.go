// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 97.
//!+

// Charcount computes counts of Unicode characters.
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

func main() {
	counts := make(map[rune]int)    // counts of Unicode characters
	var utflen [utf8.UTFMax + 1]int // count of lengths of UTF-8 encodings
	invalid := 0                    // count of invalid UTF-8 characters
	countsByCategory := make(map[string]int)

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune() // returns rune, nbytes, error
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}

		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		counts[r]++
		utflen[n]++

		if unicode.IsLetter(r) {
			countsByCategory["Letter"]++
		}
		if unicode.IsNumber(r) {
			countsByCategory["Number"]++
		}
		if unicode.IsSymbol(r) {
			countsByCategory["Symbol"]++
		}
		if unicode.IsPunct(r) {
			countsByCategory["Punct"]++
		}
		if unicode.IsSpace(r) {
			countsByCategory["Space"]++
		}
		if unicode.IsControl(r) {
			countsByCategory["Control"]++
		}
		if unicode.IsMark(r) {
			countsByCategory["Mark"]++
		}
		if unicode.IsGraphic(r) {
			countsByCategory["Graphic"]++
		}
		if unicode.IsPrint(r) {
			countsByCategory["Print"]++
		}
		if unicode.IsLower(r) {
			countsByCategory["Lower"]++
		}
		if unicode.IsUpper(r) {
			countsByCategory["Upper"]++
		}
		if unicode.IsTitle(r) {
			countsByCategory["Title"]++
		}
		if unicode.IsDigit(r) {
			countsByCategory["Digit"]++
		}
	}
	fmt.Printf("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Print("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	fmt.Print("\ntype\tcount\n")
	for t, n := range countsByCategory {
		if n > 0 {
			fmt.Printf("%s\t%d\n", t, n)
		}
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}

//!-
