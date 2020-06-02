package queue

import (
	"testing"
)

func TestEnqueue(t *testing.T) {
	tests := []struct {
		q    Queue
		item int
		want Queue
	}{
		{Queue{}, 1, Queue{1}},
		{Queue{1}, 3, Queue{1, 3}},
		{Queue{1, 3}, 5, Queue{1, 3, 5}},
	}

	for _, test := range tests {
		test.q.Enqueue(test.item)

		got := test.q
		want := test.want

		if len(got) != len(want) {
			t.Errorf("%v.Push(%d): got %d, want %d", test.q, test.item, got, want)
		}

		for i, v := range got {
			if want[i] != v {
				t.Errorf("%v.Push(%d): got %d, want %d", test.q, test.item, got, want)
			}
		}
	}
}

func TestDequeue(t *testing.T) {
	tests := []struct {
		q    Queue
		item int
		ok   bool
	}{
		{Queue{1, 3, 5}, 1, true},
		{Queue{3, 5}, 3, true},
		{Queue{5}, 5, true},
		{Queue{}, 0, false},
	}

	for _, test := range tests {
		before := make(Queue, len(test.q))
		copy(before, test.q)

		got, ok := test.q.Dequeue()

		if ok != test.ok {
			t.Errorf("%v.Dequeue(): got %d, %t; want %d, %t", test.q, got, ok, test.item, test.ok)
		}

		if len(before) == 0 {
			// only care about contents when we start with a non-empty Queue
			continue
		}

		if len(test.q) != len(before)-1 {
			t.Errorf("%v.Dequeue() should remove item: got %v, want %v", before, test.q, before[1:])
			continue
		}

		for i, v := range test.q {
			if v != before[i+1] {
				t.Errorf("%v.Dequeue() should not alter values, got %d, want %d", before, v, before[i])
			}
		}
	}
}

func TestPeek(t *testing.T) {
	tests := []struct {
		q    Queue
		item int
		ok   bool
	}{
		{Queue{1, 3, 5}, 1, true},
		{Queue{3, 5}, 3, true},
		{Queue{5}, 5, true},
		{Queue{}, 0, false},
	}

	for _, test := range tests {
		before := make(Queue, len(test.q))
		copy(before, test.q)

		got, ok := test.q.Peek()

		if ok != test.ok {
			t.Errorf("%v.Peek(): got %d, %t; want %d, %t", test.q, got, ok, test.item, test.ok)
		}

		if len(before) == 0 {
			// only care about contents when we start with a non-empty Queue
			continue
		}

		if len(test.q) != len(before) {
			t.Errorf("%v.Peek() should not remove item: got %v, want %v", before, test.q, before)
			continue
		}

		for i, v := range test.q {
			if v != before[i] {
				t.Errorf("s.Peek(%v) should not alter values, got %d, want %d", before, v, before[i])
			}
		}
	}
}

func TestEmpty(t *testing.T) {
	tests := []struct {
		s    Queue
		want bool
	}{
		{Queue{}, true},
		{Queue{1}, false},
		{Queue{1, 3}, false},
		{Queue{1, 3, 5}, false},
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
		s    Queue
		want int
	}{
		{Queue{}, 0},
		{Queue{1}, 1},
		{Queue{1, 3}, 2},
		{Queue{1, 3, 5}, 3},
	}

	for _, test := range tests {
		got := test.s.Length()
		if got != test.want {
			t.Errorf("%v.Length(): got %d, want %d", test.s, got, test.want)
		}
	}
}
