package go_log

import (
	"testing"
)

func TestInfo(t *testing.T) {
	Info("Hello", "msg", "Hello World.", "val", 1)
}
