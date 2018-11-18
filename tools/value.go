package tools

import (
	"fmt"
	"net/http"
	"strconv"
	"github.com/zhengjianwen/utils/log"
	"github.com/gorilla/mux"
	"strings"
	"sort"
	"encoding/hex"
	"crypto/md5"
	"encoding/json"
)




func GetRequestValue(r *http.Request, key string) int64 {
	val := r.FormValue(key)
	if val == "" {
		var has bool
		val, has = mux.Vars(r)[key]
		if !has || val == "" {
			val = "0"
		}
	}
	i, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		log.Error("[GetRequestValue]获取", key, "错误", err, "i:", i)
	}
	return i
}

func GetRequestValueStr(r *http.Request, key string) string {
	val := r.FormValue(key)
	if val == "" {
		var has bool
		val, has = mux.Vars(r)[key]
		if !has || val == "" {
			val = ""
		}
	}
	return val
}

func GetRequestValueBool(r *http.Request, key string) bool {
	val := r.FormValue(key)
	if val == "" {
		val, _ = mux.Vars(r)[key]
	}
	switch val {
	case "true":
		return true
	case "false":
		return false
	}
	return false
}

func GetRequestInterface(r *http.Request, result *interface{}) *interface{} {
	if err := json.NewDecoder(r.Body).Decode(result); err != nil {
		log.Errorf("[GetRequestInterface] 解析错误: %v\n", err)
		return nil
	}
	return result
}

func StrToint(str string) int64 {
	if str != ""{
		new := strings.Replace(str," ","",-1)
		i,err := strconv.ParseInt(new,10,64)
		if err != nil{
			log.Println("[StrToint] 字符串转换数字失败,",err)
		}
		return i
	}
	return 0
}

func StrMd5(str string)  string {
	h := md5.New()
	base_str := str+"hairui.574601624"
	tmp := strings.Split(base_str,"")
	sort.Strings(tmp)
	str = strings.Join(tmp,"")
	h.Write([]byte(str)) //
	cipherStr := h.Sum(nil)
	md5_key:= fmt.Sprintf("%s", hex.EncodeToString(cipherStr)) // 输出加密结果
	return md5_key
}


