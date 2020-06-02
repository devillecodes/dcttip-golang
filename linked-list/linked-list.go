package list

// Node represents a single item in the list.
type Node struct {
	v    int
	next *Node
}

// LinkedList is a collection of linked nodes.
type LinkedList struct {
	root *Node
}

// Push adds an element to the list.
func (l *LinkedList) Push(item int) {
	if l.root == nil {
		l.root = &Node{v: item}
		return
	}

	last := l.root
	for last.next != nil {
		last = last.next
	}

	last.next = &Node{v: item}
}

// Pop removes an element from the list.
func (l *LinkedList) Pop() (int, bool) {
	if l.root == nil {
		return 0, false
	}

	if l.root.next == nil {
		v := l.root.v
		l.root = nil
		return v, true
	}

	var prev *Node
	curr := l.root
	for curr.next != nil {
		prev = curr
		curr = curr.next
	}

	prev.next = nil
	return curr.v, true
}

// Get returns the element at the given index.
func (l *LinkedList) Get(index int) (int, bool) {
	if index < 0 {
		return 0, false
	}

	if l.root == nil {
		return 0, false
	}

	curr := l.root
	i := 0
	for ; i < index; i++ {
		if curr.next == nil {
			break
		}
		curr = curr.next
	}

	if i != index {
		// index wasn't found
		return 0, false
	}

	return curr.v, true
}

// Delete removes the element at the given index.
func (l *LinkedList) Delete(index int) bool {
	if index < 0 {
		return false
	}

	if l.root == nil {
		return false
	}

	if index == 0 && l.root.next == nil {
		l.root = nil
		return true
	}

	var prev *Node
	curr := l.root
	i := 0
	for ; i < index; i++ {
		if curr.next == nil {
			break
		}
		prev = curr
		curr = curr.next
	}

	if i != index {
		// index wasn't found
		return false
	}

	prev.next = nil
	return true
}

// Empty returns a boolean indicating whether the list is empty.
func (l *LinkedList) Empty() bool {
	return l.root == nil
}

// Count returns the number of items.
func (l *LinkedList) Count() int {
	if l.root == nil {
		return 0
	}

	count := 1
	next := l.root.next
	for next != nil {
		count++
		next = next.next
	}

	return count
}
