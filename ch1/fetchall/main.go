// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 17.
//!+

// Fetchall fetches URLs in parallel and reports their times and sizes.
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	const prefix = "https://dexonline.ro/definitie/"

	start := time.Now()
	ch := make(chan string)
	words := getWords(os.Args[1])
	fmt.Println(words)
	for _, word := range words {
		go fetch(prefix+word, ch) // start a goroutine
	}
	for range words {
		fmt.Println(<-ch) // receive from channel ch
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}

	nbytes, err := io.Copy(io.Discard, resp.Body)
	resp.Body.Close() // don't leak resources
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d  %s %d", secs, nbytes, url, resp.StatusCode)
}

// Get a slice of words from a file
func getWords(fileName string) []string {
	data, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
	}

	// Append to an empty string slice
	return append(make([]string, 0), strings.Split(string(data), "\n")...)
}

//!-
