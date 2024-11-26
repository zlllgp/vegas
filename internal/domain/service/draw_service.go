package service

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/zlllgp/vegas/internal/domain/entity"
	"github.com/zlllgp/vegas/internal/domain/repository"
	"github.com/zlllgp/vegas/kitex_gen/api"
	"github.com/zlllgp/vegas/pkg/errno"
	"strconv"
	"time"
)

type DrawService struct {
	rmbRepo   repository.RmbRepository
	actRepo   repository.ActivityRepository
	redisRepo repository.RedisActivityRepository
	cacheRepo repository.CacheRepository
}

func NewDrawService(
	rmbRepo repository.RmbRepository,
	actRepo repository.ActivityRepository,
	redisRepo repository.RedisActivityRepository,
	cacheRepo repository.CacheRepository) *DrawService {
	return &DrawService{
		rmbRepo: rmbRepo, actRepo: actRepo, redisRepo: redisRepo, cacheRepo: cacheRepo,
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
	safe := s.rmbRepo.IsSafe(ctx, userId)
	if !safe {
		return nil, errno.DrawErr
	}
	var activity *entity.Activity
	// first local cache
	cacheActivity, ok := s.cacheRepo.Get("activity:" + strconv.FormatInt(activityId, 10))
	if !ok {
		//second redis cache
		rActivity, err := s.redisRepo.GetActivity(ctx, activityId)
		// third db
		if errors.Is(err, errno.ActivityNotFoundErr) {
			activityDo, err := s.actRepo.GetById(ctx, activityId)
			if err != nil {
				return nil, errors.New("activity not found")
			}
			dActivity := entity.NewActivityWithModel(activityDo)
			activity = dActivity
			klog.Infof("activity from db , %+v", dActivity)
			s.redisRepo.StoreActivity(ctx, activityId, dActivity)
			s.cacheRepo.Set("activity:"+strconv.FormatInt(activityId, 10), dActivity, time.Until(time.Now()))
		} else {
			activity = rActivity
			klog.Infof("activity from redis ,  %+v", rActivity)
			s.cacheRepo.Set("activity:"+strconv.FormatInt(activityId, 10), rActivity, time.Until(time.Now()))
		}
	} else {
		if b, ok := cacheActivity.([]uint8); ok {
			if err := json.Unmarshal(b, &activity); err != nil {
				klog.Errorf("unmarshal activity err:%v", err)
			}
		}
		klog.Infof("activity from cache ,  %+v", activity)
	}

	// has wins rights

	// mock do draw
	klog.Infof("draw , planId : %d", planId)
	return &entity.DrawResult{
		Rights: []*api.Right{{Num: 1, Id: 1, Amt: "100"}},
		Time:   time.Now(),
		Err:    nil,
	}, nil
}
