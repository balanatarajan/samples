package graph

type GraphNode struct {
	Val int
	adj []*GraphNode
}

func (n *GraphNode) Adjacent() []*GraphNode {
	return n.adj
}

func (n *GraphNode) AddNode(nd *GraphNode) {
	n.adj = append(n.adj, nd)
}

func (n *GraphNode) AddNodes(nds []*GraphNode) {
	n.adj = append(n.adj, nds...)
}

func New(v int) *GraphNode {

	return &GraphNode{Val: v}
}

func (n *GraphNode) Add(v int) {
	nd := &GraphNode{Val: v}
	n.AddNode(nd)
}

func (n *GraphNode) Find(v int) *GraphNode {
	for _, val := range n.adj {

		if val.Val == v {
			return val
		}
	}

	return nil
}
