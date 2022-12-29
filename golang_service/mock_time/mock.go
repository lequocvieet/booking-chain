package mock_time

import "time"

type Clock interface {
	Now(now time.Time) time.Time
	//After(d time.Duration) <-chan time.Time
}
