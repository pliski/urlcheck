package model

import (
	"net/http"
)

type loggingRoundTripper struct {
	next   http.RoundTripper
	logger *StdLogger
}

func NewLoggingRoundTripper(next http.RoundTripper, logger *StdLogger) *loggingRoundTripper {
	return &loggingRoundTripper{
		next:   next,
		logger: logger,
	}
}

func (t loggingRoundTripper) RoundTrip(r *http.Request) (*http.Response, error) {
	// Info log
	t.logger.Info(r.URL.String(), "Requesting...")

	// The http.RoundTripper contract forbids mutating the caller's request,
	// so clone it before adding the no-cache headers.
	r = r.Clone(r.Context())
	SetNoCacheHeaders(r.Header)

	// Response error log
	res, err := t.next.RoundTrip(r)
	if err != nil {
		t.logger.Error(r.URL.String(), err.Error())
	}
	return res, err
}
