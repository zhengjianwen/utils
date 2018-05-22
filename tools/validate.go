package tools

import "regexp"

//必须输数字
func IsNumber(data string) bool {
	b,_ := regexp.MatchString("^[0-9]+$", data)
	return b
}
// 必须输字母
func IsLetter(data string) bool {
	b,_ := regexp.MatchString("^[a-zA-Z]+$", data)
	return b
}

func NotHasSpance (data string) bool{
	b,_ := regexp.MatchString("^[-$#@*.a-zA-Z0-9_,\u4e00-\u9fa5]+$", data)
	return b
}

// 以字母开头数字字母下划线结尾
func IsLetterNumber(data string) bool {
	b,_ := regexp.MatchString("^[a-zA-Z]+[_0-9a-zA-Z]+$", data)
	return b
}
