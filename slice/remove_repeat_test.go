package slice

import (
	"testing"
)

func TestRemoveRepeatInt(t *testing.T) {
	elems := []int{9, 2, 1, 9, 2, 7, 3, 3, 3, 4, 5, 6, 7, 6, 7, 7, 8, 1, 5}
	tmp := RemoveRepeatInt(elems)
	t.Logf("%v", tmp)
}

func TestRemoveRepeatInt64(t *testing.T) {
	elems := []int64{9, 2, 1, 9, 2, 7, 3, 3, 3, 4, 5, 6, 7, 6, 7, 7, 8, 1, 5}
	tmp := RemoveRepeatInt64(elems)
	t.Logf("%v", tmp)
}

func TestRemoveRepeatString(t *testing.T) {
	elems := []string{"9", "2", "1", "9", "2", "7", "3", "4", "5", "6", "7", "6", "8", "1", "5"}
	tmp := RemoveRepeatString(elems)
	t.Logf("%v", tmp)
}
