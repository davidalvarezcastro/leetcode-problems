package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddTwoNumbers_1(t *testing.T) {
	l1 := &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3, Next: nil}}}
	l2 := &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4, Next: nil}}}

	got := addTwoNumbers(l1, l2)
	want := &ListNode{Val: 7, Next: &ListNode{Val: 0, Next: &ListNode{Val: 8, Next: nil}}}

	assert.Equal(t, got, want)
}

func TestAddTwoNumbers_2(t *testing.T) {
	l1 := &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: nil}}}}}}}
	l2 := &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: nil}}}}

	got := addTwoNumbers(l1, l2)
	want := &ListNode{Val: 8, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 0, Next: &ListNode{Val: 0, Next: &ListNode{Val: 0, Next: &ListNode{Val: 1, Next: nil}}}}}}}}

	assert.Equal(t, got, want)
}

func TestAddTwoNumbers_3(t *testing.T) {
	l1 := &ListNode{Val: 9, Next: &ListNode{Val: 8, Next: &ListNode{Val: 1, Next: nil}}}
	l2 := &ListNode{Val: 9, Next: nil}

	got := addTwoNumbers(l1, l2)
	want := &ListNode{Val: 8, Next: &ListNode{Val: 9, Next: &ListNode{Val: 1, Next: nil}}}

	assert.Equal(t, got, want)
}
