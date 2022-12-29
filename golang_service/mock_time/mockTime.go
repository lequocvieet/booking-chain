package mock_time

import "time"

type mockClock struct {
}

func NewMockClock() Clock {
	return &mockClock{}
}

func (m *mockClock) Now(fakeNow time.Time) time.Time {
	return fakeNow
}
