package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	counts := map[string]int{}
    for _, word := range strings.Fields(s) {
		if _, ok := counts[word]; ok {
			counts[word]++
		} else {
			counts[word] = 1
		}
	}
	return counts
}

func main() {
	wc.Test(WordCount)
}
