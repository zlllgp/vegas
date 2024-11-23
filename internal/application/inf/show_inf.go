package inf

import (
	"context"
	"github.com/zlllgp/vegas/internal/domain/entity"
	"github.com/zlllgp/vegas/internal/domain/service"
)

type ShowService interface {
	Show(ctx context.Context, userId int64, eh string) (*entity.ShowResult, error)
}

var _ ShowService = (*service.ShowService)(nil)
