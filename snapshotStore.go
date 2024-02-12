package main

import (
	"bufio"
	"io"

	"github.com/hashicorp/raft"
)

type snapshotStore struct {
	*snapshotSink
}

type snapshotSink struct {
	*bufio.Writer
}

type snapshotReader struct {
	*bufio.Reader
}

func (s snapshotReader) Close() error {
	return nil
}

func (s snapshotSink) Close() error {
	return nil
}

func (s snapshotSink) ID() string {
	return "lol"
}

func (s snapshotSink) Cancel() error {
	return nil
}

func (s snapshotStore) Create(version raft.SnapshotVersion, index, term uint64, configuration raft.Configuration,
		configurationIndex uint64, trans raft.Transport) (raft.SnapshotSink, error) {
			return s.snapshotSink, nil
		}

func (s snapshotStore) List() ([]*raft.SnapshotMeta, error) {
	return []*raft.SnapshotMeta{}, nil
}

func (s snapshotStore) Open(id string) (*raft.SnapshotMeta, io.ReadCloser, error) {
	return &raft.SnapshotMeta{}, snapshotReader{}, nil
}