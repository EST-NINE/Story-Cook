package e

var MsgFlags = map[int]string{
	SUCCESS:       "操作成功",
	ERROR:         "操作失败",
	InvalidParams: "请求参数错误",

	ErrorExistNick:          "用户昵称已存在",
	ErrorExistUser:          "用户已存在",
	ErrorNotExistUser:       "用户不存在",
	ErrorNotCompare:         "不匹配",
	ErrorNotComparePassword: "密码不匹配",
	ErrorFailEncryption:     "加密失败",

	ErrorAuthCheckTokenFail:    "Token鉴权失败",
	ErrorAuthCheckTokenTimeout: "Token已超时",
	ErrorAuthToken:             "Token生成失败",
	ErrorAuth:                  "Token错误",

	ErrorDatabase:   "数据库操作出错，请重试",
	ErrorOss:        "OSS操作出错",
	ErrorUploadFile: "文件上传失败",
}

// GetMsg 获取状态码对应的信息
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[ERROR]
}
