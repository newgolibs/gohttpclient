package gohttpclient
import (
        "net/http"
)

type GetInterface interface {
    /**
    清除上一次的参数拼接    */
    CleanParamData()Get
    /**
    发起Get请求    */
    Invoke()(string, *http.Response)
}

/**
发起get请求，包含了连接池机制;
*/
type Get struct
{
    /*是否打开信息调试，默认关闭*/
    Debug bool
    /*发送头部的map数组(一维)*/
    Header map[string]string
    /*参数*/
    ParamData map[string][]string
    /*要请求的目标网址*/
    Url string
}



    //get -
    func (this *Get) GetDebug() bool{
        return this.Debug;
    }
    //如果是bool类型的属性，那么多了一个is前缀的函数
    func (this *Get) IsDebug() bool{
        return this.Debug;
    }

    //set -
    func (this *Get) SetDebug(Debug bool) *Get{
        this.Debug = Debug;
        return this
    }

    //get -
    func (this *Get) GetHeader() map[string]string{
        return this.Header;
    }

    //set -
    func (this *Get) SetHeader(Header map[string]string) *Get{
        this.Header = Header;
        return this
    }
    //数组格式，在单个变量设置的时候.在go语言中，map[string]string (一维map)  和 map[string][]string （二维map）
    //区分大小写索引
    func (this *Get) SetHeader_raw(key,value string) *Get {
        // 如果头部参数是第一次赋值，那么顺便初始化下
        if this.Header == nil {
            this.Header = make(map[string]string)
        }
        this.Header[key] = value
        return this
    }

    //get -
    func (this *Get) GetParamData() map[string][]string{
        return this.ParamData;
    }

    //set -
    func (this *Get) SetParamData(ParamData map[string][]string) *Get{
        this.ParamData = ParamData;
        return this
    }
    //数组格式，在单个变量设置的时候.在go语言中，map[string]string (一维map)  和 map[string][]string （二维map）
    //区分大小写索引
    func (this *Get) SetParamData_raw(key,value string) *Get {
        // 如果头部参数是第一次赋值，那么顺便初始化下
        if this.ParamData == nil {
            this.ParamData = make(map[string][]string)
        }
        this.ParamData[key] = append(this.ParamData[key], value)
        return this
    }

    //get -
    func (this *Get) GetUrl() string{
        return this.Url;
    }

    //set -
    func (this *Get) SetUrl(Url string) *Get{
        this.Url = Url;
        return this
    }





