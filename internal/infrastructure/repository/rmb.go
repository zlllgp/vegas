package repository

type RmbRepository struct{}

func (rr *RmbRepository) IsSafe(userId int64) bool {
	return true
}
