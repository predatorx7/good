package add_two_numbers

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func Solve(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{Val: 0}
	current := result
	n1 := l1
	n2 := l2
	carry := 0

	for {
		if n1 != nil {
			current.Val += n1.Val
			n1 = n1.Next
		}
		if n2 != nil {
			current.Val += n2.Val
			n2 = n2.Next
		}
		if carry > 0 {
			current.Val += carry
			carry = 0
		}
		if current.Val > 9 {
			carry = 1
			current.Val = current.Val % 10
		}
		if n1 != nil || n2 != nil || carry > 0 {
			current.Next = &ListNode{Val: 0}
			current = current.Next
		} else {
			break
		}
	}

	return result
}
