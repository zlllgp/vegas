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

type ShowResult struct {
	Rights []*api.Right
	Time   time.Time
	Err    error
}

type RightResult struct {
	Rights []*api.Right
}
