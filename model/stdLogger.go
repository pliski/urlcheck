package model

import (
	"fmt"
	"io"
	"os"
	"time"
)

type StdLogger struct {
	out  io.Writer
	info io.Writer
	err  io.Writer
}

func NewStdLogger() *StdLogger {
	return &StdLogger{
		out:  os.Stdout,
		info: os.Stdout,
		err:  os.Stderr,
	}
}

func (logger StdLogger) Error(url string, msg string) {
	if msg != "" {
		fmt.Fprintf(logger.err, "[%s] ERROR %s: %s \n", now(), url, msg)
	}
}

func (logger StdLogger) Info(url string, msg string) {
	if msg != "" {
		fmt.Fprintf(logger.info, "[%s] Info %s: %s \n", now(), url, msg)
	}
}

func (logger StdLogger) Out(url string, msg string) {
	if msg != "" {
		fmt.Fprintf(logger.out, "%s: %s \n", url, msg)
	}
}

func now() string {
	return time.Now().Format(time.ANSIC)
}
