package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	path, err := filepath.Rel(".", "ch4/wordfreq/words.txt")

	if err != nil {
		panic(err)
	}

	// Map of words to their counts
	counts := make(map[string]int)
	
	var file *os.File
	file, err = os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		word := strings.ToLower(scanner.Text())
		counts[word]++
	}

	// Print the word frequency
	for word, count := range counts {
		fmt.Printf("%s: %d\n", word, count)
	}
}