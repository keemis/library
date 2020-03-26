package slice

import (
	"testing"
)

func TestJoinInt(t *testing.T) {
	elems := []int{9, 2, 1, 9, 2, 7, 3, 3, 3, 4, 5, 6, 7, 6, 7, 7, 8, 1, 5}
	tmp := JoinInt(elems, ",")
	t.Logf("%v", tmp)
}

func TestJoinInt64(t *testing.T) {
	elems := []int64{9, 2, 1, 9, 2, 7, 3, 3, 3, 4, 5, 6, 7, 6, 7, 7, 8, 1, 5}
	tmp := JoinInt64(elems, ",")
	t.Logf("%v", tmp)
}

func TestJoinString(t *testing.T) {
	elems := []string{"9", "2", "1", "9", "2", "7", "3", "4", "5", "6", "7", "6", "8", "1", "5"}
	tmp := JoinString(elems, ",")
	t.Logf("%v", tmp)
}
