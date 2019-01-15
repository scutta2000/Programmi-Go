package main

import (
	"fmt"
	"io"
)

func compareSlice(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}
	for c := 0; c < len(a); c++ {
		if a[c] != b[c] {
			return false
		}
	}
	return true
}
func sostituisci(s, old, new []byte, n int) []byte {
	var start, end, count int
	for start = 0; start < len[s]-len[old]; start++ {
		end = start + len[old]
		if compareSlice(s[start:end], old) {
			count++
		}
		if n == count {
			//return append(s[:start], new, ...s[end:])
		}
	}
	return s
}
func main() {
	input := io.Args[1:]
	fmt.Println(input[1])
	fmt.Println(sostituisci(input[1], input[2], input[3], input[4]))
}
