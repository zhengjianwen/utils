package wx

import (
	"fmt"
	"github.com/toolkits/net/httplib"
	"github.com/zhengjianwen/utils/log"
	"encoding/json"
)

type ArticlesData struct {
	Title  string
	Description  string
	Url  string
	Picurl  string
}

type SendMsgContent struct {
	Content  string `json:"content"`//text
	//MediaId  string  //image MEDIA_ID  voice
	//Articles []map[string]ArticlesData
}

type SendMsgData struct {
	Touser 	string 	`json:"touser"`
	Msgtype string `json:"msgtype"`// text  image  voice
	Text    map[string]interface{} `json:"text"`
	Image	map[string]interface{}
	Voice	map[string]interface{}
	Mpnews  SendMsgContent
	News	SendMsgContent
	Msgmenu	map[string]interface{}
}

type SendMsgRet struct {
	Errcode  	int64
	Errmsg 		string

}

func SendMsgToUser(userOpenID ,msgtype,msg string) error {
	data := SendMsgData{
		Touser:userOpenID,
		Msgtype:msgtype,
	}
	if len(msg) <= 0{
		return fmt.Errorf("发送消息不能为空")
	}
	m := make(map[string]interface{})
	switch msgtype {
	case "text":
		m["content"] = msg
		data.Text = m
		log.Debugf("%s - %s",msgtype,msg)
	case "image":
		m["media_id"] = msg
		data.Image = m
	case "voice":
		m["media_id"] = msg
		data.Image = m
	case "video":
	case "music":
	case "news":
	case "mpnews":
		m["media_id"] = msg
		data.Image = m
	case "msgmenu":
		m["head_content"] = "您对本次服务是否满意呢?"
		list := make([]map[string]string,0)
		c1,c2 := make(map[string]string),make(map[string]string)
		c1["id"] = "101"
		c1["content"] = "满意"
		c2["id"] = "102"
		c2["content"] = "满意"
		list = append(list,c1)
		list = append(list,c2)
		m["list"] = list
		m["tail_content"] = "欢迎再次光临"
		data.Msgmenu = m
	default:
		return fmt.Errorf("消息类型不支持")
	}
	url := fmt.Sprintf(SendMsgUrl,G.Token)
	r,err := httplib.PostJSON(url,data)
	if err != nil{
		log.Errorf("[SendMsgToUser] %v",err)
		return fmt.Errorf("请求失败")
	}
	var ret SendMsgRet
	if err := json.Unmarshal(r,&ret);err != nil{
		log.Errorf("[SendMsgToUser][Unmarshal] %v",err)
		return fmt.Errorf("解析失败")
	}
	if ret.Errcode != 0{
		switch ret.Errcode {
		case -1:
			log.Debugf("[SendMsgToUser] 系统繁忙 - %s",ret.Errmsg)
		case 45047:
			log.Debugf("[SendMsgToUser] 单用户发送数量太多，用户未回复 - %s",ret.Errmsg)
		case 40001:
			log.Debugf("[SendMsgToUser] Token 过期了 - %s",ret.Errmsg)
		default:
			log.Debugf("[SendMsgToUser] 错误消息：%v",ret)
		}
		

		return fmt.Errorf("发送失败")
	}
	return nil
}
