package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLengthOfLongestSubstring_1(t *testing.T) {
	s := "abcabcbb"

	got := lengthOfLongestSubstring(s)
	want := 3

	assert.Equal(t, want, got)
}

func TestLengthOfLongestSubstring_2(t *testing.T) {
	s := "bbbbb"

	got := lengthOfLongestSubstring(s)
	want := 1

	assert.Equal(t, want, got)
}

func TestLengthOfLongestSubstring_3(t *testing.T) {
	s := "pwwkew"

	got := lengthOfLongestSubstring(s)
	want := 3

	assert.Equal(t, want, got)
}

func TestLengthOfLongestSubstring_4(t *testing.T) {
	s := "au"

	got := lengthOfLongestSubstring(s)
	want := 2

	assert.Equal(t, want, got)
}

func TestLengthOfLongestSubstring_5(t *testing.T) {
	s := "aab"

	got := lengthOfLongestSubstring(s)
	want := 2

	assert.Equal(t, want, got)
}

func TestLengthOfLongestSubstring_6(t *testing.T) {
	s := "dvdf"

	got := lengthOfLongestSubstring(s)
	want := 3

	assert.Equal(t, want, got)
}
