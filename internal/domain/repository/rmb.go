package repository

import "context"

type RmbRepository interface {
	IsSafe(ctx context.Context, userId int64) bool
}
