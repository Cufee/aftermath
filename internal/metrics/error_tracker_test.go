package metrics

import (
	"testing"
	"time"
)

func TestErrorTrackerTransitionsAndAggregation(t *testing.T) {
	cfg := Config{
		Window:             time.Minute,
		BucketInterval:     time.Second,
		EvaluationInterval: time.Second,
		MinRequests:        4,
		WarningRate:        0.30,
		CriticalRate:       0.60,
		RecoveryRate:       0.15,
	}
	tracker := NewErrorTracker(cfg)

	var transitions []Transition
	tracker.RegisterSink(func(t Transition) {
		transitions = append(transitions, t)
	})

	now := time.Now()
	for range 3 {
		tracker.Record("wargaming", "account_by_id", true)
	}
	tracker.evaluate(now)
	if tracker.OverallState() != HealthStateHealthy {
		t.Fatalf("expected healthy state before min sample threshold")
	}
	if len(transitions) != 0 {
		t.Fatalf("unexpected transition before min sample threshold: %d", len(transitions))
	}

	tracker.Record("wargaming", "account_by_id", true) // 4/4 => critical
	tracker.evaluate(now)
	if tracker.OverallState() != HealthStateCritical {
		t.Fatalf("expected critical state, got %s", tracker.OverallState())
	}

	for range 4 {
		tracker.Record("discord", "create_message", false)
	}
	tracker.Record("discord", "create_message", true)
	tracker.Record("discord", "create_message", true) // 2/6 => warning-ish for source
	tracker.evaluate(now)
	if tracker.OverallState() != HealthStateCritical {
		t.Fatalf("expected worst-state aggregation to stay critical")
	}

	for range 4 { // wargaming 4/8 => warning
		tracker.Record("wargaming", "account_by_id", false)
	}
	tracker.evaluate(now)
	if tracker.OverallState() != HealthStateWarning {
		t.Fatalf("expected warning state, got %s", tracker.OverallState())
	}

	for range 24 { // wargaming 4/32 => healthy
		tracker.Record("wargaming", "account_by_id", false)
	}
	for range 20 { // discord 2/26 => healthy
		tracker.Record("discord", "create_message", false)
	}
	tracker.evaluate(now)
	if tracker.OverallState() != HealthStateHealthy {
		t.Fatalf("expected recovered healthy state, got %s", tracker.OverallState())
	}

	if len(transitions) != 3 {
		t.Fatalf("expected exactly 3 transitions, got %d", len(transitions))
	}

	if transitions[0].From != HealthStateHealthy || transitions[0].To != HealthStateCritical {
		t.Fatalf("unexpected first transition: %+v", transitions[0])
	}
	if transitions[1].From != HealthStateCritical || transitions[1].To != HealthStateWarning {
		t.Fatalf("unexpected second transition: %+v", transitions[1])
	}
	if transitions[2].From != HealthStateWarning || transitions[2].To != HealthStateHealthy {
		t.Fatalf("unexpected third transition: %+v", transitions[2])
	}
}
