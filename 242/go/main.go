package main

import (
	"fmt"
)

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	result1 := make(map[byte]int)
	result2 := make(map[byte]int)

	for i := range s {
		result1[s[i]]++
		result2[t[i]]++
	}

	for k, v := range result1 {
		if v != result2[k] {
			return false
		}
	}

	return true
}

func main() {
	str1 := "anagram"
	str2 := "nagaram"
	expect := true
	fmt.Println(expect, isAnagram(str1, str2))
}
