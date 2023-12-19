package controller

import (
	"SparkForge/pkg/errCode"
)

// Response 基础序列化器
type Response struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
	Msg    string      `json:"msg"`
	Error  string      `json:"error"`
}

// SuccessResp 成功返回
func SuccessResp() *Response {
	return &Response{
		Status: errCode.SUCCESS,
		Data:   "操作成功",
		Msg:    errCode.GetMsg(errCode.SUCCESS),
	}
}

// SuccessWithDataResp 带data成功返回
func SuccessWithDataResp(data interface{}) *Response {
	return &Response{
		Status: errCode.SUCCESS,
		Data:   data,
		Msg:    errCode.GetMsg(errCode.SUCCESS),
	}
}

// ErrorResp 错误返回
func ErrorResp(err error, data string, code ...int) *Response {
	status := errCode.ERROR
	if code != nil {
		status = code[0]
	}

	return &Response{
		Status: status,
		Msg:    errCode.GetMsg(status),
		Data:   data,
		Error:  err.Error(),
	}
}

// DataListResp 带有总数的Data结构
type DataListResp struct {
	Item  interface{} `json:"item"`
	Total int64       `json:"total"`
}

// ListResp 带有总数的列表构建器
func ListResp(items interface{}, total int64) *Response {
	return &Response{
		Status: 200,
		Data: DataListResp{
			Item:  items,
			Total: total,
		},
		Msg: "查询列表成功",
	}
}
