package gohttpclient
import (
        "net/http"
)

type PostInterface interface {
    /**
    执行Post请求操作    */
    Invoke()(string, *http.Response)
}

/**
发起Post请求，包含了连接池机制;
*/
type Post struct
{
    /*是否打开信息调试，默认关闭*/
    Debug bool
    /*发送头部的map数组(一维)*/
    header map[string]string
    /*参数*/
    paramData string
    /*要请求的目标网址*/
    Url string
}



    //get -
    func (this *Post) GetDebug() bool{
        return this.Debug;
    }
    //如果是bool类型的属性，那么多了一个is前缀的函数
    func (this *Post) IsDebug() bool{
        return this.Debug;
    }

    //set -
    func (this *Post) SetDebug(Debug bool) *Post{
        this.Debug = Debug;
        return this
    }

    //get -
    func (this *Post) Getheader() map[string]string{
        return this.header;
    }

    //set -
    func (this *Post) Setheader(header map[string]string) *Post{
        this.header = header;
        return this
    }
    //数组格式，在单个变量设置的时候.在go语言中，map[string]string (一维map)  和 map[string][]string （二维map）
    //区分大小写索引
    func (this *Post) Setheader_raw(key,value string) *Post {
        // 如果头部参数是第一次赋值，那么顺便初始化下
        if this.header == nil {
            this.header = make(map[string]string)
        }
        this.header[key] = value
        return this
    }

    //get -
    func (this *Post) GetparamData() string{
        return this.paramData;
    }

    //set -
    func (this *Post) SetparamData(paramData string) *Post{
        this.paramData = paramData;
        return this
    }

    //get -
    func (this *Post) GetUrl() string{
        return this.Url;
    }

    //set -
    func (this *Post) SetUrl(Url string) *Post{
        this.Url = Url;
        return this
    }





