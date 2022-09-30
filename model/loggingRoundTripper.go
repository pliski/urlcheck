package model

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

type loggingRoundTripper struct {
	next http.RoundTripper
	info io.Writer
	err  io.Writer
}

func NewLoggingRoundTripper(next http.RoundTripper, info io.Writer, err io.Writer) *loggingRoundTripper {
	return &loggingRoundTripper{
		next: http.DefaultTransport,
		info: os.Stdout,
		err:  os.Stderr,
	}
}

func (logger loggingRoundTripper) RoundTrip(r *http.Request) (*http.Response, error) {
	// Info log
	fmt.Fprintf(logger.info, "[%s] %s %s\n", time.Now().Format(time.ANSIC), r.Method, r.URL.String())

	// Response error log
	res, err := logger.next.RoundTrip(r)
	if err != nil {
		fmt.Fprintf(logger.err, "[%s] ERROR %s %s %s \n", time.Now().Format(time.ANSIC), err.Error(), r.Method, r.URL.String())
	}
	return res, err
}
