package main

// import (
// 	"errors"
// 	"fmt"
// )

// type StackSlice struct {
// 	data []int
// }

// func NewStackSlice() *StackSlice {
// 	return &StackSlice{data: []int{}}
// }

// func (s *StackSlice) Push(v int) {
// 	s.data = append(s.data, v)
// }

// func (s *StackSlice) Pop() (int, error) {
// 	if len(s.data) == 0 {
// 		return 0, errors.New("stack is empty")
// 	}
// 	i := len(s.data) - 1
// 	v := s.data[i]
// 	s.data = s.data[:i]
// 	return v, nil
// }

// func (s *StackSlice) Peek() (int, error) {
// 	if len(s.data) == 0 {
// 		return 0, errors.New("stack is empty")
// 	}
// 	return s.data[len(s.data)-1], nil
// }

// func (s *StackSlice) IsEmpty() bool {
// 	return len(s.data) == 0
// }

// func (s *StackSlice) Size() int {
// 	return len(s.data)
// }

// func main() {
// 	ss := NewStackSlice()
// 	ss.Push(10)
// 	ss.Push(20)
// 	ss.Push(30)
// 	ss.Push(40)
// 	t := ss.Size()
// 	fmt.Println(t)
// }
