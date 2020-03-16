package cache

import (
	"testing"
)

func TestRun(t *testing.T) {
	// init
	Init(5 * 1024 * 1024)

	// set
	Client().Set([]byte("k"), []byte("v"), 0)
	v, err := Client().Get([]byte("k"))
	t.Logf("value: %v, err: %v", string(v), err)

	// del
	affected := Client().Del([]byte("k"))
	t.Logf("affected: %v", affected)

	// get
	v, err = Client().Get([]byte("k"))
	t.Logf("value: %v, err: %v", string(v), err)
}
