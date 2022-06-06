//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by iocli

package service

import (
	autowire "github.com/alibaba/ioc-golang/autowire"
	rpc_service "github.com/alibaba/ioc-golang/extension/autowire/rpc/rpc_service"
)

func init() {
	rpc_service.RegisterStructDescriptor(&autowire.StructDescriptor{
		Alias: "github.com/alibaba/ioc-golang/example/autowire_rpc/client/test/service/api.ComplexServiceIOCRPCClient",
		Factory: func() interface{} {
			return &ComplexService{}
		},
	})
	rpc_service.RegisterStructDescriptor(&autowire.StructDescriptor{
		Alias: "github.com/alibaba/ioc-golang/example/autowire_rpc/client/test/service/api.SimpleServiceIOCRPCClient",
		Factory: func() interface{} {
			return &SimpleService{}
		},
	})
}