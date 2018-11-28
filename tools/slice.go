package tools

import "sort"

func SortSilceInt64(data []int64) []int64 {
	tmp := make([]int,0)
	for _,v := range data{
		tmp = append(tmp,int(v))
	}
	sort.Ints(tmp)
	ret := make([]int64,0)
	for _,v := range tmp{
		ret = append(ret,int64(v))
	}
	return ret
}

func SortSilceInt(data []int) []int {
	sort.Ints(data)
	return data
}

func SortStr(data []string) []string {
	sort.Strings(data)
	return data
}