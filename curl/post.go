package curl

import (
	"encoding/json"
	"errors"

	"github.com/astaxie/beego/httplib"
)

// Post POST请求
func (u Curl) Post(url string, params map[string]string) *Curl {
	return u.request(url, params, typePost)
}

// BodyString 设置Body
func (u *Curl) BodyString(body string) (*httplib.BeegoHTTPRequest, error) {
	if !u.used {
		return nil, errors.New("please initialize first")
	}
	u.log.Debug("Body: %v", body)
	return u.BeegoHTTPRequest.Body(body), nil
}

// BodyBytes 设置Body
func (u *Curl) BodyBytes(body []byte) (*httplib.BeegoHTTPRequest, error) {
	if !u.used {
		return nil, errors.New("please initialize first")
	}
	u.log.Debug("Body: %v", string(body))
	return u.BeegoHTTPRequest.Body(body), nil
}

// BodyJSON 设置Body
func (u *Curl) BodyJSON(body interface{}) (*httplib.BeegoHTTPRequest, error) {
	if !u.used {
		return nil, errors.New("please initialize first")
	}
	if body != nil {
		byt, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		u.log.Debug("Body: %v", string(byt))
	}
	return u.BeegoHTTPRequest.JSONBody(body)
}
