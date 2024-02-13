package main

import (
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/raft"
)

var logger = log.Default()

func main() {
    leaderNode, serverAddress, serverID, err := CreateNewNode(logger)

    if err != nil {
        panic(err)
    }

    configuration := raft.Configuration{
		Servers: []raft.Server{
			{
				Suffrage: raft.Voter,
				ID:       serverID,
				Address:  serverAddress,
			},
		},
	}

    voterNode, voterNodeAddr, voterNodeID, err := CreateNewNode(logger)

    if err != nil {
        panic(err)
    }

    future := leaderNode.BootstrapCluster(configuration)

    voterFuture := leaderNode.AddVoter(voterNodeID, voterNodeAddr, 0, time.Duration(time.Second))


	for {
		time.Sleep(time.Second)
        fmt.Println(leaderNode.Stats())
        fmt.Println(voterNode.Stats())
        fmt.Println(future.Error())
        fmt.Println(voterFuture.Error())
	}
}
