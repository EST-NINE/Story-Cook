package api

import (
	"SparkForge/pkg/ctl"
	"SparkForge/pkg/e"
	"encoding/json"
	"errors"
)

// ErrorResponse 返回错误信息
func ErrorResponse(err error) *ctl.Response {

	var unmarshalTypeError *json.UnmarshalTypeError
	if errors.As(err, &unmarshalTypeError) {
		return ctl.ErrorResp(err, "JSON类型不匹配")
	}

	return ctl.ErrorResp(err, "参数错误", e.InvalidParams)
}
