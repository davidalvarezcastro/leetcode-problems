package main

import (
	"fmt"
	"math"
)

func expand(s string, left int, rigth int) int {
	for left >= 0 && rigth < len(s) && s[left] == s[rigth] {
		left--
		rigth++
	}

	return rigth - left - 1
}

func longestPalindrome(s string) string {
	end := 0
	start := 0

	for i := range s {
		len1 := expand(s, i, i)   // babad
		len2 := expand(s, i, i+1) // abba
		len := int(math.Max(float64(len1), float64(len2)))

		if len > (end - start) {
			start = i - (len-1)/2
			end = i + len/2
		}

	}

	return s[start : end+1]
}

func main() {
	str := "abba"
	fmt.Println(longestPalindrome(str))
}
