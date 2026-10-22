package scheduler

import (
	"context"
	"testing"
	"time"
)

func TestDefaultSchedulesUseUTC(t *testing.T) {
	originalLocation := time.Local
	time.Local = time.FixedZone("test-offset", -7*60*60)
	defer func() { time.Local = originalLocation }()
	s, err := New()
	if err != nil {
		t.Fatal(err)
	}
	if err := RegisterDefaultTasks(s, nil); err != nil {
		t.Fatal(err)
	}
	if len(s.jobs) != 9 {
		t.Fatalf("registered %d jobs, want 9", len(s.jobs))
	}
	// Replace workers before starting: checking scheduling must not contact services.
	for _, job := range s.jobs {
		if err := s.cron.RemoveJob(job.ID()); err != nil {
			t.Fatal(err)
		}
	}
	s.jobs = nil
	job, err := s.Add("0 9 * * *", func() {})
	if err != nil {
		t.Fatal(err)
	}
	stop, err := s.Start(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	defer stop()
	next, err := job.NextRun()
	if err != nil {
		t.Fatal(err)
	}
	if next.Location() != time.UTC || next.Hour() != 9 || next.Minute() != 0 {
		t.Fatalf("snapshot scheduled at %v, want 09:00 UTC", next)
	}
}
