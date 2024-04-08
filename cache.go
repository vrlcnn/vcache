package vcache

import (
	"context"
)

type VCache struct {
	shards []*Shard
}

func New(ctx context.Context) (*VCache, error) {
	return newVCache(ctx)
}

func NewVCache() (*VCache, error) {
	return newVCache(context.Background())
}

func newVCache(ctx context.Context) (*VCache, error) {

	vc := &VCache{
		shards: make([]*Shard, 0),
	}

	return vc, nil
}
