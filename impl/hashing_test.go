package impl_test

import (
	. "github.com/andygeiss/assert"
	"github.com/andygeiss/exceptionz/api"
	"github.com/andygeiss/exceptionz/impl"
	"golang.org/x/net/context"
	"testing"
)

func BenchmarkDefaultHashing_Hash(b *testing.B) {
	service := impl.NewDefaultHashing()
	request := &api.HashRequest{
		Plain: []byte(`de.genoip.vrs.configuration.VRSSNExtensionConfigurationFacade java.io.IOException: resource "lib/ini/system/vrs/vrs_mqconfiguration.properties" not found`),
	}
	for i := 0; i < b.N; i++ {
		service.Hash(context.Background(), request)
	}
}

func TestDefaultHashing_Hash_Should_Return_No_Error(t *testing.T) {
	service := impl.NewDefaultHashing()
	request := &api.HashRequest{
		Plain: []byte(`de.genoip.vrs.configuration.VRSSNExtensionConfigurationFacade java.io.IOException: resource "lib/ini/system/vrs/vrs_mqconfiguration.properties" not found`),
	}
	res, err := service.Hash(context.Background(), request)
	Assert(t, err, IsNil())
	Assert(t, res.Hash, IsNotEqual(uint64(0)))
}
