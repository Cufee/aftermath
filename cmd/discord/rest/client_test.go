package rest

import (
	"errors"
	"fmt"
	"testing"
)

func TestShouldCountDiscordFailure(t *testing.T) {
	tcs := []struct {
		name string
		err  error
		want bool
	}{
		{name: "nil", err: nil, want: false},
		{name: "unknown webhook", err: ErrUnknownWebhook, want: false},
		{name: "unknown interaction wrapped", err: fmt.Errorf("wrapped: %w", ErrUnknownInteraction), want: false},
		{name: "already acked", err: ErrInteractionAlreadyAcked, want: false},
		{name: "missing permissions", err: ErrMissingPermissions, want: false},
		{name: "user unreachable", err: ErrMissingUserUnreachable, want: false},
		{name: "generic", err: errors.New("boom"), want: true},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			if got := shouldCountDiscordFailure(tc.err); got != tc.want {
				t.Fatalf("shouldCountDiscordFailure(%v) = %v, want %v", tc.err, got, tc.want)
			}
		})
	}
}
