package slice

import (
	"strconv"
	"strings"
)

// 数组转字符串
func JoinInt(elems []int, sep string) string {
	var tmp = make([]string, len(elems))
	for k, v := range elems {
		tmp[k] = strconv.Itoa(v)
	}
	return JoinString(tmp, sep)
}

// 数组转字符串
func JoinInt64(elems []int64, sep string) string {
	var tmp = make([]string, len(elems))
	for k, v := range elems {
		tmp[k] = strconv.FormatInt(v, 10)
	}
	return JoinString(tmp, sep)
}

// 数组转字符串
func JoinString(elems []string, sep string) string {
	return strings.Join(elems, sep)
}
