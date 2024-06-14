package graph

import (
	"errors"
)

type NodeName = string

const (
	NodeKindPaper = "paper"
	NodeKindBoard = "board"
	NodeKindFile  = "file"
)

type NodeKind = string

type nodeMeta struct {
	name NodeName
	kind NodeKind
}

type Node struct {
	meta    *nodeMeta
	content any
}

func NewNode(nodeType NodeKind) (*Node, error) {
	if nodeType != NodeKindPaper && nodeType != NodeKindBoard && nodeType != NodeKindFile {
		return nil, errors.New("invalid value: must be 'paper', 'board', or 'file'")
	}
	meta := nodeMeta{kind: nodeType}
	return &Node{meta: &meta}, nil
}
