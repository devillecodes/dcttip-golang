package stack

// Stack represents a stack of integers.
type Stack []int

// Push adds an item onto the stack.
func (s *Stack) Push(item int) {
	*s = append(*s, item)
}

// Pop removes the next item from the stack.
func (s *Stack) Pop() (int, bool) {
	l := len(*s)

	if l == 0 {
		return 0, false
	}

	i := (*s)[l-1]
	*s = (*s)[:len(*s)-1]

	return i, true
}

// Peek returns the next item on the stack without removing it.
func (s *Stack) Peek() (int, bool) {
	l := len(*s)

	if l == 0 {
		return 0, false
	}

	i := (*s)[l-1]

	return i, true
}

// Empty indicates whether the stack contains any items.
func (s *Stack) Empty() bool {
	return len(*s) == 0
}

// Length returns the number of items on the stack.
func (s *Stack) Length() int {
	return len(*s)
}
