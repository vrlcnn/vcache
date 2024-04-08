package vcache

import "sync"

type Shard struct {
	hashmap map[uint64]uint64
	entries byteQueue
	lock    sync.RWMutex
}
