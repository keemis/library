package mysql

import (
	"testing"
)

func TestInit(t *testing.T) {
	Init()
	t.Logf("connects: %v", connects)
}
