package main

import (
	"fmt"
	"log"
	"net"
	"time"

	"github.com/google/uuid"
	"github.com/hashicorp/raft"
)

func GetFreePort() (port string, err error) {
	var a *net.TCPAddr
	if a, err = net.ResolveTCPAddr("tcp", "localhost:0"); err == nil {
		var l *net.TCPListener
		if l, err = net.ListenTCP("tcp", a); err == nil {
			defer l.Close()
			return fmt.Sprint(l.Addr().(*net.TCPAddr).Port), nil
		}
	}
	return
}

func CreateNewNode(logger *log.Logger) (raftNode *raft.Raft, serverAddress raft.ServerAddress, serverID raft.ServerID, err error) {
	config := raft.DefaultConfig()
    serverID = raft.ServerID(uuid.New().String())
	config.LocalID = serverID

    serverPort, err := GetFreePort()

    if err != nil {
        return nil, "", "", err
    }

    serverAddr := "127.0.0.1" + ":" + serverPort

	listener, err := net.Listen("tcp", serverAddr)

	if err != nil {
		return nil, "", "", err
	}

	var streamLayer raft.StreamLayer = StreamLayer{
		Listener: listener,
	}
    

	logStore := &logStore{
		logStoreMap: []raft.Log{},
	}


	raftNode, err = raft.NewRaft(config, &raft.MockFSM{}, logStore, NewStableStore(), &snapshotStore{},
		raft.NewNetworkTransport(streamLayer, 5, time.Duration(time.Second), logger.Writer()))

	return raftNode, raft.ServerAddress(serverAddr), serverID, err
}