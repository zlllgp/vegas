package app

import (
	"context"
	"github.com/zlllgp/vegas/internal/application/inf"
	"github.com/zlllgp/vegas/kitex_gen/api"
)

type VegasServiceImpl struct {
	service inf.DrawService
}

func NewVegasServiceImpl(service inf.DrawService) *VegasServiceImpl {
	return &VegasServiceImpl{service: service}
}

// Draw implements the VegasServiceImpl interface.
func (s *VegasServiceImpl) Draw(ctx context.Context, req *api.DrawRequest) (resp *api.DrawResponse, err error) {
	drawResult, err := s.service.Draw(ctx, req.User.UserId, req.Eh)
	if err != nil {
		resp = &api.DrawResponse{
			Rights: drawResult.Rights,
			Code:   "SUCCESS",
			Msg:    "WIN",
		}
		return resp, nil
	}
	return nil, err
}

// Show implements the VegasServiceImpl interface.
func (s *VegasServiceImpl) Show(ctx context.Context, req *api.ShowRequest) (resp *api.ShowResponse, err error) {
	// TODO: Your code here...
	return
}
