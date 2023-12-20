package api

import (
	"encoding/json"
	"errors"

	"story-cook-be/pkg/errCode"
	"story-cook-be/pkg/response"
)

// ErrorResponse 返回错误信息
func ErrorResponse(err error) *response.Response {

	var unmarshalTypeError *json.UnmarshalTypeError
	if errors.As(err, &unmarshalTypeError) {
		return response.Error(err, "JSON类型不匹配", errCode.InvalidParams)
	}

	return response.Error(err, "参数错误", errCode.InvalidParams)
}
