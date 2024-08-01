package use_time

import (
	"time"
)

type UseTime struct {
	t time.Time
}

func New(begin time.Time) *UseTime {
	return &UseTime{t: begin}
}

func (u *UseTime) Step() time.Duration {
	return time.Since(u.t)
}
