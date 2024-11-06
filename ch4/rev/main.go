// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 86.

// Rev reverses a slice.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	//!+array
	a := [...]int{0, 1, 2, 3, 4, 5}
	reverse(a[:])
	fmt.Println(a) // "[5 4 3 2 1 0]"
	//!-array

	//!+slice
	s := []int{0, 1, 2, 3, 4, 5}
	// Rotate s left by two positions.
	reverse(s[:2])
	reverse(s[2:])
	reverse(s)
	fmt.Println(s) // "[2 3 4 5 0 1]"
	//!-slice

	b := [32]int{0, 1, 2, 3, 4, 5}
	reverse2(&b)
	fmt.Println(b, "reverse with array") // "[5 4 3 2 1 0]"

 	rotate([]int{0, 1, 2, 3, 4, 5}, 2)

	fmt.Println(s, "rotate in a single pass")

	fmt.Println(dedupAdjacent([]string{"a", "a", "b", "c", "c", "c", "d", "e", "e"}), "remove adjacent duplicates")

	fmt.Println(string(squashSpaces([]byte("This string has \t  multiple  \n spaces."))), "squash spaces")

	fmt.Println(string(reverseByte([]byte("reverse\t"))), "reverse byte")

	// Interactive test of reverse.
	input := bufio.NewScanner(os.Stdin)
outer:
	for input.Scan() {
		var ints []int
		for _, s := range strings.Fields(input.Text()) {
			x, err := strconv.ParseInt(s, 10, 64)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				continue outer
			}
			ints = append(ints, int(x))
		}
		reverse(ints)
		fmt.Printf("%v\n", ints)
	}
	// NOTE: ignoring potential errors from input.Err()
}

//!+rev
// reverse reverses a slice of ints in place.
func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

//!-rev

// Execrise 4.3 - reverse an array using a pointer to the array
func reverse2(s *[32]int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

// Exercise 4.4 - rotate a slice by n positions in a single pass
func rotate(s []int, n int) {

	for i := 0; i < n; i++ {
		s[0], s[len(s)-1] = s[len(s)-1], s[0]
	}
}

// Exercise 4.5 - remove adjacent duplicates in a slice of strings in place
func dedupAdjacent(s []string) []string {
	i := 0
	for _, v := range s {
		if s[i] == v {
			continue
		}
		i++
		s[i] = v
	}
	return s[:i+1]
}

// Exercise 4.6 - squash adjacent spaces in a byte slice
func squashSpaces(s []byte) []byte {
	count := 0
	currIndex := 0

	for i := range s {
		isSpace := unicode.IsSpace(rune(s[i]))
		if isSpace {
			count++
		} else {
			count = 0
		}

		if count > 1 {
			continue
		}
		currIndex++

		if isSpace {
			s[currIndex-1] = ' '
		} else {
			s[currIndex-1] = s[i] 
		}
	}

	return s[:currIndex]
}

// Exercise 4.7 - reverse a byte slice
func reverseByte(s []byte) []byte {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}

	return s
}