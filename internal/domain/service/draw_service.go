package service

import (
	"context"
	"errors"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/zlllgp/vegas/internal/domain/entity"
	"github.com/zlllgp/vegas/internal/domain/repository"
	"github.com/zlllgp/vegas/kitex_gen/api"
	"strconv"
	"time"
)

type DrawService struct {
	rmbRep repository.RmbRepository
	actRep repository.ActivityRepository
}

func NewDrawService(rmbRep repository.RmbRepository, actRep repository.ActivityRepository) *DrawService {
	return &DrawService{
		rmbRep: rmbRep,
		actRep: actRep,
	}
}

func (s *DrawService) Draw(ctx context.Context, userId int64, eh string) (*entity.DrawResult, error) {
	//decode eh to activityId , planId
	activityId, err := strconv.ParseInt(eh, 10, 64)
	planId, err := strconv.ParseInt(eh, 10, 64)
	if err != nil {
		return nil, errors.New("id is not int")
	}

	// judge safe
	safe := s.rmbRep.IsSafe(ctx, userId)
	if !safe {
		return nil, errors.New("not safe")
	}
	// judge activity
	activityDo, err := s.actRep.GetById(ctx, activityId)
	if err != nil {
		return nil, errors.New("activity not found")
	}
	activity := entity.NewActivityWithModel(activityDo)
	klog.Infof("activity id : %d, name : %s", activity.ID, activity.Name)
	// has wins rights

	// mock do draw
	klog.Infof("draw , planId : %d", planId)
	return &entity.DrawResult{
		Rights: []*api.Right{{Num: 1, Id: 1, Amt: "100"}},
		Time:   time.Now(),
		Err:    nil,
	}, nil
}
