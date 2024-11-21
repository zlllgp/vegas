package service

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/zlllgp/vegas/internal/dal/wk/query"
	"github.com/zlllgp/vegas/kitex_gen/api"
)

type VegasService struct {
	ctx context.Context
}

func NewVegasService(ctx context.Context) *VegasService {
	return &VegasService{ctx: ctx}
}

func (s *VegasService) Draw(req *api.DrawRequest) (resp *api.DrawResult, err error) {
	u := query.Q.Activity
	// create
	/*activityCrate := model.Activity{CreatorId: 1, CreatorName: "test002", Name: "测试活动3", TenantId: 1}
	err = u.Create(&activityCrate)
	if err != nil {
		return nil, err
	}*/

	// transaction todo fix
	/*q := query.Use(db)
	activityTx := model.Activity{CreatorId: 1, CreatorName: "test002", Name: "测试活动4", TenantId: 1}
	err = q.Transaction(func(tx *query.Query) error {
		tx.Activity.Create(&activityTx)
		return nil
	})
	if err != nil {
		return nil, err
	}*/

	// redis

	// query
	activityResult, _ := u.Select(u.ALL).Where(u.ID.In(1)).First()
	klog.Infof("query activity : %T", activityResult)
	klog.Infof("eh : %s , userName : %s", req.Eh, req.UserName)
	resp = &api.DrawResult{Code: "SUCCESS", Msg: "", Rights: &api.RightsDTO{Id: 1, Num: 1, Amt: "10"}}
	return resp, nil
}
