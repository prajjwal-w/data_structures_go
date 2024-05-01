package linkedList

import "fmt"

type Node struct {
	Data int
	Next *Node
}

func CreateLinkedList(data []int) *Node {
	var head *Node
	for _, value := range data {
		newNode := &Node{Data: value}
		if head == nil {
			head = newNode
		} else {
			curr := head
			for curr.Next != nil {
				curr = curr.Next
			}
			curr.Next = newNode
		}
	}
	return head
}

func PrintLinkedList(head *Node) {
	curr := head
	for curr != nil {
		fmt.Print(curr.Data, " -> ")
		curr = curr.Next
	}
	fmt.Print("nil\n")
}

//use
