package tree

import (
	"fmt"
	"strings"
)

// Node represents a node in the tree.
type Node struct {
	value       int
	left, right *Node
}

func (n *Node) String() string {
	var b strings.Builder

	fmt.Fprintf(&b, "[v:%d", n.value)
	if n.left != nil {
		fmt.Fprintf(&b, ", left: %s", n.left)
	}
	if n.right != nil {
		fmt.Fprintf(&b, ", right: %s", n.right)
	}
	fmt.Fprintf(&b, "]")

	return b.String()
}

// Tree represents a Binary Search Tree.
type Tree struct {
	root *Node
}

func (t *Tree) String() string {
	if t.root == nil {
		return "[empty]"
	}
	return fmt.Sprintf("%s", t.root)
}

// Add a node to the tree.
// Returns true when the node could be added.
func (t *Tree) Add(value int) bool {
	node := &Node{value: value}

	if t.root == nil {
		t.root = node
		return true
	}

	current := t.root
	for {
		if current.value == value {
			// duplicate node found
			return false
		}

		if value < current.value {
			if current.left == nil {
				current.left = node
				return true
			}
			current = current.left
			continue
		}

		if current.right == nil {
			current.right = node
			return true
		}
		current = current.right
	}
}

// Remove a node from the tree.
// Returns true when the node could be deleted.
func (t *Tree) Remove(value int) bool {
	// TODO:
	return false
}

func nodeEqual(n1, n2 *Node) bool {
	switch {
	case n1 == nil && n2 == nil:
		return true
	case n1 == nil || n2 == nil:
		return false
	case n1.value != n2.value:
		return false
	}

	left := nodeEqual(n1.left, n2.left)
	right := nodeEqual(n1.right, n2.right)

	return left && right
}

// Equal returns true if the structure matches the provided tree.
func (t *Tree) Equal(tree *Tree) bool {
	if t.root == nil && tree.root == nil {
		return true
	}

	if t.root == nil || tree.root == nil {
		return false
	}

	return nodeEqual(t.root, tree.root)
}
