package gohttpclient

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

/**
  执行PostForm请求操作;
*/
func (this *PostForm) Invoke() (string, *http.Response) {
	var (
		request *http.Request
		err     error
		q       url.Values
	)
	// 追加请求的get参数 - 循环添加
	if this.paramData != nil {
		q = url.Values{}
		// 循环添加请求参数
		for key, value := range this.paramData {
			if this.Debug {
				fmt.Printf("循环添加请求参数 key:%+v - value:%+v \n", key, value)
			}
			for _, v := range value {
				q.Add(key, v)
			}
		}
	}
	// 构造请求的结构数据
	request, err = http.NewRequest("POST", this.Url, strings.NewReader(q.Encode()))
	if err != nil {
		fmt.Printf("err:%+v\n", err)
	}

	// 循环添加头部参数
	for key, value := range this.header {
		if this.Debug {
			fmt.Printf("循环添加头部参数 key:%+v - value:%+v \n", key, value)
		}
		request.Header.Add(key, value)
	}
	// 表单提交也有的头部
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	var (
		response    *http.Response
		body_byte   []byte
		body_string string
	)
	// 执行发送请求的操作
	response, err = (&Httpclient{Debug: this.Debug}).Invoke().Do(request)
	if err != nil {
		fmt.Printf("err:%+v\n", err)
	}
	// 在函数退出的时候，关闭和服务器自己的链接 [如果不关闭的话，那么这个链接就长久占着不释放了]
	defer response.Body.Close()
	if err != nil {
		fmt.Printf("err:%+v\n", err)
		return "", response
	}
	// 获取请求返回的内容字符串
	body_byte, err = ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("err:%+v\n", err)
		return "", response
	}
	body_string = string(body_byte)
	// 返回请求结果 和整个请求对象数据（方便后续做判断分析）
	return body_string, response
}
