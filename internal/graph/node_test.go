package graph

import (
	"goodnode/pkg/logger"
	"testing"
)

func TestNewNode(t *testing.T) {
	node, err := NewNode("AA")
	if err != nil {
		logger.Err(err)
	}
	logger.Debug(node)
}
