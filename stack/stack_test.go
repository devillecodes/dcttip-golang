package stack

import (
	"testing"
)

func TestPush(t *testing.T) {
	tests := []struct {
		s    Stack
		item int
		want Stack
	}{
		{Stack{}, 1, Stack{1}},
		{Stack{1}, 3, Stack{1, 3}},
		{Stack{1, 3}, 5, Stack{1, 3, 5}},
	}

	for _, test := range tests {
		test.s.Push(test.item)

		got := test.s
		want := test.want

		if len(got) != len(want) {
			t.Errorf("%v.Push(%d): got %d, want %d", test.s, test.item, got, want)
		}

		for i, v := range got {
			if want[i] != v {
				t.Errorf("%v.Push(%d): got %d, want %d", test.s, test.item, got, want)
			}
		}
	}
}

func TestPop(t *testing.T) {
	tests := []struct {
		s    Stack
		item int
		ok   bool
	}{
		{Stack{1, 3, 5}, 5, true},
		{Stack{1, 3}, 3, true},
		{Stack{1}, 1, true},
		{Stack{}, 0, false},
	}

	for _, test := range tests {
		before := make(Stack, len(test.s))
		copy(before, test.s)

		got, ok := test.s.Pop()

		if ok != test.ok {
			t.Errorf("%v.Pop(): got %d, %t; want %d, %t", test.s, got, ok, test.item, test.ok)
		}

		if len(before) == 0 {
			// only care about contents when we start with a non-empty stack
			continue
		}

		if len(test.s) != len(before)-1 {
			t.Errorf("%v.Pop() should remove item: got %v, want %v", test.s, before, test.s)
		}

		for i, v := range test.s {
			if v != before[i] {
				t.Errorf("%v.Pop() should not alter remaining values: got %d, want %d", before, v, before[i])
			}
		}
	}
}

func TestPeek(t *testing.T) {
	tests := []struct {
		s    Stack
		item int
		ok   bool
	}{
		{Stack{1, 3, 5}, 5, true},
		{Stack{1, 3}, 3, true},
		{Stack{1}, 1, true},
		{Stack{}, 0, false},
	}

	for _, test := range tests {
		before := make(Stack, len(test.s))
		copy(before, test.s)

		got, ok := test.s.Peek()

		if ok != test.ok {
			t.Errorf("%v.Peek(): got %d, %t; want %d, %t", test.s, got, ok, test.item, test.ok)
		}

		if len(before) == 0 {
			// only care about contents when we start with a non-empty stack
			continue
		}

		if len(test.s) != len(before) {
			t.Errorf("%v.Peek() should not remove item: got %v, want %v", before, test.s, before)
			continue
		}

		for i, v := range test.s {
			if v != before[i] {
				t.Errorf("s.Peek(%v) should not alter values, got %d, want %d", before, v, before[i])
			}
		}
	}
}

func TestEmpty(t *testing.T) {
	tests := []struct {
		s    Stack
		want bool
	}{
		{Stack{}, true},
		{Stack{1}, false},
		{Stack{1, 3}, false},
		{Stack{1, 3, 5}, false},
	}

	for _, test := range tests {
		got := test.s.Empty()
		if got != test.want {
			t.Errorf("%v.Empty(): got %t, want %t", test.s, got, test.want)
		}
	}
}

func TestLength(t *testing.T) {
	tests := []struct {
		s    Stack
		want int
	}{
		{Stack{}, 0},
		{Stack{1}, 1},
		{Stack{1, 3}, 2},
		{Stack{1, 3, 5}, 3},
	}

	for _, test := range tests {
		got := test.s.Length()
		if got != test.want {
			t.Errorf("%v.Length(): got %d, want %d", test.s, got, test.want)
		}
	}
}
