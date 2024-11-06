package main

import (
	"fmt"
	"os"
)

func main() {
	first := os.Args[1]
	second := os.Args[2]

	if len(first) != len(second) {
		fmt.Println("the strings are not anagrams")
		return
	}

	if areAnagrams(first, second) {
		fmt.Println("the strings are anagrams")
	} else {
		fmt.Println("the strings are not anagrams")
	}

}

func areAnagrams(s1, s2 string) bool {
	lettersMap := make(map[byte]int)

	// Count the letters in the first string
	for _, c := range s1 {
		lettersMap[byte(c)]++
	}

	// Count the letters in the second string
	for _, c := range s2 {
		lettersMap[byte(c)]--
	}

	for _, count := range lettersMap {
		if count != 0 {
			return false
		}
	}

	return true
}
