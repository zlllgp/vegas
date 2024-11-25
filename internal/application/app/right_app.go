package app

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/zlllgp/vegas/internal/application/inf"
	"github.com/zlllgp/vegas/kitex_gen/api"
	"github.com/zlllgp/vegas/pkg/errno"
)

type RightServiceImpl struct {
	rightService inf.RightService
}

func NewRightServiceImpl(rightService inf.RightService) *RightServiceImpl {
	return &RightServiceImpl{}
}

// QueryRight implements the RightServiceImpl interface.
func (s *RightServiceImpl) QueryRight(ctx context.Context, req *api.RightRequest) (resp *api.RightResponse, err error) {
	rights, err := s.rightService.Find(ctx, req.GetUser().GetUserId(), 1, req.GetPageReq(), req.GetSfs()...)
	if err != nil {
		klog.Errorf("Query right err:%v", err)
		resp = &api.RightResponse{
			Rights: nil,
			Base: &api.BaseResponse{
				Code: errno.DataBaseErr.GetCode(),
				Msg:  errno.DataBaseErr.GetMsg()},
		}
		return resp, nil
	}

	var dtos []*api.Right
	for _, t := range rights {
		dtos = append(dtos, t.ToDTO())
	}

	return &api.RightResponse{
		Rights: dtos,
		Base: &api.BaseResponse{
			Code: errno.Success.GetCode(),
			Msg:  errno.Success.GetMsg()},
	}, nil
}
