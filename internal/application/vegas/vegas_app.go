package vegas

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/zlllgp/vegas/internal/infrastructure/dal/wk/query"
	"github.com/zlllgp/vegas/internal/infrastructure/repository"
	"github.com/zlllgp/vegas/kitex_gen/api"
)

type VegasServiceImpl struct {
	rmbRep *repository.RmbRepository
}

func NewVegasServiceImpl(rmbRep *repository.RmbRepository) *VegasServiceImpl {
	return &VegasServiceImpl{rmbRep: rmbRep}
}

// Draw implements the VegasServiceImpl interface.
func (s *VegasServiceImpl) Draw(ctx context.Context, req *api.DrawRequest) (resp *api.DrawResponse, err error) {
	// judge safe
	safe := s.rmbRep.IsSafe(req.User.UserId)
	if !safe {
		return
	}
	// judge activity

	// has wins rights

	// do draw

	q := query.Q.Activity
	// activity
	activity, _ := q.Select(q.ALL).Where(q.ID.In(1)).First()
	klog.Infof("query activity : %+v", activity)

	// create
	/*activityCrate := model.Activity{CreatorId: 1, CreatorName: "test002", Name: "测试活动3", TenantId: 1}
	err = q.Create(&activityCrate)
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

	right := &api.Right{Id: 1, Num: 1, Amt: "10"}
	resp = &api.DrawResponse{Code: "SUCCESS", Msg: "", Rights: []*api.Right{right}}
	return resp, nil
}

// Show implements the VegasServiceImpl interface.
func (s *VegasServiceImpl) Show(ctx context.Context, req *api.ShowRequest) (resp *api.ShowResponse, err error) {
	// TODO: Your code here...
	return
}
