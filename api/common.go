package api

import (
	"encoding/json"
	"errors"

	"SparkForge/pkg/ctl"
	"SparkForge/pkg/e"
)

// ErrorResponse 返回错误信息
func ErrorResponse(err error) *ctl.Response {

	var unmarshalTypeError *json.UnmarshalTypeError
	if errors.As(err, &unmarshalTypeError) {
		return ctl.ErrorResp(err, "JSON类型不匹配", e.InvalidParams)
	}

	return ctl.ErrorResp(err, "参数错误", e.InvalidParams)
}
