package main

import (
	"net/http"
	"encoding/json"
	"bytes"
	"github.com/zhengjianwen/utils/log"
	"github.com/PuerkitoBio/goquery"
	"golang.org/x/text/encoding"
)


func main()  {
	url := "http://www.tianshuba.com/reading/61/61904/18541659.html"
	data := getUrl(url,nil)
	log.Info(data)

}

func getUrl(url string,data interface{}) string {
	client := &http.Client{}
	tmp, err := json.Marshal(data)
	if err != nil {
		panic("请求参数不正确")
	}
	req, err := http.NewRequest("GET", url, bytes.NewBuffer(tmp))
	req.Header.Add("Content-Type", "application/json;charset=utf-8;")
	req.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8")
	req.Header.Add("Cache-Control", "max-age=0")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Cookie", "Hm_lvt_4070987fd5c0d65cb4c4a00de66865d5=1561689127; UM_distinctid=16b9becf9173a0-00907f9294ffab-38395f03-1aeaa0-16b9becf9186b4; CNZZDATA1260819367=1825031779-1561684797-null%7C1561684797; Hm_lpvt_4070987fd5c0d65cb4c4a00de66865d5=1561689144; PHPSESSID=c5837e20eebcc603b215360305aa75fb; BLR=61%2361904%2318541659%23%u4E11%u5973%u79CD%u7530%uFF1A%u5C71%u91CC%u6C49%u5BA0%u59BB%u65E0%u5EA6%23%u7AE0%u8282%u76EE%u5F55%201%u7B2C1%u7AE0%u4E11%u5AB3")
	req.Header.Add("f-None-Match","W/'81ca33b67ebed21:0'")
	req.Header.Add("Upgrade-Insecure-Requests","1")
	req.Header.Add("User-Agent","Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/72.0.3626.121 Safari/537.36")
	req.Header.Add("Accept-Language", "zh-CN,zh;q=0.9,zh-TW;q=0.8,en;q=0.7")
	if err != nil {
		log.Errorf("[HttpGet] 创建请求错误，%v", err)
		panic("请求参数错误")
	}
	res, err := client.Do(req)
	if err != nil {
		log.Debugf("[HttpGet]请求错误: %v", err)
		panic("请求错误")
	}
	defer res.Body.Close()
	doc, err := goquery.NewDocumentFromReader(res.Body)
	txt := getText(doc)

	return txt

}

func getText(doc *goquery.Document) string {
	log.Println(doc.Text())
	text := ""
	doc.Find("#BookText div").Each(func(i int, s *goquery.Selection) {
		// name
		text += s.Text()
	})
	return text
}