package beeControll

import (
	"testing"

	"github.com/astaxie/beego/context"
	"github.com/keemis/library/errs"
)

// 模拟创建一个HTTP请求
func createRequest() *BaseController {
	req := &BaseController{}
	req.Init(context.NewContext(), "x", "x", nil)
	req.Prepare()
	return req
}

func TestRun(t *testing.T) {
	req := createRequest()

	req.ApiSuccess(map[string]interface{}{
		"id": 33,
	})

	req.ApiError("error message", -1000)

	req.ApiErrorf("error message: %v", "php")

	req.ApiResult(WithCode(-1100), WithMsg("error"))

	err := errs.NewWithOption(errs.WithCode(-2211), errs.WithMsg("error"))
	req.ApiStruct(err)

}
