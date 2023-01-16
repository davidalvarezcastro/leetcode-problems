package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGroupAnagrams_1(t *testing.T) {
	str := []string{"eat", "tea", "tan", "ate", "nat", "bat"}

	got := groupAnagrams(str)
	want := [][]string{{"eat", "tea", "ate"}, {"tan", "nat"}, {"bat"}}

	assert.Equal(t, want, got)
}

func TestGroupAnagrams_2(t *testing.T) {
	str := []string{""}

	got := groupAnagrams(str)
	want := [][]string{{""}}

	assert.Equal(t, want, got)
}

func TestGroupAnagrams_3(t *testing.T) {
	str := []string{"a"}

	got := groupAnagrams(str)
	want := [][]string{{"a"}}

	assert.Equal(t, want, got)
}
