package main

import "fmt"

type IStack interface {
	Push(interface{})
	Pop() interface{}
	IsEmpty() bool
}

type StringIStack struct {
	stack []string
}

func (s *StringIStack) Push(v interface{}) {
	s.stack = append(s.stack, v.(string))
}

func (s *StringIStack) Pop() interface{} {
	if s.IsEmpty() {
		return nil
	}
	lastIndex := len(s.stack) - 1
	v := s.stack[lastIndex]
	s.stack = s.stack[:lastIndex]
	return v
}

func (s *StringIStack) IsEmpty() bool {
	return len(s.stack) == 0
}

func main() {
	s := &StringIStack{}
	s.Push("a")
	s.Push("b")
	s.Push("c")
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
}
