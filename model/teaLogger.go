package model

import (
	"fmt"
)

// type errType string

type TeaLogger struct {
	errList *[]string
}

func NewTeaLogger(errList *[]string) *TeaLogger {
	return &TeaLogger{
		errList: errList,
	}
}

func (logger TeaLogger) appendMsg(msg string) {
	*logger.errList = append(*logger.errList, msg)
	fmt.Println(msg)
}

func (logger TeaLogger) Error(url string, msg string) {
	if msg != "" {
		logger.appendMsg(fmt.Sprintf("[%s] ERROR %s: %s \n", now(), url, msg))
	}
}

func (logger TeaLogger) Info(url string, msg string) {
	if msg != "" {
		logger.appendMsg(fmt.Sprintf("[%s] Info %s: %s \n", now(), url, msg))
	}
}

func (logger TeaLogger) Out(url string, msg string) {
	if msg != "" {
		logger.appendMsg(fmt.Sprintf("%s: %s \n", url, msg))
	}
}
