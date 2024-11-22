package rpc

import "context"

type RmbRepository struct{}

func NewRmbRepository() *RmbRepository {
	return &RmbRepository{}
}

func (r *RmbRepository) IsSafe(ctx context.Context, userId int64) bool {
	return true
}
