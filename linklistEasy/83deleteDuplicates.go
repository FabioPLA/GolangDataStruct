package main

import "fmt"

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	cur := head
	for cur.Next != nil {
		if cur.Val == cur.Next.Val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return head
}
func main() {
	head := &ListNode{}
	head.Val = 1
	head.Next.Val = 1
	head.Next.Next.Val = 2
	fmt.Println(deleteDuplicates(head))
}
