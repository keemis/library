package timer

import (
	"testing"
	"time"
)

func TestRun(t *testing.T) {
	// FormatCN
	str := FormatCN(time.Now())
	t.Logf("time: %v", str)

	// ParseCN
	ts, err := ParseCN("2020-03-16 15:53:17")
	t.Logf("time: %v, error: %v", ts, err)
}
