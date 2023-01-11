package main

import (
	"fmt"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func lengthOfLongestSubstring(s string) int {
	distance := 0
	prev := -1
	result := 0

	for i, v := range s {

		if pos := strings.LastIndex(s[:i], string(v)); pos > prev {
			prev = pos
		}

		if distance = i - prev; distance > result {
			result = distance
		}
	}

	return result
}

func main() {
	s := "dvdf"
	fmt.Println(lengthOfLongestSubstring(s))
}
