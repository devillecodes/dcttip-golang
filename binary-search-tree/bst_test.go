package tree

import "testing"

func TestEqual(t *testing.T) {
	tests := []struct {
		desc string
		t1   *Tree
		t2   *Tree
		want bool
	}{
		{
			"empty",
			&Tree{},
			&Tree{},
			true,
		},
		{
			"one root nil",
			&Tree{},
			&Tree{root: &Node{value: 5}},
			false,
		},
		{
			"other root nil",
			&Tree{root: &Node{value: 5}},
			&Tree{},
			false,
		},
		{
			"only root equal",
			&Tree{root: &Node{value: 5}},
			&Tree{root: &Node{value: 5}},
			true,
		},
		{
			"only root not equal",
			&Tree{root: &Node{value: 5}},
			&Tree{root: &Node{value: 7}},
			false,
		},
		{
			"nested child nil",
			&Tree{root: &Node{value: 5, left: &Node{value: 2}}},
			&Tree{root: &Node{value: 5, left: &Node{value: 2}, right: &Node{value: 7}}},
			false,
		},
		{
			"nested other child nil",
			&Tree{root: &Node{value: 5, left: &Node{value: 2}, right: &Node{value: 7}}},
			&Tree{root: &Node{value: 5, left: &Node{value: 2}}},
			false,
		},
		{
			"nested equal",
			&Tree{root: &Node{value: 5, left: &Node{value: 2}, right: &Node{value: 7}}},
			&Tree{root: &Node{value: 5, left: &Node{value: 2}, right: &Node{value: 7}}},
			true,
		},
		{
			"nested not equal",
			&Tree{root: &Node{value: 5, left: &Node{value: 2}, right: &Node{value: 7}}},
			&Tree{root: &Node{value: 5, left: &Node{value: 3}, right: &Node{value: 7}}},
			false,
		},
	}

	for _, test := range tests {
		got := test.t1.Equal(test.t2)
		want := test.want

		if got != want {
			t.Errorf("%v: got %t, want %t", test.desc, got, want)
		}
	}
}

func TestAdd(t *testing.T) {
	tests := []struct {
		desc  string
		tree  *Tree
		value int
		want  *Tree
		added bool
	}{
		{
			"empty",
			&Tree{},
			5,
			&Tree{root: &Node{value: 5}},
			true,
		},
		{
			"duplicate root",
			&Tree{root: &Node{value: 5}},
			5,
			&Tree{root: &Node{value: 5}},
			false,
		},
		{
			"less than root",
			&Tree{root: &Node{value: 5}},
			3,
			&Tree{root: &Node{value: 5, left: &Node{value: 3}}},
			true,
		},
		{
			"greater than root",
			&Tree{root: &Node{value: 5}},
			7,
			&Tree{root: &Node{value: 5, right: &Node{value: 7}}},
			true,
		},
		{
			"add child left, left",
			&Tree{root: &Node{value: 5, left: &Node{value: 2}, right: &Node{value: 7}}},
			1,
			&Tree{root: &Node{value: 5, left: &Node{value: 2, left: &Node{value: 1}}, right: &Node{value: 7}}},
			true,
		},
		{
			"add child left, right",
			&Tree{root: &Node{value: 5, left: &Node{value: 2}, right: &Node{value: 7}}},
			3,
			&Tree{root: &Node{value: 5, left: &Node{value: 2, right: &Node{value: 3}}, right: &Node{value: 7}}},
			true,
		},
		{
			"add child right, left",
			&Tree{root: &Node{value: 5, left: &Node{value: 2}, right: &Node{value: 7}}},
			6,
			&Tree{root: &Node{value: 5, left: &Node{value: 2}, right: &Node{value: 7, left: &Node{value: 6}}}},
			true,
		},
		{
			"add child right, right",
			&Tree{root: &Node{value: 5, left: &Node{value: 2}, right: &Node{value: 7}}},
			8,
			&Tree{root: &Node{value: 5, left: &Node{value: 2}, right: &Node{value: 7, right: &Node{value: 8}}}},
			true,
		},
		{
			"nested duplicate",
			&Tree{root: &Node{value: 5, left: &Node{value: 2}, right: &Node{value: 7}}},
			7,
			&Tree{root: &Node{value: 5, left: &Node{value: 2}, right: &Node{value: 7}}},
			false,
		},
	}

	for _, test := range tests {
		ga := test.tree.Add(test.value)
		wa := test.added

		if ga != wa {
			t.Errorf("%v: got added %t, want added %t", test.desc, ga, wa)
			continue
		}

		gt := test.tree
		wt := test.want
		if !gt.Equal(wt) {
			t.Errorf("%v:\ngot\t%s\nwant\t%s", test.desc, gt, wt)
		}
	}
}

