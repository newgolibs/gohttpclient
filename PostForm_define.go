package gohttpclient
import (
        "net/http"
)

type PostFormInterface interface {
    /**
    执行PostForm请求操作    */
    Invoke()(string, *http.Response)
}

/**
发起表单类型的请求方式;
*/
type PostForm struct
{
    /*是否打开信息调试，默认关闭*/
    Debug bool
    /*发送头部的map数组(一维)*/
    header map[string]string
    /*参数*/
    paramData map[string][]string
    /*要请求的目标网址*/
    Url string
}



    //get -
    func (this *PostForm) GetDebug() bool{
        return this.Debug;
    }
    //如果是bool类型的属性，那么多了一个is前缀的函数
    func (this *PostForm) IsDebug() bool{
        return this.Debug;
    }

    //set -
    func (this *PostForm) SetDebug(Debug bool) *PostForm{
        this.Debug = Debug;
        return this
    }

    //get -
    func (this *PostForm) Getheader() map[string]string{
        return this.header;
    }

    //set -
    func (this *PostForm) Setheader(header map[string]string) *PostForm{
        this.header = header;
        return this
    }
    //数组格式，在单个变量设置的时候.在go语言中，map[string]string (一维map)  和 map[string][]string （二维map）
    //区分大小写索引
    func (this *PostForm) Setheader_raw(key,value string) *PostForm {
        // 如果头部参数是第一次赋值，那么顺便初始化下
        if this.header == nil {
            this.header = make(map[string]string)
        }
        this.header[key] = value
        return this
    }

    //get -
    func (this *PostForm) GetparamData() map[string][]string{
        return this.paramData;
    }

    //set -
    func (this *PostForm) SetparamData(paramData map[string][]string) *PostForm{
        this.paramData = paramData;
        return this
    }
    //数组格式，在单个变量设置的时候.在go语言中，map[string]string (一维map)  和 map[string][]string （二维map）
    //区分大小写索引
    func (this *PostForm) SetparamData_raw(key,value string) *PostForm {
        // 如果头部参数是第一次赋值，那么顺便初始化下
        if this.paramData == nil {
            this.paramData = make(map[string][]string)
        }
        this.paramData[key] = append(this.paramData[key], value)
        return this
    }

    //get -
    func (this *PostForm) GetUrl() string{
        return this.Url;
    }

    //set -
    func (this *PostForm) SetUrl(Url string) *PostForm{
        this.Url = Url;
        return this
    }





