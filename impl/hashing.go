package impl

import (
	"crypto/md5"
	"fmt"
	"github.com/andygeiss/exceptionz/api"
	"golang.org/x/net/context"
)

// DefaultHashing uses a MD5 impementation to create a hash for a given plaintext.
type DefaultHashing struct{}

// NewDefaultHashing creates a new HashingServer implementation and returns its address.
func NewDefaultHashing() api.HashingServer {
	return &DefaultHashing{}
}

// Hash creates an unsigned 64-bit integer by using MD5.
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
