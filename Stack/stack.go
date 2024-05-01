package stack

import "fmt"

//Stack represents a stack data structure
type Stack struct {
	items []interface{}
}

//NewStack creates a new empty stack
func NewStack() *Stack {
	return &Stack{items: []interface{}{}}
}

//Push adds an element to the top of the stack
func (s *Stack) Push(item interface{}) {
	s.items = append(s.items, item)
}

//Pop removes and returns the top element from the stack
func (s *Stack) Pop() (interface{}, error) {
	if len(s.items) == 0 {
		return nil, fmt.Errorf("Stack is empty")
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item, nil
}

//len returns the number of elements in the stack
func (s *Stack) Len() int {
	return len(s.items)
}

//PrintStack prints the stack
func (s *Stack) PrintStack() {
	for _, value := range s.items {
		fmt.Print(value, " ")
	}
}
