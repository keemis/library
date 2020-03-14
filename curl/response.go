package curl

import (
	"compress/gzip"
	"errors"
	"io/ioutil"
	"net/http"
)

// ResponseBytes 返回结果
func (u *Curl) ResponseBytes() ([]byte, error) {
	if !u.used {
		return nil, errors.New("please initialize first")
	}
	_, byt, err := u.Response()
	return byt, err
}

// ResponseString 返回结果
func (u *Curl) ResponseString() (string, error) {
	if !u.used {
		return "", errors.New("please initialize first")
	}
	_, byt, err := u.Response()
	if err != nil {
		return "", err
	}
	return string(byt), err
}

// Response 返回结果
func (u *Curl) Response() (*http.Response, []byte, error) {
	resp, err := u.BeegoHTTPRequest.Response()
	if err != nil {
		u.log.Warn("curl <<< Response http failed, err: %v", err)
		return nil, nil, err
	}
	byt, err := u.decodeResponse(resp)
	if err != nil {
		u.log.Warn("curl <<< Response decode failed, err: %v", err)
		return nil, nil, err
	}
	u.log.Debug("curl <<< Response success, code: %v, body: %v ", resp.StatusCode, string(byt))
	return resp, byt, err
}

// decodeResponse 解析Response
func (u *Curl) decodeResponse(resp *http.Response) ([]byte, error) {
	if resp.Body == nil {
		return nil, nil
	}
	defer resp.Body.Close()
	if resp.Header.Get("Content-Encoding") == "gzip" {
		reader, err := gzip.NewReader(resp.Body)
		if err != nil {
			return nil, err
		}
		byt, err := ioutil.ReadAll(reader)
		return byt, err
	}
	byt, err := ioutil.ReadAll(resp.Body)
	return byt, err
}
