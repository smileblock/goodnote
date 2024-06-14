package graph

import "goodnode/internal/common"

type EdgeName = string

type Edge struct {
	id         common.Id
	name       EdgeName
	sourceNode *Node
	targetNode *Node
}
