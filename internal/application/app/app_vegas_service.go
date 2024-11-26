package app

import (
	"context"
	"github.com/zlllgp/vegas/internal/domain/service"
	"github.com/zlllgp/vegas/kitex_gen/api"
	"github.com/zlllgp/vegas/pkg/errno"
)

type VegasServiceImpl struct {
	drawService service.DrawService
}

func NewVegasServiceImpl(drawService service.DrawService) *VegasServiceImpl {
	return &VegasServiceImpl{drawService: drawService}
}

// Draw implements the VegasServiceImpl interface.
func (s *VegasServiceImpl) Draw(ctx context.Context, req *api.DrawRequest) (resp *api.DrawResponse, err error) {
	//biz logic in domain service
	drawResult, err := s.drawService.Draw(ctx, req.User.UserId, req.Eh)
	if err != nil {
		resp = &api.DrawResponse{
			Rights: nil,
			Base: &api.BaseResponse{
				Code: errno.DrawErr.GetCode(),
				Msg:  errno.DrawErr.GetMsg()},
		}
		return resp, nil
	}
	return &api.DrawResponse{
		Rights: drawResult.Rights,
		Base: &api.BaseResponse{
			Code: errno.Success.GetCode(),
			Msg:  errno.Success.GetMsg()},
	}, err
}

// Show implements the VegasServiceImpl interface.
func (s *VegasServiceImpl) Show(ctx context.Context, req *api.ShowRequest) (resp *api.ShowResponse, err error) {
	// TODO: Your code here...
	return
}
