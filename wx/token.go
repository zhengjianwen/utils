package wx

import (
	"fmt"
	"github.com/toolkits/net/httplib"
	"github.com/zhengjianwen/utils/log"
	"encoding/json"
)

type TokenRet struct {
	AccessToken string `json:"access_token"`
	ExpiresIn	int64 	`json:"expires_in"`
}

func GetToken() string {
	if G != nil{
		url := fmt.Sprintf(TokenGetUrl,G.Appid,G.AppScret)
		r := httplib.Get(url)
		resp, err := r.Bytes()
		if err != nil{
			log.Errorf("[GET][TOKEN] 请求失败 %v",err)
			return ""
		}
		//log.Debugf("[GET][TOKEN] BODY： %v",string(resp))
		var data TokenRet
		err = json.Unmarshal(resp,&data)
		if err != nil{
			log.Fatalf("[GET][Unmarshal] %v",err)
		}
		if &data != nil{
			log.Debugf("[GET][TOKEN] 获取TOKEN:%s",data.AccessToken)
			return data.AccessToken
		}
	}
	log.Fatalf("[SYS] 系统未初始化数据")
	return ""
}
