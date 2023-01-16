package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContainsDuplicate_1(t *testing.T) {
	input := []int{1, 2, 3, 1}
	got := containsDuplicate(input)
	want := true

	assert.Equal(t, want, got)
}

func TestContainsDuplicate_2(t *testing.T) {
	input := []int{1, 2, 3, 4}
	got := containsDuplicate(input)
	want := false

	assert.Equal(t, want, got)
}

func TestContainsDuplicate_3(t *testing.T) {
	input := []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}
	got := containsDuplicate(input)
	want := true

	assert.Equal(t, want, got)
}
