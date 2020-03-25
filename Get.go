package gohttpclient

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

/**
清空参数
*/
func (this *Get) CleanParamData() *Get {
	this.ParamData = nil
	return this
}

/**
发出get请求
*/
func (this *Get) Invoke() (string, *http.Response) {
	var (
		request *http.Request
		err     error
	)
	// 准备请求 ，生成句柄
	request, err = http.NewRequest("GET", this.Url, nil)
	if err != nil {
		fmt.Printf("err:%+v\n", err)
	}
	// 循环添加头部参数
	for key, value := range this.Header {
		if this.Debug {
			fmt.Printf("循环添加头部参数 key:%+v - value:%+v \n", key, value)
		}
		request.Header.Add(key, value)
	}
	// 追加请求的get参数 - 循环添加
	if this.ParamData != nil {
		q := request.URL.Query()
		// 循环添加请求参数
		for key, value := range this.ParamData {
			if this.Debug {
				fmt.Printf("循环添加请求参数 key:%+v - value:%+v \n", key, value)
			}
			for _, v := range value {
				q.Add(key, v)
			}
		}
		request.URL.RawQuery = q.Encode()
	}

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
