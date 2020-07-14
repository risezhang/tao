package slice

type Int64Slice []int64

func (s Int64Slice) InterfaceSlice() []interface{} {
	var ss []interface{}
	for _, v := range s {
		ss = append(ss, v)
	}
	return ss
}

func (s Int64Slice) Has(item int64) bool {
	if s == nil {
		return false
	}
	for _, v := range s {
		if v == item {
			return true
		}
	}

	return false
}

func UniqueInt64s(ss []int64) []int64 {
	uniq := make(map[int64]bool)
	for _, v := range ss {
		uniq[v] = true
	}

	var res []int64
	for key := range uniq {
		res = append(res, key)
	}

	return res
}
