package mysql

import (
	"testing"
)

func TestDB(t *testing.T) {
	Init()
	o := New()
	db := o.DB("test")
	t.Logf("db: %v", db)
}

func TestQuery(t *testing.T) {
	Init()
	o := New()
	db := o.DB("test")
	cnt := 0
	db.Table("school").Count(&cnt)
	t.Logf("cnt: %v", cnt)
}
