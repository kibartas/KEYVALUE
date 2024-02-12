package main


type stableStore struct {}


func (s stableStore) Set(key []byte, val []byte) error {
	return nil
}

func (s stableStore) Get(key []byte) ([]byte, error) {
	return []byte{}, nil
}

func (s stableStore) SetUint64(key []byte, val uint64) error {
	return nil
}

func (s stableStore) GetUint64(key []byte) (uint64, error) {
	return 0, nil
}
