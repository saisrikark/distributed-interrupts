package node

import (
	"bufio"
	"context"
	"fmt"
	"log/slog"
	"net"
	"os"
)

type node struct {
	id     int
	logger *slog.Logger
}

type NodeConfig struct {
	ID int
}

func NewNode(nc NodeConfig) *node {
	return &node{
		id:     nc.ID,
		logger: slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{})).With("node", nc.ID),
	}
}

func (n *node) Run(ctx context.Context) error {
	n.logger.InfoContext(ctx, "running node")
	go n.runListner(ctx)

	defer func() {
		n.logger.InfoContext(ctx, "shutting down")
	}()

	<-ctx.Done()
	return nil
}

func (n *node) runListner(ctx context.Context) error {
	ln, err := net.Listen("tcp", ":0")
	if err != nil {
		return fmt.Errorf("unable to create listner %v", err)
	}

	n.logger.InfoContext(ctx, "started listener", "addr", ln.Addr(), "err", err)

	for {
		conn, err := ln.Accept()
		if err != nil {
			return fmt.Errorf("unable to accept from listner")
		}

		go n.handleConnection(ctx, conn)
	}
}

func (n *node) handleConnection(ctx context.Context, conn net.Conn) error {
	defer conn.Close()

	for {
		line, _, err := bufio.NewReader(conn).ReadLine()
		if err != nil {
			n.logger.WarnContext(ctx, "unable to read line", "err", err)
			break
		}

		n.logger.InfoContext(ctx, "received message", "msg", string(line))
	}

	return nil
}
