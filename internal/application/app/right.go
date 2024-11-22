package app

import (
	"context"
	"github.com/zlllgp/vegas/kitex_gen/api"
)

type RightServiceImpl struct {
}

func NewRightServiceImpl() *RightServiceImpl {
	return &RightServiceImpl{}
}

// QueryRight implements the RightServiceImpl interface.
func (s *RightServiceImpl) QueryRight(ctx context.Context, req *api.RightRequest) (resp *api.RightResponse, err error) {

	return
}
