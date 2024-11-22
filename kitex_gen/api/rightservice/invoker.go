// Code generated by Kitex v0.10.1. DO NOT EDIT.

package rightservice

import (
	server "github.com/cloudwego/kitex/server"
	api "github.com/zlllgp/vegas/kitex_gen/api"
)

// NewInvoker creates a server.Invoker with the given handler and options.
func NewInvoker(handler api.RightService, opts ...server.Option) server.Invoker {
	var options []server.Option

	options = append(options, opts...)

	s := server.NewInvoker(options...)
	if err := s.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	if err := s.Init(); err != nil {
		panic(err)
	}
	return s
}
