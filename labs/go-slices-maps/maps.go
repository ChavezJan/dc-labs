package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {

	Rmap := map[string]int{}
	for _, word := range strings.Fields(s) {
		Rmap[word]++
	}
	return Rmap
}

func main() {
	wc.Test(WordCount)
}
