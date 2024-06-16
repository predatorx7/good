package main

import (
	"fmt"
	"good/pkg/add_two_numbers"
)

func printListNode(n *add_two_numbers.ListNode) {
	current := n
	for current != nil {
		fmt.Printf("%d", current.Val)
		current = current.Next
		if current != nil {
			fmt.Print(", ")
		} else {
			fmt.Print("\n")
		}
	}
}
func sample1() {
	a := &add_two_numbers.ListNode{
		Val: 2,
		Next: &add_two_numbers.ListNode{
			Val: 4,
			Next: &add_two_numbers.ListNode{
				Val: 3,
			},
		},
	}
	b := &add_two_numbers.ListNode{
		Val: 5,
		Next: &add_two_numbers.ListNode{
			Val: 6,
			Next: &add_two_numbers.ListNode{
				Val: 4,
			},
		},
	}

	result := add_two_numbers.Solve(a, b)

	printListNode(result)
}

func sample2() {
	a := &add_two_numbers.ListNode{
		Val: 9,
		Next: &add_two_numbers.ListNode{
			Val: 9,
			Next: &add_two_numbers.ListNode{
				Val: 9,
			},
		},
	}
	b := &add_two_numbers.ListNode{
		Val: 9,
		Next: &add_two_numbers.ListNode{
			Val: 9,
		},
	}

	result := add_two_numbers.Solve(a, b)

	printListNode(result)
}
func main() {
	sample1()

	sample2()
}
