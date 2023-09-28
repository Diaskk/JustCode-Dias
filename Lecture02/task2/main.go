package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

	l1 := &ListNode{Val: 10, Next: &ListNode{Val: 1000, Next: &ListNode{Val: 1337}}}
	l2 := &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 6}}}

	printList(mergeTwoLists(l1, l2))
}
func printList(head *ListNode) {
	for head != nil {
		fmt.Print(head.Val)
		head = head.Next
		if head != nil {
			fmt.Print(", ")
		}
	}
	fmt.Println()
}
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	temp := &ListNode{}
	current := temp

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			current.Next = l1
			l1 = l1.Next
		} else {
			current.Next = l2
			l2 = l2.Next
		}
		current = current.Next
	}
	if l1 != nil {
		current.Next = l1
	}
	if l2 != nil {
		current.Next = l2
	}

	return temp.Next
}
