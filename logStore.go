package main

import (
	"errors"

	"github.com/hashicorp/raft"
)

type logStore struct {
	logStoreMap []raft.Log
}

func (l logStore) FirstIndex() (uint64, error) {
	if len(l.logStoreMap) < 1 {
		return 0, nil
	}
	return 0, nil
}

func (l logStore) LastIndex() (uint64, error) {
	if len(l.logStoreMap) < 1 {
		return 0, nil
	}
	return uint64(len(l.logStoreMap) - 1), nil
}

func (l logStore) GetLog(index uint64, log *raft.Log) error {
	if len(l.logStoreMap) < int(index) {
		return errors.New("Index out of range")
	}
	log = &l.logStoreMap[index]
	return nil
}

func (l *logStore) StoreLog(log *raft.Log) error {
	l.logStoreMap = append(l.logStoreMap, *log)
	return nil
}

func (l *logStore) StoreLogs(logs []*raft.Log) error {
	for _, v := range logs {
		err := l.StoreLog(v)
		if err != nil {
			return err
		}
	}
	return nil
}

func (l *logStore) DeleteRange(min, max uint64) error {
	// TODO: add func
	return nil
}
