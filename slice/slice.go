package slice

// 数组去重
func RemoveRepeatInt(ids []int) []int {
	var res []int
	for _, v := range ids {
		if len(res) == 0 {
			res = append(res, v)
		} else {
			for k, id := range res {
				if v == id {
					break
				}
				if k == len(res)-1 {
					res = append(res, v)
				}
			}
		}
	}
	return res
}

// 数组去重
func RemoveRepeatInt64(ids []int64) []int64 {
	var res []int64
	for _, v := range ids {
		if len(res) == 0 {
			res = append(res, v)
		} else {
			for k, id := range res {
				if v == id {
					break
				}
				if k == len(res)-1 {
					res = append(res, v)
				}
			}
		}
	}
	return res
}

// 数组去重
func RemoveRepeatString(ids []string) []string {
	var res []string
	for _, v := range ids {
		if len(res) == 0 {
			res = append(res, v)
		} else {
			for k, id := range res {
				if v == id {
					break
				}
				if k == len(res)-1 {
					res = append(res, v)
				}
			}
		}
	}
	return res
}
