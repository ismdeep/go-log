package log

import (
	"testing"
)

func TestWithContext(t *testing.T) {
	ctx := NewTraceContext("07a313fc-9ee9-11ed-b3d1-33fcf0cd96d8")
	WithContext(ctx).Info("hello")
}

func TestNew(t *testing.T) {
	logger, err := New("console://[stdout]?level=debug&time_encoder=rfc3339&caller_encoder=short&trace_level=error")
	if err != nil {
		panic(err)
	}

	ctx := NewTraceContext("07a313fc-9ee9-11ed-b3d1-33fcf0cd96d8")
	logger.WithContext(ctx).Info("hello")
}
