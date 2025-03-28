package main

import (
	"context"
	"flag"
	"os"
	"os/signal"

	"github.com/saisrikark/distributed-interrupts/internal/pkg/node"
)

func main() {
	number := flag.Int("number", 3, "define the number of nodes you want to run")

	if *number <= 0 {
		panic("you need to define at least one node")
	}

	ctx, cancel := context.WithCancel(context.Background())

	nodeCompletion := make([]chan struct{}, *number)

	for i := range *number {
		newnode := node.NewNode(node.NodeConfig{
			ID: i,
		})
		nodeCompletion[i] = make(chan struct{})

		go func() {
			defer func() {
				nodeCompletion[i] <- struct{}{}
			}()
			newnode.Run(ctx)
		}()
	}

	notifyChan := make(chan os.Signal, 1)
	signal.Notify(notifyChan, os.Interrupt)

	<-notifyChan
	cancel()

	for i := range *number {
		<-nodeCompletion[i]
	}
}
