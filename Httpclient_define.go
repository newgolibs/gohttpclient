package gohttpclient

import (
    "context"
    "net"
    "net/http"
)

type HttpclientInterface interface {
	/**
	  生成新的配置    */
	Invoke() *http.Client
	/**
	  拦截请求的请求，可以在其中打印调试一些细节    */
	printLocalDial() func(ctx context.Context, network, addr string) (net.Conn, error)
}

/**
持久链接池配置，实现单例模式。所有请求方式公用的单一实例;
*/
type Httpclient struct {
	/*是否打开信息调试，默认关闭*/
	Debug bool
}

// get -
func (this *Httpclient) GetDebug() bool {
	return this.Debug
}
func (this *Httpclient) IsDebug() bool {
	return this.Debug
}

// set -
func (this *Httpclient) SetDebug(Debug bool) *Httpclient {
	this.Debug = Debug
    return this
}
