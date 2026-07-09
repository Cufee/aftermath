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
			if tc.err != nil {
				if got := IsNonActionable(tc.err); got == tc.want {
					t.Fatalf("IsNonActionable(%v) = %v, want %v", tc.err, got, !tc.want)
				}
			}
		})
	}
}

func TestIsNonActionableMessage(t *testing.T) {
	tcs := []struct {
		name string
		msg  string
		want bool
	}{
		{name: "empty", msg: "", want: false},
		{name: "generic", msg: "boom", want: false},
		{name: "exact missing permissions", msg: ErrMissingPermissions.Error(), want: true},
		{name: "wrapped missing permissions", msg: "handler returned an error: discord api: missing permissions", want: true},
		{name: "unknown interaction", msg: ErrUnknownInteraction.Error(), want: true},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			if got := IsNonActionableMessage(tc.msg); got != tc.want {
				t.Fatalf("IsNonActionableMessage(%q) = %v, want %v", tc.msg, got, tc.want)
			}
		})
	}
}
