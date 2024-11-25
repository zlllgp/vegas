package app

import (
	"context"
	"github.com/zlllgp/vegas/internal/application/inf"
	"github.com/zlllgp/vegas/internal/infrastructure/errors"
	"github.com/zlllgp/vegas/kitex_gen/api"
)

type VegasServiceImpl struct {
	drawService inf.DrawService
}

func NewVegasServiceImpl(drawService inf.DrawService) *VegasServiceImpl {
	return &VegasServiceImpl{drawService: drawService}
}

// Draw implements the VegasServiceImpl interface.
func (s *VegasServiceImpl) Draw(ctx context.Context, req *api.DrawRequest) (resp *api.DrawResponse, err error) {
	//biz logic in domain service , but use inf
	drawResult, err := s.drawService.Draw(ctx, req.User.UserId, req.Eh)
	if err != nil {
		resp = &api.DrawResponse{
			Rights: nil,
			Base: &api.BaseResponse{
				Code: errors.DrawErr.GetCode(),
				Msg:  errors.DrawErr.GetMsg()},
		}
		return resp, nil
	}
	return &api.DrawResponse{
		Rights: drawResult.Rights,
		Base: &api.BaseResponse{
			Code: errors.Success.GetCode(),
			Msg:  errors.Success.GetMsg()},
	}, err
}

// Show implements the VegasServiceImpl interface.
func (s *VegasServiceImpl) Show(ctx context.Context, req *api.ShowRequest) (resp *api.ShowResponse, err error) {
	// TODO: Your code here...
	return
}
