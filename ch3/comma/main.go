// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 73.

// Comma prints its argument numbers with a comma at each power of 1000.
//
// Example:
//
//	$ go build gopl.io/ch3/comma
//	$ ./comma 1 12 123 1234 1234567890
//	1
//	12
//	123
//	1,234
//	1,234,567,890
package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", commaFloatingPoint(os.Args[i]))
	}
}

// !+
// comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

func comma2(s string) string {
	n := len(s)

	if n <= 3 {
		return s
	}

	var buf bytes.Buffer
	remainder := n % 3
	commas := n / 3

	if remainder == 0 {
		commas -= 1
	}

	readBytes := 0

	for i := n - 1; i >= 0; i-- {
		buf.WriteByte(s[i])
		readBytes++

		if commas > 0 && readBytes%3 == 0 {
			buf.WriteString(",")
			commas--
		}
	}

	reverseBuf(&buf)

	return buf.String()
}

func commaFloatingPoint(s string) string {
	prefix := ""

	// Check for sign
	if strings.HasPrefix(s, "+") || strings.HasPrefix(s, "-") {
		prefix = s[:1]
		s = s[1:]
	}

	// Split upon decimal point
	parts := strings.Split(s, ".")

	if len(parts) == 1 {
		return prefix + comma2(parts[0])
	}

	return prefix + comma2(parts[0]) + "." + parts[1]
}

func reverseBuf(buf *bytes.Buffer) {
	data := buf.Bytes()
	for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {
		data[i], data[j] = data[j], data[i]
	}

	buf.Reset()
	buf.Write(data)
}

//!-
