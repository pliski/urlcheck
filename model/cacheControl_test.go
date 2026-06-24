package model

import (
	"io"
	"net/http"
	"testing"
)

// fakeRoundTripper records the request it receives and returns 200 OK.
type fakeRoundTripper struct{ got *http.Request }

func (f *fakeRoundTripper) RoundTrip(r *http.Request) (*http.Response, error) {
	f.got = r
	return &http.Response{StatusCode: http.StatusOK, Body: http.NoBody}, nil
}

func TestLoggingRoundTripperSetsNoCacheHeaders(t *testing.T) {
	fake := &fakeRoundTripper{}
	logger := &StdLogger{out: io.Discard, info: io.Discard, err: io.Discard}
	rt := NewLoggingRoundTripper(fake, logger)

	req, _ := http.NewRequest(http.MethodHead, "http://example.com", nil)
	if _, err := rt.RoundTrip(req); err != nil {
		t.Fatalf("RoundTrip: %v", err)
	}

	// The injected transport must actually be used (regression test for the
	// previously hardcoded http.DefaultTransport).
	if fake.got == nil {
		t.Fatal("next transport was not called")
	}
	if got := fake.got.Header.Get("Cache-Control"); got != "no-cache" {
		t.Errorf("Cache-Control = %q, want %q", got, "no-cache")
	}
	if got := fake.got.Header.Get("Pragma"); got != "no-cache" {
		t.Errorf("Pragma = %q, want %q", got, "no-cache")
	}
	// The caller's request must not be mutated (RoundTripper clone contract).
	if req.Header.Get("Cache-Control") != "" {
		t.Error("caller's request was mutated; expected it to be cloned")
	}
}
