package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsAnagram_1(t *testing.T) {
	str1 := "anagram"
	str2 := "nagaram"

	got := isAnagram(str1, str2)
	want := true

	assert.Equal(t, want, got)
}

func TestIsAnagram_2(t *testing.T) {
	str1 := "rat"
	str2 := "car"

	got := isAnagram(str1, str2)
	want := false

	assert.Equal(t, want, got)
}

func TestIsAnagram_3(t *testing.T) {
	str1 := "aacc"
	str2 := "ccac"

	got := isAnagram(str1, str2)
	want := false

	assert.Equal(t, want, got)
}
