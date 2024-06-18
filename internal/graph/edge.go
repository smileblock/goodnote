package graph

type Edge struct {
	from *Node
	to   *Node
}

func NewEdge(from *Node, to *Node) Edge {
	return Edge{from: from, to: to}
}
