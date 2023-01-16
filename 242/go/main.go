package main

import (
	"fmt"
)

func isAnagram(s string, t string) bool {
	result := [26]int{}

	for _, v := range s {
		result[v-'a']++
	}

	for _, v := range t {
		result[v-'a']--
	}

	return result == [26]int{}
}

func main() {
	str1 := "anagram"
	str2 := "nagaram"
	expect := true
	fmt.Println(expect, isAnagram(str1, str2))
}
