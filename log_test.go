package log

import (
	"testing"
)

func TestInfo(t *testing.T) {
	Info("Hello Info", "msg", "Hello World.", "val", 1)
	Debug("Hello Debug", "msg", "Hello World.", "val", 1)
	Warn("Hello Warn", "msg", "Hello World.", "val", 1)
	Error("Hello Error", "msg", "Hello World.", "val", 1)
	Panic("Hello Panic", "msg", "Hello World.", "val", 1)
	Fatal("Hello Fatal", "msg", "Hello World.", "val", 1)
}
