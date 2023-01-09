package main

import "fmt"

func twoSum(nums []int, target int) []int {
	resuts := make(map[int]int)

	for index, value := range nums {
		if other_index, ok := resuts[target-value]; ok {
			return []int{other_index, index}
		}
		resuts[value] = index
	}

	return nil
}

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}
