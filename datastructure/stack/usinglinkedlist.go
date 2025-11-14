package main

import "errors"

type Node struct {
	data int
	next *Node
}

type StackList struct {
	head *Node
}

func NewStackList() *StackList {
	return &StackList{head: nil}
}

func (s *StackList) Push(v int) {
	newNode := &Node{data: v, next: s.head}
	s.head = newNode
}

func (s *StackList) Peek() (int, error) {
	if s.head == nil {
		return 0, errors.New("sjhbheb")
	}
	return s.head.data, nil
}

func (s *StackList) Pop() (int, error) {
	if s.head == nil {
		return 0, errors.New("sjhbheb")
	}
	val := s.head.data
	s.head = s.head.next
	return val,nil
}

func (s *StackList) Size() int {
	count := 0
	for cur := s.head; cur != nil ;cur = cur.next{
		count++
	}
	return count
}

func main() {

}