func TestRemove(t *testing.T) {
	tests := []struct {
		desc    string
		tree    *Tree
		value   int
		want    *Tree
		removed bool
	}{
		{
			"empty",
			&Tree{},
			5,
			&Tree{},
			false,
		},
		{
			"root",
			&Tree{root: &Node{value: 5}},
			5,
			&Tree{},
			true,
		},
		{
			"root not found",
			&Tree{root: &Node{value: 5}},
			3,
			&Tree{root: &Node{value: 5}},
			false,
		},
		{
			"root node, left",
			&Tree{root: &Node{value: 5, left: &Node{value: 2}}},
			5,
			&Tree{root: &Node{value: 2}},
			true,
		},
		{
			"root node, right",
			&Tree{root: &Node{value: 5, right: &Node{value: 8}}},
			5,
			&Tree{root: &Node{value: 8}},
			true,
		},
		{
			"leaf left, left",
			&Tree{root: &Node{value: 5, left: &Node{value: 2, left: &Node{value: 1}}, right: &Node{value: 7}}},
			1,
			&Tree{root: &Node{value: 5, left: &Node{value: 2}, right: &Node{value: 7}}},
			true,
		},
		{
			"leaf left, right",
			&Tree{root: &Node{value: 5, left: &Node{value: 2, right: &Node{value: 3}}, right: &Node{value: 7}}},
			3,
			&Tree{root: &Node{value: 5, left: &Node{value: 2}, right: &Node{value: 7}}},
			true,
		},
		{
			"leaf right, left",
			&Tree{root: &Node{value: 5, left: &Node{value: 2}, right: &Node{value: 7, left: &Node{value: 6}}}},
			6,
			&Tree{root: &Node{value: 5, left: &Node{value: 2}, right: &Node{value: 7}}},
			true,
		},
		{
			"leaf right, right",
			&Tree{root: &Node{value: 5, left: &Node{value: 2}, right: &Node{value: 7, right: &Node{value: 8}}}},
			8,
			&Tree{root: &Node{value: 5, left: &Node{value: 2}, right: &Node{value: 7}}},
			true,
		},
		{
			"one child, left",
			&Tree{root: &Node{value: 5, left: &Node{value: 2}, right: &Node{value: 7, left: &Node{value: 6}}}},
			7,
			&Tree{root: &Node{value: 5, left: &Node{value: 2}, right: &Node{value: 6}}},
			true,
		},
		{
			"one child, right",
			&Tree{root: &Node{value: 5, left: &Node{value: 2}, right: &Node{value: 7, right: &Node{value: 8}}}},
			7,
			&Tree{root: &Node{value: 5, left: &Node{value: 2}, right: &Node{value: 8}}},
			true,
		},
		{
			"two children, left",
			&Tree{root: &Node{value: 5, left: &Node{value: 2}, right: &Node{value: 7, left: &Node{value: 6}, right: &Node{value: 8}}}},
			7,
			&Tree{root: &Node{value: 5, left: &Node{value: 2}, right: &Node{value: 8, left: &Node{value: 6}}}},
			true,
		},
		{
			"two children, right",
			&Tree{root: &Node{value: 8, left: &Node{value: 4, left: &Node{value: 1}, right: &Node{value: 7}}, right: &Node{value: 12}}},
			4,
			&Tree{root: &Node{value: 8, left: &Node{value: 7, left: &Node{value: 1}}, right: &Node{value: 12}}},
			true,
		},
		{
			"nested not found",
			&Tree{root: &Node{value: 8, left: &Node{value: 4, left: &Node{value: 1}, right: &Node{value: 7}}, right: &Node{value: 7, left: &Node{value: 6}, right: &Node{value: 8}}}},
			9,
			&Tree{root: &Node{value: 8, left: &Node{value: 4, left: &Node{value: 1}, right: &Node{value: 7}}, right: &Node{value: 7, left: &Node{value: 6}, right: &Node{value: 8}}}},
			false,
		},
	}

	for _, test := range tests {
		gr := test.tree.Remove(test.value)
		wr := test.removed

		if gr != wr {
			t.Errorf("%v: got removed %t, want removed %t", test.desc, gr, wr)
			continue
		}

		gt := test.tree
		wt := test.want
		if !gt.Equal(wt) {
			t.Errorf("%v:\ngot\t%s\nwant\t%s", test.desc, gt, wt)
		}
	}
}
