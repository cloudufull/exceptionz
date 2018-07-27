package impl

import (
	"crypto/md5"
	"fmt"
	"github.com/andygeiss/hashing/api"
	"golang.org/x/net/context"
)

// DefaultHashing ...
type DefaultHashing struct{}

// NewDefaultHashing ...
func NewDefaultHashing() api.HashingServer {
	return &DefaultHashing{}
}

func (h *DefaultHashing) Hash(ctx context.Context, req *api.HashRequest) (*api.HashResponse, error) {
	return &api.HashResponse{
		Hash: hashWithMD5(req.Plain),
	}, nil
}

func hashWithMD5(data []byte) uint64 {
	hex := fmt.Sprintf("%x", md5.Sum(data))
	var hash uint64
	for i, dig := range hex {
		hash += uint64(dig) << uint(i)
	}
	return hash
}

func hashWithCustomShiftXor(data []byte) uint64 {
	var hash uint64
	for i, dig := range data {
		hash += uint64(dig) << uint(i)
	}
	return hash
}
