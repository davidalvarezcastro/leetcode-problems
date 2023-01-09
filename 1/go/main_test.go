package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTwoSum_1(t *testing.T) {
	got := twoSum([]int{2, 7, 11, 15}, 9)
	want := []int{0, 1}

	assert.Equal(t, got, want)
}

func TestTwoSum_2(t *testing.T) {
	got := twoSum([]int{3, 3}, 6)
	want := []int{0, 1}

	assert.Equal(t, got, want)
}

func TestTwoSum_3(t *testing.T) {
	got := twoSum([]int{3, 2, 4}, 6)
	want := []int{1, 2}

	assert.Equal(t, got, want)
}
