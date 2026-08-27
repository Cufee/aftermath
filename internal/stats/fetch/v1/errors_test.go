package fetch

import (
	"fmt"
	"testing"

	"github.com/cufee/am-wg-proxy-next/v2/client/common"
	"github.com/pkg/errors"
)

func TestParseWargamingErrorClassifiesUnavailableResponses(t *testing.T) {
	tests := []struct {
		name string
		err  error
	}{
		{
			name: "upstream service unavailable",
			err:  fmt.Errorf("request failed: %w", common.ErrSourceNotAvailable),
		},
		{
			name: "unexpected response content type",
			err:  fmt.Errorf("request failed: %w", common.ErrUnexpectedContentType),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := parseWargamingError(test.err); !errors.Is(got, ErrSourceNotAvailable) {
				t.Fatalf("expected ErrSourceNotAvailable, got %v", got)
			}
		})
	}
}
