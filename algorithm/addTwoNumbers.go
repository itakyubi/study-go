package main

func main() {

}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	cur := head
	carry := 0

	for l1 != nil || l2 != nil || carry > 0 {
		sum := carry
		if l1 != nil {
			sum += l1.Val
		}
		if l2 != nil {
			sum += l2.Val
		}
		if sum < 10 {
			carry = 0
		} else {
			carry = 1
			sum -= 10
		}

		next := &ListNode{Val: sum}
		cur.Next = next
		cur = next

		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}

	return head.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}
