package service

import (
	"context"
	"github.com/zlllgp/vegas/internal/dal/model"
	"github.com/zlllgp/vegas/internal/dal/query"
	"github.com/zlllgp/vegas/kitex_gen/api"
)

type VegasService struct {
	ctx context.Context
}

func NewVegasService(ctx context.Context) *VegasService {
	return &VegasService{ctx: ctx}
}

func (s *VegasService) Draw(req *api.DrawRequest) (resp *api.DrawResult, err error) {
	// TODO: Your code here...
	activity := &model.Activity{CreatorId: 1, CreatorName: "test001", Name: "测试活动"}
	query.Activity.Create(&activity)
	resp = &api.DrawResult{Code: "SUCCESS", Msg: "", Rights: &api.RightsDTO{Id: 1, Num: 1, Amt: "10"}}
	return
}
