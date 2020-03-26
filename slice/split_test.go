package slice

import (
	"testing"
)

func TestSplitInt(t *testing.T) {
	s := "9,2,1,9,2,7,3,3,3,4,5,6,7,6,7,7,8,1,5"
	tmp, err := SplitInt(s, ",")
	t.Logf("result: %v, error: %v", tmp, err)
}

func TestSplitInt64(t *testing.T) {
	s := "9,2,1,9,2,7,3,3,3,4,5,6,7,6,7,7,8,1,5"
	tmp, err := SplitInt64(s, ",")
	t.Logf("result: %v, error: %v", tmp, err)
}

func TestSplitString(t *testing.T) {
	s := "9,2,1,9,2,7,3,3,3,4,5,6,7,6,7,7,8,1,5"
	tmp := SplitString(s, ",")
	t.Logf("result: %v", tmp)
}
