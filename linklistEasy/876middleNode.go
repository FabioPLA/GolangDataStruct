package main

func middleNode(head *ListNode) *ListNode {
	newHead := &ListNode{Next: head}
	slow, fast := newHead, newHead
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	if fast == nil {
		return slow
	}
	return slow.Next
}
func main() {

}
