package curl

// Get GET请求
func (u Curl) Get(url string, params map[string]string) *Curl {
	return u.request(url, params, typeGet)
}
