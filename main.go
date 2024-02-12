package main

import (
	"fmt"

	"github.com/hashicorp/raft"
)


func main() {
    config := raft.DefaultConfig()
    config.LocalID = raft.ServerID("123")

    logStore := logStore{
        logStoreMap: []raft.Log{},
    }

    err := raft.BootstrapCluster(config, logStore, stableStore{}, snapshotStore{}, &raft.NetworkTransport{}, raft.Configuration{
        Servers: []raft.Server{
            {
                Suffrage: raft.Voter,
                ID: "123",
                Address: "127.0.0.1",
            },
        },
    })
    fmt.Println(err)

    raft, err := raft.NewRaft(config, &raft.MockFSM{}, logStore, stableStore{}, snapshotStore{}, &raft.NetworkTransport{})

    fmt.Println(raft, err)
}