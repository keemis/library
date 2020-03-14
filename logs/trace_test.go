package logs

import (
	"testing"
)

func TestGenerator(t *testing.T) {
	id := Generator()
	t.Logf("Generator uint64: %v ", id)
}

func TestNewTraceID(t *testing.T) {
	id := NewTraceID()
	t.Logf("TraceID: %v", id)
}
