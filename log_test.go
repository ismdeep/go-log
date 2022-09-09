package log

import (
	"errors"
	"testing"
)

func TestInfo(t *testing.T) {
	Info("Hello Info", "msg", "Hello World.", "val", 1)
	Debug("Hello Debug", "msg", "Hello World.", "val", 1)
	Warn("Hello Warn", "msg", "Hello World.", "val", 1)
	Error("Hello Error", String("msg", "Hello World."), Any("val", 1), FieldErr(errors.New("error msg")))
}
