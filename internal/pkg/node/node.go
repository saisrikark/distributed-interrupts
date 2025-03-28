package node

import (
	"context"
	"log/slog"
)

type node struct {
	id int
}

type NodeConfig struct {
	ID int
}

func NewNode(nc NodeConfig) *node {
	return &node{
		id: nc.ID,
	}
}

func (n *node) Run(ctx context.Context) error {
	slog.InfoContext(ctx, "running node", "id", n.id)
	defer func() {
		slog.InfoContext(ctx, "shutting down node", "id", n.id)
	}()
	<-ctx.Done()
	return nil
}
