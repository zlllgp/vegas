package inf

import (
	"context"
	"github.com/zlllgp/vegas/internal/domain/entity"
	"github.com/zlllgp/vegas/internal/domain/service"
)

//go:generate mockgen -source=draw_inf.go -destination=draw_inf_mock.go -package=inf DrawService
type DrawService interface {
	Draw(ctx context.Context, userId int64, eh string) (*entity.DrawResult, error)
}

var _ DrawService = (*service.DrawService)(nil)
