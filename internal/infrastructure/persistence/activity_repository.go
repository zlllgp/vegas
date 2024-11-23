package persistence

import (
	"context"
	"github.com/zlllgp/vegas/internal/domain/entity"
	"github.com/zlllgp/vegas/internal/infrastructure/dal/wk/model"
	"github.com/zlllgp/vegas/internal/infrastructure/dal/wk/query"
	"gorm.io/gen"
)

type ActivityRepository struct{}

func NewActivityRepository() *ActivityRepository {
	return &ActivityRepository{}
}

func (r *ActivityRepository) getFirst(ctx context.Context, conditions ...gen.Condition) (*model.Activity, error) {
	conditions = append(conditions, query.Q.Activity.IsDel.In(0))
	activity, err := query.Q.WithContext(ctx).Activity.Where(conditions...).First()
	if err != nil {
		return nil, err
	}
	return activity, nil
}

func (r *ActivityRepository) GetById(ctx context.Context, id int64) (*model.Activity, error) {
	return r.getFirst(ctx, query.Q.Activity.ID.Eq(id))
}

func (r *ActivityRepository) CreateActivity(ctx context.Context, name string) error {
	activity := entity.NewActivity(
		entity.WithName(name))
	return query.Q.WithContext(ctx).Activity.Create(activity.ToModel())
}
