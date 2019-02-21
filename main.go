package main

import (
	"github.com/zhengjianwen/utils/wx"
	"time"
	"fmt"
)

func main()  {
	wx.Init()
	go wx.SyncToken()
	userid := "o9FYQ1OY3fLKPLUvV6RdnnU2un-A"
	//userid := "o9FYQ1C68VzP-abo5XEP09iRXgZw"
	t := time.NewTicker(time.Second * 5)
	i := 0
	for range t.C{
		i ++
		go wx.SendMsgToUser(userid,"text",fmt.Sprintf("hello 蛋蛋 - 第%d次",i))
	}

	select{}
}
