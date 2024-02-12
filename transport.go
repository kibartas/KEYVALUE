package main

import (
	"github.com/hashicorp/raft"
)

type transport struct {}

func (t *transport) Consumer() <-chan raft.RPC {
	return make(<-chan raft.RPC)
}

func (t transport) LocalAddr() raft.ServerAddress {
	return "127.0.0.1"
}

// func AppendEntriesPipeline(id raft.ServerID, target raft.ServerAddress) (raft.AppendPipeline, error) {
// 	return interface{}{}, nil
// }

// // AppendEntries sends the appropriate RPC to the target node.
// func AppendEntries(id raft.ServerID, target raft.ServerAddress, args *raft.AppendEntriesRequest, resp *raft.AppendEntriesResponse) error

// // RequestVote sends the appropriate RPC to the target node.
// func RequestVote(id raft.ServerID, target raft.ServerAddress, args *raft.RequestVoteRequest, resp *raft.RequestVoteResponse) error

// // InstallSnapshot is used to push a snapshot down to a follower. The data is read from
// // the ReadCloser and streamed to the client.
// func InstallSnapshot(id raft.ServerID, target raft.ServerAddress, args *raft.InstallSnapshotRequest, resp *raft.InstallSnapshotResponse, data io.Reader) error

// // EncodePeer is used to serialize a peer's address.
// func EncodePeer(id raft.ServerID, addr raft.ServerAddress) []byte

// // DecodePeer is used to deserialize a peer's address.
// func DecodePeer([]byte) raft.ServerAddress

// // SetHeartbeatHandler is used to setup a heartbeat handler
// // as a fast-pass. This is to avoid head-of-line blocking from
// // disk IO. If a Transport does not support this, it can simply
// // ignore the call, and push the heartbeat onto the Consumer channel.
// func SetHeartbeatHandler(cb func(rpc raft.RPC))

// // TimeoutNow is used to start a leadership transfer to the target node.
// func TimeoutNow(id raft.ServerID, target raft.ServerAddress, args *raft.TimeoutNowRequest, resp *raft.TimeoutNowResponse) error