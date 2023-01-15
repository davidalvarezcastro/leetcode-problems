package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestPalindrome_1(t *testing.T) {
	str := "babad"

	got := longestPalindrome(str)
	want := "aba"

	assert.Equal(t, want, got)
}

func TestLongestPalindrome_2(t *testing.T) {
	str := "cbbd"

	got := longestPalindrome(str)
	want := "bb"

	assert.Equal(t, want, got)
}

func TestLongestPalindrome_3(t *testing.T) {
	str := "aedavadaa"

	got := longestPalindrome(str)
	want := "davad"

	assert.Equal(t, want, got)
}
