package model

import "time"

type LoggerURl interface {
	StdLogger | TeaLogger
	Error(url string, msg string)
	Info(url string, msg string)
	Out(url string, msg string)
}

func now() string {
	return time.Now().Format(time.ANSIC)
}
