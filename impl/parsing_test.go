package impl_test

import (
	. "github.com/andygeiss/assert"
	"github.com/andygeiss/exceptionz/api"
	"github.com/andygeiss/exceptionz/impl"
	"golang.org/x/net/context"
	"testing"
)

func BenchmarkDefaultParsing_ParseExceptions(b *testing.B) {
	hashing := impl.NewDefaultHashing()
	service := impl.NewDefaultParsing("../testdata/srv-one.log", hashing)
	request := &api.ParseExceptionsRequest{}
	for i := 0; i < b.N; i++ {
		service.ParseExceptions(context.Background(), request)
	}
}

func TestDefaultParsing_ParseExceptions_Should_Return_Without_Error(t *testing.T) {
	hashing := impl.NewDefaultHashing()
	service := impl.NewDefaultParsing("../testdata/srv-one.log", hashing)
	_, err := service.ParseExceptions(context.Background(), &api.ParseExceptionsRequest{})
	Assert(t, err, IsNil())
}
