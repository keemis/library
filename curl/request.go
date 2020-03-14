package curl

import (
	"bytes"
	"crypto/tls"
	"net/url"
	"sort"
	"strings"
	"time"

	"github.com/astaxie/beego/httplib"
)

// request 设置Request
func (u Curl) request(url string, params map[string]string, typ requestType) *Curl {
	req := u
	req.BeegoHTTPRequest.Retries(2)
	req.BeegoHTTPRequest.SetTimeout(2*time.Second, 3*time.Second)
	req.BeegoHTTPRequest.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: u.tlsSecure})
	req.BeegoHTTPRequest.Header("User-Agent", "curl/9.9.9")
	url = BuildURL(url, params)
	if typ == typeGet {
		req.used = true
		req.log.Debug("curl >>> Get: %v", url)
		req.BeegoHTTPRequest = httplib.Get(url)
	} else if typ == typePost {
		req.used = true
		req.log.Debug("curl >>> Post: %v", url)
		req.BeegoHTTPRequest = httplib.Post(url)
	}
	return &req
}

// BuildURL 编码URL
func BuildURL(url string, params map[string]string) string {
	queryStr := BuildQuery(params)
	if queryStr != "" {
		character := ""
		if !strings.Contains(url, "?") {
			character = "?"
		} else {
			character = "&"
		}
		url = url + character + queryStr
	}
	return url
}

// BuildQuery 编码QUERY参数
func BuildQuery(params map[string]string) string {
	paramBody := ""
	if params != nil && len(params) > 0 {
		var keys []string
		for k := range params {
			keys = append(keys, k)
		}
		sort.Strings(keys)
		var buf bytes.Buffer
		for _, k := range keys {
			v := params[k]
			buf.WriteString(url.QueryEscape(k))
			buf.WriteByte('=')
			buf.WriteString(url.QueryEscape(v))
			buf.WriteByte('&')
		}
		paramBody = buf.String()
		paramBody = paramBody[0 : len(paramBody)-1]
	}
	return paramBody
}
