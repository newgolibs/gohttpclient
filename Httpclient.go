package gohttpclient

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"sync"
	"time"
)

// 生成客户端链接请求的 - 单例模式
var once_httpclient sync.Once
var httpclient_instance *http.Client

/**
  生成新的配置;
*/
func (this *Httpclient) Invoke() *http.Client {
	/**
	返回http客户端的链接句柄。保证每次只生成一个
	*/
	once_httpclient.Do(func() {
		if this.Debug {
			fmt.Printf("%+v\n", "新建链接")
		}
		httpclient_instance = &http.Client{
			Transport: &http.Transport{MaxIdleConns: 3, MaxIdleConnsPerHost: 3, DialContext: this.printLocalDial,},
			// 链接10秒默认超时
			Timeout: time.Duration(10) * time.Second,}
	})
	return httpclient_instance
}

/**
  拦截请求的请求，可以在其中打印调试一些细节;
*/
func (this *Httpclient) printLocalDial(ctx context.Context, network, addr string) (net.Conn, error) {
	dial := net.Dialer{
		Timeout: 30 * time.Second,
		// 保存10分钟的长久链接
		KeepAlive: 10 * 60 * time.Second,
	}

	conn, err := dial.Dial(network, addr)
	if err != nil {
		return conn, err
	}
	if this.Debug {
		fmt.Println("connect done, use", conn.LocalAddr().String())
	}

	return conn, err
}
