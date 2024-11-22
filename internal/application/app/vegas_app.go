package app

import (
	"context"
	"github.com/zlllgp/vegas/internal/domain/service"
	"github.com/zlllgp/vegas/kitex_gen/api"
)

type VegasApplication struct {
	service *service.DrawService
}

func NewVegasServiceImpl(service *service.DrawService) *VegasApplication {
	return &VegasApplication{service: service}
}

// Draw implements the VegasServiceImpl interface.
func (s *VegasApplication) Draw(ctx context.Context, req *api.DrawRequest) (resp *api.DrawResponse, err error) {
	drawResult, err := s.service.Draw(ctx, req.User, req.Eh)
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
func (s *VegasApplication) Show(ctx context.Context, req *api.ShowRequest) (resp *api.ShowResponse, err error) {
	// TODO: Your code here...
	return
}
