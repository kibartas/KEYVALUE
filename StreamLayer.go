package main

import (
	"context"
	"net"
	"time"

	"github.com/hashicorp/raft"
)

type StreamLayer struct {
	net.Listener
}

func (s StreamLayer) Dial(address raft.ServerAddress, timeout time.Duration) (net.Conn, error) {
	var d net.Dialer
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	return d.DialContext(ctx, "tcp", string(address))
}
