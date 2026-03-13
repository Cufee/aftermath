package metrics

import (
	"context"
	"math"
	"sort"
	"sync/atomic"
	"time"
)

type HealthState string

const (
	HealthStateHealthy  HealthState = "healthy"
	HealthStateWarning  HealthState = "warning"
	HealthStateCritical HealthState = "critical"
)

type ErrorObserver interface {
	Record(source, operation string, failed bool)
}

type SourceHealth struct {
	Source    string
	State     HealthState
	Total     int
	Failed    int
	ErrorRate float64
}

type Transition struct {
	At      time.Time
	From    HealthState
	To      HealthState
	Sources []SourceHealth
}

type Config struct {
	Window             time.Duration
	BucketInterval     time.Duration
	EvaluationInterval time.Duration

	MinRequests int

	WarningRate  float64
	CriticalRate float64
	RecoveryRate float64
}

func DefaultErrorTrackerConfig() Config {
	return Config{
		Window:             time.Minute * 5,
		BucketInterval:     time.Second * 10,
		EvaluationInterval: time.Second * 10,
		MinRequests:        40,
		WarningRate:        0.30,
		CriticalRate:       0.60,
		RecoveryRate:       0.15,
	}
}

type bucketStats struct {
	Total  int
	Failed int
}

type errorBucket struct {
	Start   time.Time
	Sources map[string]bucketStats
}

type recordEvent struct {
	source string
	failed bool
	at     time.Time
}

type ErrorTracker struct {
	cfg     Config
	records chan recordEvent

	buckets     []errorBucket
	sourceState map[string]HealthState
	sinks       []func(Transition)

	overall atomic.Value
}

func NewErrorTracker(cfg Config) *ErrorTracker {
	cfg = normalizeConfig(cfg)
	bucketsCount := int(math.Ceil(float64(cfg.Window) / float64(cfg.BucketInterval)))
	if bucketsCount < 1 {
		bucketsCount = 1
	}

	buckets := make([]errorBucket, bucketsCount)
	for i := range buckets {
		buckets[i] = errorBucket{Sources: make(map[string]bucketStats)}
	}

	t := &ErrorTracker{
		cfg:         cfg,
		records:     make(chan recordEvent, 4096),
		buckets:     buckets,
		sourceState: map[string]HealthState{},
	}
	t.overall.Store(HealthStateHealthy)
	return t
}

func (t *ErrorTracker) RegisterSink(sink func(Transition)) {
	if sink == nil {
		return
	}
	t.sinks = append(t.sinks, sink)
}

func (t *ErrorTracker) Record(source, _ string, failed bool) {
	if source == "" {
		source = "unknown"
	}
	select {
	case t.records <- recordEvent{source: source, failed: failed, at: time.Now()}:
	default:
	}
}

func (t *ErrorTracker) Start(ctx context.Context) {
	ticker := time.NewTicker(t.cfg.EvaluationInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case rec := <-t.records:
			t.applyRecord(rec)
		case now := <-ticker.C:
			t.evaluate(now)
		}
	}
}

func (t *ErrorTracker) OverallState() HealthState {
	if v := t.overall.Load(); v != nil {
		return v.(HealthState)
	}
	return HealthStateHealthy
}

func (t *ErrorTracker) applyRecord(rec recordEvent) {
	bucket := t.currentBucket(rec.at)
	stats := bucket.Sources[rec.source]
	stats.Total++
	if rec.failed {
		stats.Failed++
	}
	bucket.Sources[rec.source] = stats
}

func (t *ErrorTracker) drainRecords() {
	for {
		select {
		case rec := <-t.records:
			t.applyRecord(rec)
		default:
			return
		}
	}
}

