package main

import (
	"strings"
)

func WordFrequency(s string) map[string]int {
	frequency := make(map[string]int)
	words := strings.Fields(s)
	for _, word := range words {
		frequency[word]++
	}
	return frequency
}
