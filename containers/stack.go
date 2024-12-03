package containers

type Stack[T any] struct {
	Items []*T
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{}
}
func (s *Stack[T]) Push(item *T) {
	s.Items = append(s.Items, item)
}
func (s *Stack[T]) Pop() *T {
	if len(s.Items) == 0 {
		return nil
	}
	item := s.Items[len(s.Items)-1]
	s.Items = s.Items[:len(s.Items)-1]
	return item
}
func (s *Stack[T]) Peek() *T {
	if len(s.Items) == 0 {
		return nil
	}
	return s.Items[len(s.Items)-1]
}
func (s *Stack[T]) Clear() {
	s.Items = []*T{}
}
func (s *Stack[T]) Len() int {
	return len(s.Items)
}
func (s *Stack[T]) IsEmpty() bool {
	return len(s.Items) == 0
}
