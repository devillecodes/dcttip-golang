package queue

// Queue represents a queue of integers.
type Queue []int

// Enqueue adds an item to the back of the queue.
func (q *Queue) Enqueue(item int) {
	*q = append(*q, item)
}

// Dequeue removes the first item from the queue and returns it.
func (q *Queue) Dequeue() (int, bool) {
	if len(*q) == 0 {
		return 0, false
	}

	i := (*q)[0]
	*q = (*q)[1:]

	return i, true
}

// Peek returns the first item in the queue without removing it.
func (q *Queue) Peek() (int, bool) {
	l := len(*q)

	if l == 0 {
		return 0, false
	}

	i := (*q)[0]

	return i, true
}

// Empty indicates whether the stack contains any items.
func (q *Queue) Empty() bool {
	return len(*q) == 0
}

// Length returns the number of items on the stack.
func (q *Queue) Length() int {
	return len(*q)
}
