package main

func main() {

}

func addTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {
	rev1 := reverseListNode(l1)
	rev2 := reverseListNode(l2)

	head := &ListNode{}
	carry := 0
	for rev1 != nil || rev2 != nil || carry > 0 {
		sum := carry
		if rev1 != nil {
			sum += rev1.Val
		}
		if rev2 != nil {
			sum += rev2.Val
		}
		if sum < 10 {
			carry = 0
		} else {
			carry = 1
			sum -= 10
		}

		next := &ListNode{Val: sum}
		next.Next = head.Next
		head.Next = next

		if rev1 != nil {
			rev1 = rev1.Next
		}
		if rev2 != nil {
			rev2 = rev2.Next
		}
	}
	return head.Next
}

func reverseListNode(l *ListNode) *ListNode {
	head := &ListNode{}
	head.Next = l
	cur := head.Next
	next := &ListNode{}

	for cur != nil && cur.Next != nil {
		next = cur.Next
		cur.Next = next.Next
		next.Next = head.Next
		head.Next = next
	}
	return head.Next
}

func reverseList(head *ListNode) *ListNode {
	var pre, cur *ListNode = nil, head
	for cur != nil {
		nxt := cur.Next
		cur.Next = pre
		pre = cur
		cur = nxt
	}
	return pre
}
