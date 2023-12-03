package ctl

import "SparkForge/pkg/e"

// Response 基础序列化器
type Response struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
	Msg    string      `json:"msg"`
	Error  string      `json:"error"`
}

// DataListResp 带有总数的Data结构
type DataListResp struct {
	Item  interface{} `json:"item"`
	Total int64       `json:"total"`
}

// ListResp 带有总数的列表构建器
func ListResp(items interface{}, total int64) Response {
	return Response{
		Status: 200,
		Data: DataListResp{
			Item:  items,
			Total: total,
		},
		Msg: "ok",
	}
}

// SuccessResp 成功返回
func SuccessResp(code ...int) *Response {
	status := e.SUCCESS
	if code != nil {
		status = code[0]
	}

	r := &Response{
		Status: status,
		Data:   "操作成功",
		Msg:    e.GetMsg(status),
	}

	return r
}

// SuccessWithDataResp 带data成功返回
func SuccessWithDataResp(data interface{}, code ...int) *Response {
	status := e.SUCCESS
	if code != nil {
		status = code[0]
	}

	r := &Response{
		Status: status,
		Data:   data,
		Msg:    e.GetMsg(status),
	}

	return r
}

// ErrorResp 错误返回
func ErrorResp(err error, data string, code ...int) *Response {
	status := e.ERROR
	if code != nil {
		status = code[0]
	}

	r := &Response{
		Status: status,
		Msg:    e.GetMsg(status),
		Data:   data,
		Error:  err.Error(),
	}

	return r
}
