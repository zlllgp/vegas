package main

import (
	"context"
	"github.com/zlllgp/vegas/biz/service"
	"github.com/zlllgp/vegas/kitex_gen/api"
)

type VegasImpl struct{}

// Draw implements the VegasImpl interface.
func (s *VegasImpl) Draw(ctx context.Context, req *api.DrawRequest) (resp *api.DrawResult, err error) {
	resp, err = service.NewVegasService(ctx).Draw(req)
	return resp, err
}
