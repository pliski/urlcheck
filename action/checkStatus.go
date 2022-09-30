package action

import (
	"context"
	"errors"
	"net/http"
	"time"
	"urlcheck/model"
)

func IsStatusOK(urlRequest string, timeout uint) bool {
	return getStatus(urlRequest, time.Duration(timeout)*time.Second) < http.StatusInternalServerError
}

func getStatus(urlRequest string, timeout time.Duration) int {

	// shared Logger
	logger := model.NewStdLogger()

	// Client
	client := &http.Client{
		Transport: model.NewLoggingRoundTripper(
			http.DefaultTransport,
			logger,
		),
	}

	// Request
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	req, err := http.NewRequest(http.MethodHead, urlRequest, nil)
	if err != nil {
		logger.Error(urlRequest, err.Error())
		return 999
	}
	req = req.WithContext(ctx)

	// Status
	status, statusCode := sendRequest(client, req, logger)
	logger.Info(urlRequest, status)
	return statusCode

}

func sendRequest(c *http.Client, req *http.Request, logger *model.StdLogger) (string, int) {
	res, err := c.Do(req)
	if err != nil {
		// The error will already be logged by loggingRoundTripper
		if errors.Is(err, context.DeadlineExceeded) {
			return "Timeout", 999
		}
		return "Error", 999
	}
	defer res.Body.Close()
	return res.Status, res.StatusCode
}
