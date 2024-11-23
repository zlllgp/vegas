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

func (r *ActivityRepository) Create(ctx context.Context, name string) error {
	activity := entity.NewActivity(
		entity.WithName(name))
	return query.Q.WithContext(ctx).Activity.Create(activity.ToModel())
}

func (r *ActivityRepository) GetById(ctx context.Context, id int64) (*model.Activity, error) {
	return r.getFirst(ctx, query.Q.Activity.ID.Eq(id))
}

func (r *ActivityRepository) UpdateName(ctx context.Context, activity *model.Activity) error {
	_, err := query.Q.WithContext(ctx).Activity.UpdateColumns(activity)
	return err
}

func (r *ActivityRepository) IsNameExist(ctx context.Context, name string) (bool, error) {
	return r.isExist(ctx, query.Q.Activity.Name.Eq(name))
}

func (r *ActivityRepository) count(ctx context.Context, conditions ...gen.Condition) (int64, error) {
	conditions = append(conditions, query.Q.Activity.IsDel.In(0))
	return query.Q.WithContext(ctx).Activity.Where(conditions...).Count()
}

func (r *ActivityRepository) isExist(ctx context.Context, conditions ...gen.Condition) (bool, error) {
	count, err := r.count(ctx, conditions...)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
