package node

import "context"

type Node interface {
	Run(ctx context.Context) error
}
