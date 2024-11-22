package inf

import (
	"context"
	"github.com/zlllgp/vegas/internal/domain/entity"
	"github.com/zlllgp/vegas/internal/domain/service"
)

type DrawService interface {
	Draw(ctx context.Context, userId int64, eh string) (*entity.DrawResult, error)
}

var _ DrawService = (*service.DrawService)(nil)
