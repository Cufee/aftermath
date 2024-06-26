package stats

import (
	"time"

	"github.com/cufee/aftermath/internal/stats/fetch"
)

type Metadata struct {
	Stats   fetch.AccountStatsOverPeriod
	Timings map[string]time.Duration
	timers  map[string]time.Time
}

func (m *Metadata) Timer(name string) func() {
	if m.timers == nil {
		m.timers = make(map[string]time.Time)
	}
	if m.Timings == nil {
		m.Timings = make(map[string]time.Duration)
	}

	m.timers[name] = time.Now()

	return func() {
		m.Timings[name] = time.Since(m.timers[name]).Round(time.Millisecond)
	}
}

func (m *Metadata) TotalTime() time.Duration {
	var total time.Duration
	for _, duration := range m.Timings {
		total += duration
	}
	return total.Round(time.Millisecond)
}
