package service

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/spf13/cast"
	"github.com/zlllgp/vegas/internal/domain/entity"
	"github.com/zlllgp/vegas/internal/domain/repository"
	"github.com/zlllgp/vegas/internal/infrastructure/cache"
	"github.com/zlllgp/vegas/pkg/errno"
	"strconv"
	"time"
)

type ActivityService interface {
	GetActivity(ctx context.Context, activityId int64) (*entity.Activity, error)
}

type ActivityServiceImpl struct {
	actRepo   repository.ActivityRepository
	redisRepo repository.RedisActivityRepository
	cacheRepo repository.CacheRepository
}

func NewActivityServiceImpl(
	actRepo repository.ActivityRepository,
	redisRepo repository.RedisActivityRepository,
	cacheRepo repository.CacheRepository) *ActivityServiceImpl {
	return &ActivityServiceImpl{actRepo: actRepo, redisRepo: redisRepo, cacheRepo: cacheRepo}
}

func (s *ActivityServiceImpl) GetActivity(ctx context.Context, activityId int64) (*entity.Activity, error) {
	return s.cacheGetActivity(ctx, activityId)

	//return s.localCacheGetActivity(ctx, activityId)
}

func (s *ActivityServiceImpl) localCacheGetActivity(ctx context.Context, activityId int64) (*entity.Activity, error) {
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

	return activity, nil
}

func (s *ActivityServiceImpl) cacheGetActivity(ctx context.Context, activityId int64) (*entity.Activity, error) {
	var activity *entity.Activity
	// use rcache with redis + db

	activityStr, err := cache.RCache.Fetch("activity:"+strconv.FormatInt(activityId, 10), time.Hour, func() (string, error) {
		activityDo, err := s.actRepo.GetById(ctx, activityId)
		klog.Infof("activity from db , %+v", activityDo)
		if err != nil || activityDo == nil {
			return "", errors.New("activity not found")
		}
		return cast.ToString(activityDo), nil
	})
	if err != nil {
		json.Unmarshal([]byte(activityStr), &activity)
	}
	return activity, nil
}
