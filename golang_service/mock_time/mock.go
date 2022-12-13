package mock_time

import "time"

type Clock interface {
	Now() time.Time
	After(d time.Duration) <-chan time.Time
}
