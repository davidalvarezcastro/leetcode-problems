package main

import (
	"fmt"
)

func containsDuplicate(nums []int) bool {
	result := make(map[int]bool)

	for _, v := range nums {
		if _, ok := result[v]; ok {
			return true
		}

		result[v] = true

	}

	return false
}

func main() {
	input := []int{1, 2, 3, 1}
	expect := true
	fmt.Println(expect, containsDuplicate(input))
}
