package action

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func IsStatusOK(urlRequest string, timeout uint) bool {
	return GetStatus(urlRequest, time.Duration(timeout)*time.Second) < http.StatusInternalServerError
}

func GetStatus(urlRequest string, timeout time.Duration) int {

	// Request Timeout
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	client := &http.Client{
		Transport: &loggingRoundTripper{
			next: http.DefaultTransport,
			info: os.Stdout,
			err:  os.Stderr,
		},
		// Timeout: time.Duration(1 * time.Second), // Client total deadline (includes initializations, retries, response times, etc)
	}

	req, err := http.NewRequest(http.MethodHead, urlRequest, nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[%s] NewRequest Error url: %s method: %s\n", time.Now().Format(time.ANSIC), urlRequest, http.MethodHead)
		return 999
	}
	req = req.WithContext(ctx)

	status, statusCode := nowDo(client, req)
	fmt.Fprintf(os.Stdout, "%s: %s\n", urlRequest, status)
	return statusCode

}

func nowDo(c *http.Client, req *http.Request) (string, int) {
	res, err := c.Do(req)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			return "Timeout", 999
		}
		return "Error", 999
	}
	defer res.Body.Close()
	return res.Status, res.StatusCode
}

type loggingRoundTripper struct {
	next http.RoundTripper
	info io.Writer
	err  io.Writer
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
