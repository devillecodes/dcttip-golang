package list

import "testing"

func last(node *Node) Node {
	if node.next == nil {
		return *node
	}
	return last(node.next)
}

func TestPush(t *testing.T) {
	tests := []struct {
		desc string
		list LinkedList
		item int
		want LinkedList
	}{
		{
			"empty list",
			LinkedList{},
			5,
			LinkedList{&Node{v: 5}},
		},
		{
			"non-empty list",
			LinkedList{&Node{v: 5}},
			7,
			LinkedList{&Node{v: 5, next: &Node{v: 7}}},
		},
	}

	for _, test := range tests {
		test.list.Push(test.item)
		got := last(test.list.root).v
		want := last(test.want.root).v
		if got != want {
			t.Errorf("%s: got %d, want %d", test.desc, got, want)
		}
	}
}

func TestPop(t *testing.T) {
	tests := []struct {
		desc  string
		list  LinkedList
		after LinkedList
		v     int
		ok    bool
	}{
		{
			"empty",
			LinkedList{},
			LinkedList{},
			0,
			false,
		},
		{
			"one item",
			LinkedList{&Node{v: 5}},
			LinkedList{},
			5,
			true,
		},
		{
			"two items",
			LinkedList{&Node{v: 1, next: &Node{v: 3}}},
			LinkedList{&Node{v: 1}},
			3,
			true,
		},
		{
			"three items",
			LinkedList{&Node{v: 1, next: &Node{v: 3, next: &Node{v: 5}}}},
			LinkedList{&Node{v: 1, next: &Node{v: 3}}},
			5,
			true,
		},
	}

	for _, test := range tests {
		gv, gok := test.list.Pop()
		wv, wok := test.v, test.ok
		if gv != wv || gok != wok {
			t.Errorf("%s: got (%d, %t), want (%d, %t)", test.desc, gv, gok, wv, wok)
		}

		if test.after.root == nil {
			// only care about contents when we expect a non-empty list
			continue
		}

		got := last(test.list.root)
		want := last(test.after.root)
		if got != want {
			t.Errorf("%s: got %v, want %v", test.desc, got, want)
		}

		if test.after.root.next == nil {
			// only care about contents when we expect a list with more than one item
			continue
		}

		got = last(test.list.root)
		want = last(test.after.root)
		if got != want {
			t.Errorf("%s: got %v, want %v", test.desc, got, want)
		}
	}
}

func TestGet(t *testing.T) {

	tests := []struct {
		desc  string
		index int
		list  LinkedList
		after LinkedList
		v     int
		ok    bool
	}{
		{
			"empty",
			0,
			LinkedList{},
			LinkedList{},
			0,
			false,
		},
		{
			"index out of bounds",
			2,
			LinkedList{&Node{v: 5}},
			LinkedList{&Node{v: 5}},
			0,
			false,
		},
		{
			"negative index",
			-1,
			LinkedList{&Node{v: 5}},
			LinkedList{&Node{v: 5}},
			0,
			false,
		},
		{
			"one item",
			0,
			LinkedList{&Node{v: 5}},
			LinkedList{&Node{v: 5}},
			5,
			true,
		},
		{
			"two items",
			1,
			LinkedList{&Node{v: 1, next: &Node{v: 3}}},
			LinkedList{&Node{v: 1, next: &Node{v: 3}}},
			3,
			true,
		},
		{
			"three items",
			2,
			LinkedList{&Node{v: 1, next: &Node{v: 3, next: &Node{v: 5}}}},
			LinkedList{&Node{v: 1, next: &Node{v: 3, next: &Node{v: 5}}}},
			5,
			true,
		},
	}

	for _, test := range tests {
		gv, gok := test.list.Get(test.index)
		wv, wok := test.v, test.ok
		if gv != wv || gok != wok {
			t.Errorf("%s: got (%d, %t), want (%d, %t)", test.desc, gv, gok, wv, wok)
		}

		if test.after.root == nil {
			// only care about contents when we expect a non-empty list
			continue
		}

		got := last(test.list.root)
		want := last(test.after.root)
		if got != want {
			t.Errorf("%s: got %v, want %v", test.desc, got, want)
		}

		if test.after.root.next == nil {
			// only care about contents when we expect a list with more than one item
			continue
		}

		got = last(test.list.root)
		want = last(test.after.root)
		if got != want {
			t.Errorf("%s: got %v, want %v", test.desc, got, want)
		}
	}
}

func TestDelete(t *testing.T) {
	tests := []struct {
		desc  string
		index int
		list  LinkedList
		after LinkedList
		ok    bool
	}{
		{
			"empty",
			0,
			LinkedList{},
			LinkedList{},
			false,
		},
		{
			"index out of bounds",
			2,
			LinkedList{&Node{v: 5}},
			LinkedList{&Node{v: 5}},
			false,
		},
		{
			"negative index",
			-1,
			LinkedList{&Node{v: 5}},
			LinkedList{&Node{v: 5}},
			false,
		},
		{
			"one item",
			0,
			LinkedList{&Node{v: 5}},
			LinkedList{},
			true,
		},
		{
			"two items",
			1,
			LinkedList{&Node{v: 1, next: &Node{v: 3}}},
			LinkedList{&Node{v: 1}},
			true,
		},
		{
			"three items",
			2,
			LinkedList{&Node{v: 1, next: &Node{v: 3, next: &Node{v: 5}}}},
			LinkedList{&Node{v: 1, next: &Node{v: 3}}},
			true,
		},
	}

	for _, test := range tests {
		got := test.list.Delete(test.index)
		want := test.ok
		if got != want {
			t.Errorf("%s: got %v, want %v", test.desc, got, want)
		}

		if test.after.root == nil {
			// only care about contents when we expect a non-empty list
			continue
		}

		gl := last(test.list.root)
		wl := last(test.after.root)
		if gl != wl {
			t.Errorf("%s: got %v, want %v", test.desc, gl, wl)
		}

		if test.after.root.next == nil {
			// only care about contents when we expect a list with more than one item
			continue
		}

		gl = last(test.list.root)
		wl = last(test.after.root)
		if gl != wl {
			t.Errorf("%s: got %v, want %v", test.desc, gl, wl)
		}
	}
}

func TestEmpty(t *testing.T) {
	tests := []struct {
		desc string
		list LinkedList
		want bool
	}{
		{
			"empty list",
			LinkedList{},
			true,
		},
		{
			"non-empty list",
			LinkedList{&Node{v: 5}},
			false,
		},
	}

	for _, test := range tests {
		got := test.list.Empty()
		want := test.want
		if got != want {
			t.Errorf("%s: got %t, want %t", test.desc, got, want)
		}
	}
}

func TestCount(t *testing.T) {
	tests := []struct {
		desc string
		list LinkedList
		want int
	}{
		{
			"empty",
			LinkedList{},
			0,
		},
		{
			"one item",
			LinkedList{&Node{v: 5}},
			1,
		},
		{
			"two items",
			LinkedList{&Node{v: 1, next: &Node{v: 3}}},
			2,
		},
		{
			"three items",
			LinkedList{&Node{v: 1, next: &Node{v: 3, next: &Node{v: 5}}}},
			3,
		},
	}

	for _, test := range tests {
		got := test.list.Count()
		want := test.want
		if got != want {
			t.Errorf("%s: got %d, want %d", test.desc, got, want)
		}
	}
}
