package service

import (
	"context"
	"errors"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/zlllgp/vegas/internal/domain/entity"
	"github.com/zlllgp/vegas/internal/domain/repository"
	"github.com/zlllgp/vegas/kitex_gen/api"
	"github.com/zlllgp/vegas/pkg/errno"
	"strconv"
	"time"
)

type DrawService interface {
	Draw(ctx context.Context, userId int64, eh string) (*entity.DrawResult, error)
}

type DrawServiceImpl struct {
	rmbRepo    repository.RmbRepository
	actService ActivityService
}

func NewDrawServiceImpl(
	rmbRepo repository.RmbRepository,
	actService ActivityService) *DrawServiceImpl {
	return &DrawServiceImpl{
		rmbRepo: rmbRepo, actService: actService,
	}
}

func (s *DrawServiceImpl) Draw(ctx context.Context, userId int64, eh string) (*entity.DrawResult, error) {
	//decode eh to activityId , planId
	activityId, err := strconv.ParseInt(eh, 10, 64)
	if err != nil {
		return nil, errors.New("id is not int")
	}

	// judge safe
	safe := s.rmbRepo.IsSafe(ctx, userId)
	if !safe {
		return nil, errno.DrawErr
	}

	activity, err := s.actService.GetActivity(ctx, activityId)
	if err != nil {
		return nil, errno.ActivityNotFoundErr
	}
	klog.Infof("activity:%+v", activity)

	// wins rights

	// mock do draw
	return &entity.DrawResult{
		Rights: []*api.Right{{Num: 1, Id: 1, Amt: "100"}},
		Time:   time.Now(),
		Err:    nil,
	}, nil
}
