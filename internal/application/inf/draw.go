package inf

import (
	"context"
	"github.com/zlllgp/vegas/internal/domain/entity"
	"github.com/zlllgp/vegas/internal/domain/service"
	"github.com/zlllgp/vegas/kitex_gen/api"
)

type DrawService interface {
	Draw(ctx context.Context, user *api.User, eh string) (*entity.DrawResult, error)
}

var _ DrawService = (*service.DrawService)(nil)
