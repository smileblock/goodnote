package graph

import (
	"testing"
)

func TestNewNode(t *testing.T) {
	node, err := NewNode("1", "node")
	if err != nil {
		t.Error(err)
	}
	t.Log(node)
}
