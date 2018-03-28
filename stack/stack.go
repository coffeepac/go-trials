package stack

type sNode struct {
	Data interface{}
	Next *sNode
}

type Stack struct {
	Head *sNode
	Tail *sNode
}
type StackOps interface {
	Push(interface{}) bool
	Pop(interface{}) *sNode
}

func (s *Stack) Push(n interface{}) bool {
	last := &sNode{Data: n}
	if s.Tail == nil && s.Head == nil {
		s.Head = last
		s.Tail = last
		return true
	}
	s.Tail.Next = last
	s.Tail = last
	return true
}

func (s *Stack) Pop() *sNode {
	next := s.Head
	s.Head = s.Head.Next
	if s.Head == nil {
		s.Tail = nil
	}
	return next
}
