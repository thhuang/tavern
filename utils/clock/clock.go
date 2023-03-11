package clock

import "time"

type Clock interface {
	Now() time.Time
}

type fixedClock struct {
	now time.Time
}

func New(now time.Time) *fixedClock {
	return &fixedClock{now: now}
}

func (c *fixedClock) Now() time.Time {
	return c.now.UTC()
}
