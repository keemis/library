package slice

import (
	"strconv"
	"strings"
)

// 字符串转数组
func SplitInt(s string, sep string) ([]int, error) {
	s = strings.Trim(s, " ")
	if s == "" {
		return nil, nil
	}
	arr := SplitString(s, sep)
	var res = make([]int, len(arr))
	for k, v := range arr {
		elem, err := strconv.Atoi(v)
		if err != nil {
			return nil, err
		}
		res[k] = elem
	}
	return res, nil
}

// 字符串转数组
func SplitInt64(s string, sep string) ([]int64, error) {
	s = strings.Trim(s, " ")
	if s == "" {
		return nil, nil
	}
	arr := SplitString(s, sep)
	var res = make([]int64, len(arr))
	for k, v := range arr {
		elem, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return nil, err
		}
		res[k] = elem
	}
	return res, nil
}

// 字符串转数组
func SplitString(s string, sep string) []string {
	return strings.Split(s, sep)
}
