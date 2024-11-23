package inf

import (
	"context"
	"github.com/zlllgp/vegas/internal/domain/entity"
	"github.com/zlllgp/vegas/internal/domain/service"
	"github.com/zlllgp/vegas/kitex_gen/api"
)

//go:generate mockgen -source=right_inf.go -destination=right_inf_mock.go -package=inf RightService
type RightService interface {
	Find(ctx context.Context, userId int64, planId int64, pageRequest *api.PaginationRequest, searchFields ...*api.SearchField) ([]*entity.Right, error)
}

var _ RightService = (*service.RightService)(nil)
