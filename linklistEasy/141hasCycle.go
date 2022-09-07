package main

//快慢指针
func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	slow, fast := head, head.Next
	for slow != fast {
		if fast == nil || fast.Next == nil {
			return false
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return true
}

//哈希表
//func hasCycle(head *ListNode) bool{
//	seen := map[*ListNode]struct{}{}
//	for head != nil{
//		if _,ok :=seen[head];ok{
//			return true
//		}
//		seen[head] = struct{}{}
//		head = head.Next
//	}
//	return false
//}
func main() {

}
