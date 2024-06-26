package main

import (
	"context"
	"github.com/zlllgp/vegas/kitex_gen/api"
)

type VegasImpl struct{}

// Draw implements the VegasImpl interface.
func (s *VegasImpl) Draw(ctx context.Context, req *api.DrawRequest) (resp *api.DrawResult, err error) {
	// TODO: Your code here...
	resp = &api.DrawResult{Code: "SUCCESS", Msg: "", Rights: &api.RightsDTO{Id: 1, Num: 1, Amt: "10"}}
	return
}
