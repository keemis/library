package copier

import (
	"testing"
	"time"

	"github.com/jinzhu/copier"
)

//
// https://github.com/jinzhu/copier
//

type User struct {
	ID         int
	Name       string
	Age        int
	CreateTime time.Time
}

type Employee struct {
	ID     int
	Name   string
	Age    int
	Role   string
	Gender int
}

func TestRun(t *testing.T) {
	user := User{
		ID:   11,
		Name: "simon",
		Age:  23,
	}
	employee := Employee{}
	_ = copier.Copy(&employee, &user)
	t.Logf("user: %+v", user)
	t.Logf("employee: %+v", employee)
}
