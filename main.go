package main

import (
	"fmt"
	"net"
	"time"

	"github.com/hashicorp/raft"
)

func main() {
	config := raft.DefaultConfig()
	config.LocalID = raft.ServerID("123")

	logStore := &logStore{
		logStoreMap: []raft.Log{},
	}

	err := raft.BootstrapCluster(config, logStore, &stableStore{}, &snapshotStore{}, &raft.NetworkTransport{}, raft.Configuration{
		Servers: []raft.Server{
			{
				Suffrage: raft.Voter,
				ID:       "123",
				Address:  "127.0.0.1",
			},
		},
	})
	fmt.Println(err)

	listener, err := net.Listen("tcp", "127.0.0.1:3333")

	if err != nil {
		panic(err)
	}

	fmt.Println(listener.Addr())

	var streamLayer raft.StreamLayer = StreamLayer{
		Listener: listener,
	}

	raft, err := raft.NewRaft(config, &raft.MockFSM{}, logStore, &stableStore{}, &snapshotStore{},
		raft.NewNetworkTransport(streamLayer, 1, time.Duration(10), nil))
	fmt.Println(raft, err)
	for {
		time.Sleep(time.Second)
		fmt.Println(raft.Stats())
	}
	// raft.BootstrapCluster()
}
