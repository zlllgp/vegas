package repository

import "context"

//go:generate mockgen -source=rmb_repository.go -destination=mock_repository/rmb_repository_mock.go -package=mock_repository RmbRepository
type RmbRepository interface {
	IsSafe(ctx context.Context, userId int64) bool
}
