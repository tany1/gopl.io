// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 16.
//!+

// Fetch prints the content found at each specified URL.
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	const httpPrefix = "http://"

	for _, url := range os.Args[1:] {
		hasHttp := strings.HasPrefix(url, httpPrefix)

		// Override url in case prefix is missing
		if !hasHttp {
			url = httpPrefix + url
		}

		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}

		b, err := io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\nstatus code: %v", url, err, resp.StatusCode)
			os.Exit(1)
		}
		fmt.Printf("%d\n", b)
		fmt.Println(resp.StatusCode, resp.Status)
	}
}

//!-
