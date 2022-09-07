package main

func mergeTwoListsII(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil && list2 == nil {
		return nil
	}
	if (list1 == nil && list2 != nil) || (list2 == nil && list1 != nil) {
		if list1 != nil {
			return list1
		} else {
			return list2
		}
	}
	dummy := &ListNode{0, nil}
	temp := dummy
	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			temp.Next = list1
			list1 = list1.Next
		} else {
			temp.Next = list2
			list2 = list2.Next
		}
		temp = temp.Next
	}
	if list1 != nil {
		temp.Next = list1
	}
	if list2 != nil {
		temp.Next = list2
	}
	return dummy.Next
}
func main() {

}
