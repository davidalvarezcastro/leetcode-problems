package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{}
	tmp := result

	for l1 != nil || l2 != nil {
		if l1 != nil {
			tmp.Val += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			tmp.Val += l2.Val
			l2 = l2.Next
		}

		if tmp.Val > 9 {
			tmp.Val -= 10
			tmp.Next = &ListNode{Val: 1}
		} else {
			if l1 != nil || l2 != nil {
				tmp.Next = &ListNode{}
			}
		}

		tmp = tmp.Next
	}

	return result
}

func main() {
	l1 := &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: nil}}}}}}}
	l2 := &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: nil}}}}
	result := addTwoNumbers(l1, l2)

	for result != nil {
		fmt.Println(result.Val)
		result = result.Next
	}
}
