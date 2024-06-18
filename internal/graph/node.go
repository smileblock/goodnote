package graph

type NodeId string

type NodeName string

type Node struct {
	id   NodeId
	name NodeName
}

func NewNode(id NodeId, name NodeName) (*Node, error) {
	return &Node{id: id, name: name}, nil
}
