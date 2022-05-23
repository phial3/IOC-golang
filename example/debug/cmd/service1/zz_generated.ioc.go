//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by iocli

package service1

import (
	"github.com/alibaba/ioc-golang/autowire"
	"github.com/alibaba/ioc-golang/autowire/singleton"
)

func init() {
	singleton.RegisterStructDescriptor(&autowire.StructDescriptor{
		Interface: new(Service1),
		Factory: func() interface{} {
			return &Impl1{}
		},
	})
}
