package persistence

import (
	"context"
	"github.com/zlllgp/vegas/internal/infrastructure/dal"
	"github.com/zlllgp/vegas/internal/infrastructure/dal/yoda/model"
	"github.com/zlllgp/vegas/internal/infrastructure/dal/yoda/query"
	"gorm.io/gen"
)

type RightRepository struct{}

func NewRightRepository() *RightRepository {
	return &RightRepository{}
}

func (r *RightRepository) Query(ctx context.Context, offset, limit int, conditions ...gen.Condition) ([]*model.UserRight, error) {
	return r.list(ctx, r.findByPage(offset, limit), conditions...)
}

func (r *RightRepository) GetById(ctx context.Context, userId int64) (*model.UserRight, error) {
	return r.first(ctx, query.Q.UserRight.UserID.Eq(userId))
}

func (r *RightRepository) list(ctx context.Context, queryFunc func(queryDo dal.QueryDo[*model.UserRight]) ([]*model.UserRight, error),
	conditions ...gen.Condition) ([]*model.UserRight, error) {
	handler := query.Q.WithContext(ctx).UserRight.Where(
		conditions...,
	)
	return queryFunc(handler)
}

func (r *RightRepository) findByPage(offset, limit int) func(queryDo dal.QueryDo[*model.UserRight]) ([]*model.UserRight, error) {
	return func(queryDo dal.QueryDo[*model.UserRight]) ([]*model.UserRight, error) {
		result, _, err := queryDo.FindByPage(offset, limit)
		return result, err
	}
}

func (r *RightRepository) first(ctx context.Context, conditions ...gen.Condition) (*model.UserRight, error) {
	return query.Q.WithContext(ctx).UserRight.Where(
		conditions...,
	).First()
}
