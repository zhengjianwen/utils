package tools

import (
	"strings"
	"strconv"
	"crypto/rand"
	"time"

	"github.com/kless/osutil/user/crypt/sha512_crypt"
)

// EncryptPassword 对密码加盐加密
func EncryptPassword(password string) string {
	hash := sha512_crypt.New()
	encrypt, _ := hash.Generate([]byte(password), []byte("$6$"+ "ssadwacz21231"))
	return encrypt
}


// 必须是时间戳
func StrtoTimestamp(data string) int64 {
	if data == ""{
		return 0
	}
	dlist := strings.Split(data,"")
	if len(dlist) != 10{
		return 0
	}
	ret,err:= strconv.ParseInt(data,10,64)
	if err != nil{
		return 0
	}
	return ret
}

// genToken 生成随机字符串
func GenToken(n int) string {
	const alphanum = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	var bytes = make([]byte, n)
	rand.Read(bytes)
	for i, b := range bytes {
		bytes[i] = alphanum[b%byte(len(alphanum))]
	}
	return string(bytes)
}

// 生成数字随机字符串
func GenNubToken(n int) string {
	const alphanum = "0123456789"
	var bytes = make([]byte, n)
	rand.Read(bytes)
	for i, b := range bytes {
		bytes[i] = alphanum[b%byte(len(alphanum))]
	}
	return string(bytes)
}

// 生成字母随机字符串
func GenLetterToken(n int) string {
	const alphanum = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	var bytes = make([]byte, n)
	rand.Read(bytes)
	for i, b := range bytes {
		bytes[i] = alphanum[b%byte(len(alphanum))]
	}
	return string(bytes)
}