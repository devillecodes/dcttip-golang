package graph

import "testing"

func TestAdd(t *testing.T) {
	tests := []struct {
		desc string
		g    Graph
		n    *Node
		want Graph
	}{
		{
			"nil node",
			Graph{nodes: []*Node{{value: 5}}},
			nil,
			Graph{nodes: []*Node{{value: 5}}},
		},
		{
			"empty",
			Graph{},
			&Node{value: 5},
			Graph{nodes: []*Node{{value: 5}}},
		},
		{
			"non-empty",
			Graph{nodes: []*Node{{value: 1}, {value: 3}}},
			&Node{value: 5},
			Graph{nodes: []*Node{{value: 1}, {value: 3}, {value: 5}}},
		},
	}

	for _, test := range tests {
		test.g.Add(test.n)

		got := test.g
		want := test.want

		lg := len(got.nodes)
		lw := len(want.nodes)
		if lg != lw {
			t.Errorf("%v: got %d node(s), want %d node(s)", test.desc, lg, lw)
			continue
		}

		for i, v := range want.nodes {
			got := got.nodes[i].value
			want := v.value
			if got != want {
				t.Errorf("%v: got %d, want %d", test.desc, got, want)
			}
		}
	}
}

func TestRemove(t *testing.T) {
	tests := []struct {
		desc string
		g    Graph
		v    int
		want Graph
	}{
		{
			"empty",
			Graph{},
			5,
			Graph{},
		},
		{
			"one node",
			Graph{nodes: []*Node{{value: 5}}},
			5,
			Graph{},
		},
		{
			"three nodes; no edges",
			Graph{nodes: []*Node{{value: 1}, {value: 3}, {value: 5}}},
			3,
			Graph{nodes: []*Node{{value: 1}, {value: 5}}},
		},
		{
			"three nodes; undirected; with edges",
			Graph{
				nodes: []*Node{
					{value: 1, edges: []*Node{{value: 3}, {value: 5}}},
					{value: 3, edges: []*Node{{value: 1}, {value: 5}}},
					{value: 5, edges: []*Node{{value: 1}, {value: 3}}},
				},
			},
			3,
			Graph{
				nodes: []*Node{
					{value: 1, edges: []*Node{{value: 5}}},
					{value: 5, edges: []*Node{{value: 1}}},
				},
			},
		},
		{
			"three nodes; directed; with edges",
			Graph{
				directed: true, nodes: []*Node{
					{value: 1, edges: []*Node{{value: 3}, {value: 5}}},
					{value: 3, edges: []*Node{{value: 5}}},
					{value: 5},
				},
			},
			3,
			Graph{
				directed: true, nodes: []*Node{
					{value: 1, edges: []*Node{{value: 5}}},
					{value: 5},
				},
			},
		},
	}

	for _, test := range tests {
		test.g.Remove(test.v)

		got := test.g
		want := test.want

		lg := len(got.nodes)
		lw := len(want.nodes)
		if lg != lw {
			t.Errorf("%v: got %d node(s), want %d node(s)", test.desc, lg, lw)
			continue
		}

		for i, vn := range want.nodes {
			got := got.nodes[i]
			want := vn
			if got.value != want.value {
				t.Errorf("%v: got %d, want %d", test.desc, got.value, want.value)
				continue
			}

			ge := got.edges
			we := want.edges
			if len(ge) != len(we) {
				t.Errorf("%v: got %d edge(s), want %d edge(s)", test.desc, len(ge), len(we))
				continue
			}

			for i, ve := range we {
				got := ge[i].value
				want := ve.value
				if got != want {
					t.Errorf("%v: got %d, want %d", test.desc, got, want)
				}
			}
		}
	}
}

func TestNode(t *testing.T) {
	tests := []struct {
		desc string
		n    *Node
		v    int
		ok   bool
	}{
		{
			"found",
			&Node{value: 5},
			5,
			true,
		},
		{
			"not found",
			&Node{value: 5},
			3,
			false,
		},
	}

	for _, test := range tests {
		g := Graph{nodes: []*Node{test.n}}

		got, ok := g.Node(test.v)

		if ok != test.ok {
			t.Errorf("%v: got _, %t; want _, %t", test.desc, ok, test.ok)
		}

		if !test.ok {
			// we don't care about value when ok is false
			continue
		}

		if got != test.n {
			t.Errorf("%v: got %v, want %v", test.desc, got, test.n)
		}
	}
}

func TestAddEdge(t *testing.T) {
	tests := []struct {
		desc   string
		g      Graph
		v1, v2 int
		want   Graph
	}{
		{
			"empty",
			Graph{},
			1, 5,
			Graph{},
		},
		{
			"one node",
			Graph{nodes: []*Node{{value: 5}}},
			1, 5,
			Graph{nodes: []*Node{{value: 5}}},
		},
		{
			"three nodes; undirected",
			Graph{
				nodes: []*Node{
					{value: 1, edges: []*Node{{value: 3}}},
					{value: 3, edges: []*Node{{value: 1}, {value: 5}}},
					{value: 5, edges: []*Node{{value: 3}}},
				},
			},
			1, 5,
			Graph{
				nodes: []*Node{
					{value: 1, edges: []*Node{{value: 3}, {value: 5}}},
					{value: 3, edges: []*Node{{value: 1}, {value: 5}}},
					{value: 5, edges: []*Node{{value: 3}, {value: 1}}},
				},
			},
		},
		{
			"three nodes; directed",
			Graph{
				directed: true, nodes: []*Node{
					{value: 1, edges: []*Node{{value: 3}}},
					{value: 3, edges: []*Node{{value: 5}}},
					{value: 5},
				},
			},
			1, 5,
			Graph{
				directed: true, nodes: []*Node{
					{value: 1, edges: []*Node{{value: 3}, {value: 5}}},
					{value: 3, edges: []*Node{{value: 5}}},
					{value: 5},
				},
			},
		},
	}

	for _, test := range tests {
		test.g.AddEdge(test.v1, test.v2)

		got := test.g
		want := test.want

		lg := len(got.nodes)
		lw := len(want.nodes)
		if lg != lw {
			t.Errorf("%v: got %d node(s), want %d node(s)", test.desc, lg, lw)
			continue
		}

		for i, vn := range want.nodes {
			got := got.nodes[i]
			want := vn
			if got.value != want.value {
				t.Errorf("%v: got %d, want %d", test.desc, got.value, want.value)
				continue
			}

			ge := got.edges
			we := want.edges
			if len(ge) != len(we) {
				t.Errorf("%v: got %d edge(s), want %d edge(s)", test.desc, len(ge), len(we))
				continue
			}

			for i, ve := range we {
				got := ge[i].value
				want := ve.value
				if got != want {
					t.Errorf("%v: got %d, want %d", test.desc, got, want)
				}
			}
		}
	}
}
