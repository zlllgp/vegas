package right

import (
	"context"
	"github.com/zlllgp/vegas/kitex_gen/api"
)

type RightServiceImpl struct {
	ctx context.Context
}

func NewRightServiceImpl(ctx context.Context) *RightServiceImpl {
	return &RightServiceImpl{ctx: ctx}
}

// QueryRight implements the RightServiceImpl interface.
func (s *RightServiceImpl) QueryRight(ctx context.Context, req *api.RightRequest) (resp *api.RightResponse, err error) {

	return
}
