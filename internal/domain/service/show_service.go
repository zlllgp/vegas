package service

import (
	"context"
	"github.com/zlllgp/vegas/internal/domain/entity"
	"github.com/zlllgp/vegas/internal/infrastructure/persistence"
)

type ShowService interface {
	Show(ctx context.Context, userId int64, eh string) (*entity.ShowResult, error)
}

type ShowServiceImpl struct {
	actRepo persistence.ActivityRepository
}

func (s *ShowServiceImpl) Show(ctx context.Context, userId int64, eh string) (*entity.ShowResult, error) {
	//TODO implement me
	panic("implement me")
}
