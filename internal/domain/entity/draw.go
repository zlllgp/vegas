package entity

import (
	"github.com/zlllgp/vegas/kitex_gen/api"
	"time"
)

type DrawResult struct {
	Rights []*api.Right
	Time   time.Time
	Err    error
}
