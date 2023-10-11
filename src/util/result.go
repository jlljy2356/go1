package util

import (
	"gin-blog/types" //这可能是一个自定义的包，用于定义和处理 Web 应用程序中的对象、数据模型或数据结构。
	"net/http"       //用于处理 HTTP 协议相关的功能，例如创建 HTTP 服务器、处理 HTTP 请求和响应
	"time"           //用于处理 HTTP 协议相关的功能，例如创建 HTTP 服务器、处理 HTTP 请求和响应

	"github.com/gin-gonic/gin" //这是 Go 语言标准库中的一个包，用于处理 HTTP 协议相关的功能，例如创建 HTTP 服务器、处理 HTTP 请求和响应
)

// 返回结果
type Result struct {
	Time time.Time   `json:"time"` //Time：表示返回结果的时间戳。
	Code int         `json:"code"` //Code：表示返回结果的状态码。
	Msg  string      `json:"msg"`  //Msg：表示返回结果的消息。
	Data interface{} `json:"data"` //Data：表示返回结果的数据对象。
}

// 成功
func Succes(c *gin.Context, data interface{}) {
	if data == nil {
		data = gin.H{} //则将其赋值为一个空的 gin.H 对象。gin.H 是 Gin 框架提供的方便构建字典的快捷方式。
	}
	res := Result{}                                           //res := Result{}：创建一个空的 Result 结构体对象，并将其赋值给变量 res。
	res.Time = time.Now()                                     //将当前时间赋值给 res 的 Time 字段，表示返回结果的时间戳。
	res.Code = int(types.ApiCode.SUCCESS)                     //将 types.ApiCode.SUCCESS 的值转换为 int 类型，并赋值给 res 的 Code 字段，表示返回结果的状态码。
	res.Msg = types.ApiCode.GetMessage(types.ApiCode.SUCCESS) //调用 types.ApiCode.GetMessage 方法，根据状态码 types.ApiCode.SUCCESS 获取对应的消息，并赋值给 res 的 Msg 字段，表示返回结果的消息。
	res.Data = data                                           //将传入的 data 赋值给 res 的 Data 字段，表示返回结果的数据对象。

	c.JSON(http.StatusOK, res)
}

// 出错
func Error(c *gin.Context, code int, msg string) {
	res := Result{}
	res.Time = time.Now()
	res.Code = code
	res.Msg = msg
	res.Data = gin.H{}

	c.JSON(http.StatusOK, res)
}
