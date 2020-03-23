package mysql

import (
	"testing"

	"github.com/keemis/library/logs"
)

func TestQuery(t *testing.T) {
	log := getLogger()
	Init( /* db list conf */ )
	o := New(WithLogger(log))
	db := o.DB( /* db name */ )
	cnt := 0
	db.Table("school").Count(&cnt)
	t.Logf("cnt: %v", cnt)
}

func getLogger() logs.Logger {
	logs.Init(nil)
	log := logs.New(logs.WithTraceID("158493062888748290"))
	return log
}
