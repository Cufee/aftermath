package alerts

import (
	"encoding/json"
	"io"
	"testing"
	"time"

	"github.com/rs/zerolog"
)

func TestLogWriterDoesNotBlockWhenFull(t *testing.T) {
	w := NewLogWriter(1)

	if _, err := w.Write([]byte(`{"message":"first"}`)); err != nil {
		t.Fatal(err)
	}

	wrote := make(chan struct{})
	go func() {
		_, _ = w.Write([]byte(`{"message":"second"}`))
		close(wrote)
	}()

	select {
	case <-wrote:
	case <-time.After(time.Second):
		t.Fatal("write blocked when alert log buffer was full")
	}

	if dropped := w.Dropped(); dropped != 1 {
		t.Fatalf("expected one dropped log record, got %d", dropped)
	}
}

func TestLogWriterPreservesRecordsAcrossSmallReads(t *testing.T) {
	w := NewLogWriter(1)
	record := []byte(`{"level":"error","message":"this record is larger than the read buffer"}` + "\n")

	if _, err := w.Write(record); err != nil {
		t.Fatal(err)
	}

	var got []byte
	buf := make([]byte, 7)
	for len(got) < len(record) {
		n, err := w.Read(buf)
		if err != nil && err != io.EOF {
			t.Fatal(err)
		}
		got = append(got, buf[:n]...)
	}

	if string(got) != string(record) {
		t.Fatalf("record changed across reads: %q", got)
	}
}

func TestLogWriterAllowsReaderToLogWithoutDeadlocking(t *testing.T) {
	w := NewLogWriter(1)
	logger := zerolog.New(w)

	readerReturned := make(chan struct{})
	go func() {
		dec := json.NewDecoder(w)
		var event map[string]any
		if err := dec.Decode(&event); err != nil {
			t.Errorf("decode first log event: %v", err)
			return
		}

		logger.Info().Msg("nested log from alert reader")
		close(readerReturned)
	}()

	logger.Error().Msg("first event")

	select {
	case <-readerReturned:
	case <-time.After(time.Second):
		t.Fatal("alert reader blocked while logging back through LogWriter")
	}

	futureWriteReturned := make(chan struct{})
	go func() {
		logger.Info().Msg("future access log")
		close(futureWriteReturned)
	}()

	select {
	case <-futureWriteReturned:
	case <-time.After(time.Second):
		t.Fatal("future log write blocked after alert reader logged")
	}
}
