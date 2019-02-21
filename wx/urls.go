package wx

var (
	TokenGetUrl = "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s"
	SendMsgUrl = "https://api.weixin.qq.com/cgi-bin/message/custom/send?access_token=%s"
)
