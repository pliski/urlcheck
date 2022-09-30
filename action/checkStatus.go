package action

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"
	"urlcheck/model"
)

func IsStatusOK(urlRequest string, timeout uint) bool {
	return getStatus(urlRequest, time.Duration(timeout)*time.Second) < http.StatusInternalServerError
}

func getStatus(urlRequest string, timeout time.Duration) int {

	// Client
	client := &http.Client{
		Transport: model.NewLoggingRoundTripper(
			http.DefaultTransport,
			os.Stdout,
			os.Stderr,
		),
	}

	// Request
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	req, err := http.NewRequest(http.MethodHead, urlRequest, nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[%s] NewRequest Error url: %s method: %s\n", time.Now().Format(time.ANSIC), urlRequest, http.MethodHead)
		return 999
	}
	req = req.WithContext(ctx)

	// Status
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
