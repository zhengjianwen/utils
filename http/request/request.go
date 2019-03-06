package request

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"bytes"
	"github.com/zhengjianwen/utils/log"
	"encoding/json"
)

func HttpGetJson(url string, data, resp interface{}) (error) {
	client := &http.Client{}
	tmp, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("请求参数不正确")
	}
	req, err := http.NewRequest("GET", url, bytes.NewBuffer(tmp))
	req.Header.Add("Content-Type", "application/json;charset=utf-8;")
	req.Header.Add("Accept-Language", "zh")
	if err != nil {
		log.Errorf("[HttpGet] 创建请求错误，%v", err)
		return fmt.Errorf("请求参数错误")
	}
	res, err := client.Do(req)
	if err != nil {
		log.Debugf("[HttpGet]请求错误: %v", err)
		return fmt.Errorf("请求错误")
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Errorf("[HttpGet] %v  ret:%s", err, string(body))
		return fmt.Errorf("获取数据错误")
	}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		log.Errorf("[HttpGet] %v  ret:%s", err, string(body))
		return fmt.Errorf("获取数据错误")
	}

	return nil
}

func HttpPostJson(url string, data, resp interface{}) (error) {
	client := &http.Client{}
	tmp, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("请求参数不正确")
	}
	req, err := http.NewRequest("GET", url, bytes.NewBuffer(tmp))
	req.Header.Add("Content-Type", "application/json;charset=utf-8;")
	req.Header.Add("Accept-Language", "zh")
	if err != nil {
		log.Errorf("[HttpGet] 创建请求错误，%v", err)
		return fmt.Errorf("请求参数错误")
	}
	res, err := client.Do(req)
	if err != nil {
		log.Debugf("[HttpGet]请求错误: %v", err)
		return fmt.Errorf("请求错误")
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Errorf("[HttpGet] %v  ret:%s", err, string(body))
		return fmt.Errorf("获取数据错误")
	}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		log.Errorf("[HttpGet] %v  ret:%s", err, string(body))
		return fmt.Errorf("获取数据错误")
	}

	return nil
}
