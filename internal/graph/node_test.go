package graph

import (
	"testing"
)

func TestNewNode(t *testing.T) {
	_, err := NewNode(NodeKindFile)
	if err != nil {
		print(err)
	}
}
