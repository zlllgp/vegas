package repository

type RightRepository struct{}

func (rr *RightRepository) GetById(userId int64) bool {
	return true
}
