package slice

import (
	"testing"
)

func TestRemoveRepeatInt(t *testing.T) {
	ids := []int{9, 2, 1, 9, 2, 7, 3, 3, 3, 4, 5, 6, 7, 6, 7, 7, 8, 1, 5}
	tmp := RemoveRepeatInt(ids)
	t.Logf("%v", tmp)
}

func TestRemoveRepeatInt64(t *testing.T) {
	ids := []int64{9, 2, 1, 9, 2, 7, 3, 3, 3, 4, 5, 6, 7, 6, 7, 7, 8, 1, 5}
	tmp := RemoveRepeatInt64(ids)
	t.Logf("%v", tmp)
}

func TestRemoveRepeatString(t *testing.T) {
	ids := []string{"9", "2", "1", "9", "2", "7", "3", "4", "5", "6", "7", "6", "8", "1", "5"}
	tmp := RemoveRepeatString(ids)
	t.Logf("%v", tmp)
}
