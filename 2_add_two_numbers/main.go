package main

import "fmt"

// TODO: FINISH!

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) Print() {
	arr := []int{}

	node := l
	for node != nil {
		arr = append(arr, node.Val)
		node = node.Next
	}

	fmt.Println(arr)
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	response := ListNode{}

	var sum int
	first := l1
	second := l2
	node := &response
	for {
		var val1, val2 int

		if first != nil {
			val1 = first.Val
		}
		if second != nil {
			val2 = second.Val
		}
		if first == nil && second == nil && sum == 0 {
			break
		}

		sum = val1 + val2
		fmt.Println(sum)

		if sum+node.Val >= 9 {
			node.Val += 9 - node.Val
		} else {
			node.Val += sum
		}

		if sum >= 9 {
			sum = sum - 9
		} else {
			sum = 0
		}

		if first != nil {
			first = first.Next
		}

		if second != nil {
			second = second.Next
		}

		if first == nil && second == nil && sum == 0 {
			break
		}

		node.Next = &ListNode{}
		node = node.Next
	}

	return &response
}

func main() {
	example := ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 1,
			},
		},
	}

	example1 := ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 1,
			},
		},
	}

	result := addTwoNumbers(&example, &example1)

	result.Print()

	example = ListNode{
		Val: 9,
		Next: &ListNode{
			Val: 9,
			Next: &ListNode{
				Val: 9,
			},
		},
	}

	example1 = ListNode{
		Val: 9,
		Next: &ListNode{
			Val: 9,
			Next: &ListNode{
				Val: 9,
			},
		},
	}

	result = addTwoNumbers(&example, &example1)

	result.Print()

}