func (t *ErrorTracker) evaluate(now time.Time) {
	t.drainRecords()

	statsBySource := t.windowStats(now)
	nextSourceState := make(map[string]HealthState, len(statsBySource))

	sources := make([]SourceHealth, 0, len(statsBySource))
	overallState := HealthStateHealthy

	for source, stats := range statsBySource {
		var rate float64
		if stats.Total > 0 {
			rate = float64(stats.Failed) / float64(stats.Total)
		}

		state := stateForRate(t.sourceState[source], stats.Total, rate, t.cfg)
		nextSourceState[source] = state
		if severity(state) > severity(overallState) {
			overallState = state
		}

		sources = append(sources, SourceHealth{
			Source:    source,
			State:     state,
			Total:     stats.Total,
			Failed:    stats.Failed,
			ErrorRate: rate,
		})
	}
	sort.Slice(sources, func(i, j int) bool { return sources[i].Source < sources[j].Source })

	transition := Transition{
		At:      now,
		From:    t.overall.Load().(HealthState),
		To:      overallState,
		Sources: sources,
	}

	t.sourceState = nextSourceState
	t.overall.Store(overallState)

	if transition.From == transition.To {
		return
	}

	for _, sink := range t.sinks {
		sink(transition)
	}
}

func (t *ErrorTracker) windowStats(now time.Time) map[string]bucketStats {
	cutoff := now.Add(-t.cfg.Window)
	statsBySource := make(map[string]bucketStats)

	for _, bucket := range t.buckets {
		if bucket.Start.IsZero() || bucket.Start.Before(cutoff) {
			continue
		}

		for source, stats := range bucket.Sources {
			agg := statsBySource[source]
			agg.Total += stats.Total
			agg.Failed += stats.Failed
			statsBySource[source] = agg
		}
	}
	return statsBySource
}

func (t *ErrorTracker) currentBucket(now time.Time) *errorBucket {
	slot := int((now.UnixNano() / int64(t.cfg.BucketInterval)) % int64(len(t.buckets)))
	start := now.Truncate(t.cfg.BucketInterval)
	bucket := &t.buckets[slot]
	if !bucket.Start.Equal(start) {
		bucket.Start = start
		clear(bucket.Sources)
	}
	return bucket
}

func stateForRate(previous HealthState, total int, rate float64, cfg Config) HealthState {
	if total < cfg.MinRequests {
		return HealthStateHealthy
	}

	switch previous {
	case HealthStateCritical:
		if rate <= cfg.RecoveryRate {
			return HealthStateHealthy
		}
		if rate >= cfg.CriticalRate {
			return HealthStateCritical
		}
		return HealthStateWarning

	case HealthStateWarning:
		if rate <= cfg.RecoveryRate {
			return HealthStateHealthy
		}
		if rate >= cfg.CriticalRate {
			return HealthStateCritical
		}
		return HealthStateWarning

	default:
		if rate >= cfg.CriticalRate {
			return HealthStateCritical
		}
		if rate >= cfg.WarningRate {
			return HealthStateWarning
		}
		return HealthStateHealthy
	}
}

func normalizeConfig(cfg Config) Config {
	def := DefaultErrorTrackerConfig()

	if cfg.Window <= 0 {
		cfg.Window = def.Window
	}
	if cfg.BucketInterval <= 0 {
		cfg.BucketInterval = def.BucketInterval
	}
	if cfg.EvaluationInterval <= 0 {
		cfg.EvaluationInterval = def.EvaluationInterval
	}
	if cfg.MinRequests <= 0 {
		cfg.MinRequests = def.MinRequests
	}
	if cfg.WarningRate <= 0 || cfg.WarningRate >= 1 {
		cfg.WarningRate = def.WarningRate
	}
	if cfg.CriticalRate <= 0 || cfg.CriticalRate >= 1 {
		cfg.CriticalRate = def.CriticalRate
	}
	if cfg.RecoveryRate < 0 || cfg.RecoveryRate >= 1 {
		cfg.RecoveryRate = def.RecoveryRate
	}
	if cfg.CriticalRate < cfg.WarningRate {
		cfg.CriticalRate = cfg.WarningRate
	}
	if cfg.RecoveryRate > cfg.WarningRate {
		cfg.RecoveryRate = def.RecoveryRate
	}

	return cfg
}

func severity(state HealthState) int {
	switch state {
	case HealthStateCritical:
		return 2
	case HealthStateWarning:
		return 1
	default:
		return 0
	}
}
