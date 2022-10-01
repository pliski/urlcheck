package action

import (
	"context"
	"errors"
	"net/http"
	"time"
	"urlcheck/model"
)

type ClientModel[L model.Logger] struct {
	logger *L
	client *http.Client
}

func IsStatusOK(urlRequest string, timeout uint) bool {
	// Client
	var clientModel *ClientModel[model.StdLogger] = NewClient()

	return getStatus(urlRequest, time.Duration(timeout)*time.Second, clientModel) < http.StatusInternalServerError
}

func NewClient() *ClientModel[model.StdLogger] {
	// shared Logger
	var logger *model.StdLogger = model.NewStdLogger()

	// Client
	client := &http.Client{
		Transport: model.NewLoggingRoundTripper(
			http.DefaultTransport,
			logger,
		),
	}

	cm := ClientModel[model.StdLogger]{
		logger,
		client,
	}

	return &cm
}

func NewClientTea(errList *[]string) *ClientModel[model.TeaLogger] {
	// shared Logger
	logger := model.NewTeaLogger(errList)

	// Client
	client := &http.Client{
		Transport: http.DefaultTransport,
	}

	return &ClientModel[model.TeaLogger]{
		logger,
		client,
	}
}

func getStatus[L model.Logger](urlRequest string, timeout time.Duration, clientModel *ClientModel[L]) int {

	var cm *ClientModel[L] = clientModel
	var logger L = *cm.logger

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
	status, statusCode := sendRequest(clientModel.client, req, logger)
	logger.Info(urlRequest, status)
	return statusCode

}

func sendRequest[L model.Logger](c *http.Client, req *http.Request, logger L) (string, int) {
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
