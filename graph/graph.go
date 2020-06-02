package graph

// Node represents a graph node with its edge references.
type Node struct {
	value int
	edges []*Node
}

// Graph represents a directed or undirected graph.
type Graph struct {
	directed bool
	nodes    []*Node
}

// Add a node to the graph.
func (g *Graph) Add(n *Node) {
	if n == nil {
		return
	}
	g.nodes = append(g.nodes, n)
}

func removeIndex(n []*Node, i int) []*Node {
	copy(n[i:], n[i+1:]) // Shift n[i+1:] left one index.
	n[len(n)-1] = nil    // Erase last element (write zero value).
	n = n[:len(n)-1]     // Truncate slice.
	return n
}

// Remove nodes and edges with the provided value.
func (g *Graph) Remove(v int) {
	var nodesToRemove []int
	for in, n := range g.nodes {
		if n.value == v {
			// mark node for removal
			nodesToRemove = append(nodesToRemove, in)
			continue
		}

		var edgesToRemove []int
		for ie, e := range n.edges {
			if e.value == v {
				edgesToRemove = append(edgesToRemove, ie)
			}
		}
		for _, i := range edgesToRemove {
			n.edges = removeIndex(n.edges, i)
		}
	}

	for _, i := range nodesToRemove {
		g.nodes = removeIndex(g.nodes, i)
	}
}

// Node gets a node with the provided value.
func (g *Graph) Node(v int) (*Node, bool) {
	for _, n := range g.nodes {
		if n.value == v {
			return n, true
		}
	}
	return nil, false
}

// AddEdge adds an edge link between two nodes based on their value.
func (g *Graph) AddEdge(v1, v2 int) {
	n1, ok1 := g.Node(v1)
	n2, ok2 := g.Node(v2)

	if !ok1 || !ok2 {
		return
	}

	n1.edges = append(n1.edges, n2)

	if !g.directed {
		n2.edges = append(n2.edges, n1)
	}
}
