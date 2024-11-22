package repository

import (
	"context"
	"github.com/zlllgp/vegas/internal/infrastructure/dal/wk/model"
	"github.com/zlllgp/vegas/internal/infrastructure/dal/wk/query"
	"gorm.io/gen"
)

type ActivityRepository struct{}

func NewActivityRepository() *ActivityRepository {
	return &ActivityRepository{}
}

func (ar *ActivityRepository) getFirst(ctx context.Context, conditions ...gen.Condition) (*model.Activity, error) {
	conditions = append(conditions, query.Q.Activity.IsDel)
	return nil, nil
}

func (ar *ActivityRepository) GetById(ctx context.Context, id int64) (*model.Activity, error) {
	//return ar.GetById(ctx)
	return nil, nil
}
