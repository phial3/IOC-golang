//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by ioc-go-cli

package redis

import (
	autowire "github.com/alibaba/ioc-golang/autowire"
	"github.com/alibaba/ioc-golang/autowire/normal"
)

func init() {
	normal.RegisterStructDescriptor(&autowire.StructDescriptor{
		Interface: new(Redis),
		Factory: func() interface{} {
			return &Impl{}
		},
		ParamFactory: func() interface{} {
			var _ configInterface = &Config{}
			return &Config{}
		},
		ConstructFunc: func(i interface{}, p interface{}) (interface{}, error) {
			param := p.(configInterface)
			impl := i.(*Impl)
			return param.New(impl)
		},
	})
}

type configInterface interface {
	New(impl *Impl) (*Impl, error)
}