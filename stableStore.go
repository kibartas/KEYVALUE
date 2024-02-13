package main

import "fmt"

type stableStore struct{
	byteMap map[string][]byte
	uint64Map map[string]uint64
}

func NewStableStore() *stableStore {
	return &stableStore{
		byteMap: map[string][]byte{},
		uint64Map: map[string]uint64{},
	}
}

func (s *stableStore) Set(key []byte, val []byte) error {
	fmt.Println("Set was called", key, val)
	strKey := fmt.Sprint(key)
	s.byteMap[strKey] = val
	return nil
}

func (s *stableStore) Get(key []byte) ([]byte, error) {
	fmt.Println("Get was called", key)
	strKey := fmt.Sprint(key)
	if val, ok := s.byteMap[strKey]; ok {
		return val, nil
	} 
	return []byte{}, nil
}

func (s *stableStore) SetUint64(key []byte, val uint64) error {
	fmt.Println("SetUint64 was called", string(key[:]), val)
	strKey := fmt.Sprint(key)
	s.uint64Map[strKey] = val
	return nil
}

func (s *stableStore) GetUint64(key []byte) (uint64, error) {
	fmt.Println("GetUint64 was called", string(key[:]))
	strKey := fmt.Sprint(key)
	if val, ok := s.uint64Map[strKey]; ok {
		return val, nil
	}
	return 0, nil
}
