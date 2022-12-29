package mock_time

import "time"

type realClock struct {
}

func NewRealClock() Clock {
	return &realClock{}
}

func (m *realClock) Now(now time.Time) time.Time {
	return time.Now()
}
