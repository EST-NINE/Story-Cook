package api

import (
	"encoding/json"
	"errors"

	"SparkForge/pkg/controller"
	"SparkForge/pkg/errMsg"
)

// ErrorResponse 返回错误信息
func ErrorResponse(err error) *controller.Response {

	var unmarshalTypeError *json.UnmarshalTypeError
	if errors.As(err, &unmarshalTypeError) {
		return controller.ErrorResp(err, "JSON类型不匹配", errMsg.InvalidParams)
	}

	return controller.ErrorResp(err, "参数错误", errMsg.InvalidParams)
}
