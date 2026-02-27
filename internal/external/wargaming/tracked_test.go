package wargaming

import (
	"context"
	"errors"
	"testing"

	"github.com/cufee/am-wg-proxy-next/v2/types"
)

type observerCall struct {
	Source    string
	Operation string
	Failed    bool
}

type observerSpy struct {
	calls []observerCall
}

func (o *observerSpy) Record(source, operation string, failed bool) {
	o.calls = append(o.calls, observerCall{
		Source:    source,
		Operation: operation,
		Failed:    failed,
	})
}

type searchOnlyClient struct {
	Client
	err error
}

func (s *searchOnlyClient) SearchAccounts(ctx context.Context, realm types.Realm, query string, opts ...types.Option) ([]types.Account, error) {
	return nil, s.err
}

func TestTrackedClientRecordsFailuresAndSuccess(t *testing.T) {
	baseErr := errors.New("boom")
	base := &searchOnlyClient{err: baseErr}
	observer := &observerSpy{}

	tracked := NewTrackedClient(base, observer)
	_, _ = tracked.SearchAccounts(context.Background(), types.RealmNorthAmerica, "name")
	if len(observer.calls) != 1 {
		t.Fatalf("expected 1 observer call, got %d", len(observer.calls))
	}
	if observer.calls[0].Source != "wargaming" || observer.calls[0].Operation != "search_accounts" || !observer.calls[0].Failed {
		t.Fatalf("unexpected observer call: %+v", observer.calls[0])
	}

	base.err = nil
	_, _ = tracked.SearchAccounts(context.Background(), types.RealmNorthAmerica, "name")
	if len(observer.calls) != 2 {
		t.Fatalf("expected 2 observer calls, got %d", len(observer.calls))
	}
	if observer.calls[1].Failed {
		t.Fatalf("expected success call to be marked as non-failure")
	}
}
