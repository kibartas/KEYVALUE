package main

import (
	"github.com/hashicorp/raft"
)

type logStore struct {
	logStoreMap []raft.Log
}

func (l *logStore) FirstIndex() (uint64, error) {
	logger.Println("FirstIndex called")
	if len(l.logStoreMap) < 1 {
		return 0, nil
	}
	return 1, nil
}

func (l *logStore) LastIndex() (uint64, error) {
	logger.Println("LastIndex called")
	if len(l.logStoreMap) < 1 {
		return 0, nil
	}
	return uint64(len(l.logStoreMap)), nil
}

func (l *logStore) GetLog(index uint64, log *raft.Log) error {
	logger.Println("GetLog called with index", index)
	if len(l.logStoreMap) < int(index) {
		return nil
	}
	logger.Println("GetLog get here with value ", &l.logStoreMap[index - 1])
	log = &l.logStoreMap[index - 1]
	return nil
}

func (l *logStore) StoreLog(log *raft.Log) error {
	logger.Println("StoreLog called with " + string(log.Data))
	l.logStoreMap = append(l.logStoreMap, *log)
	logger.Println(l.logStoreMap)
	return nil
}

func (l *logStore) StoreLogs(logs []*raft.Log) error {
	logger.Println("StoreLogs called")
	for _, v := range logs {
		err := l.StoreLog(v)
		if err != nil {
			return err
		}
	}
	return nil
}

func (l *logStore) DeleteRange(min, max uint64) error {
	logger.Println("DeleteRange called")
	// TODO: add func
	return nil
}
