package errMsg

const (
	SUCCESS       = 200 // 操作成功
	ERROR         = 500 // 操作失败
	InvalidParams = 400 // 请求参数错误

	ErrorExistNick          = 10001 // 用户昵称已存在
	ErrorExistUser          = 10002 // 用户已存在
	ErrorNotExistUser       = 10003 // 用户不存在
	ErrorNotCompare         = 10004 // 不匹配
	ErrorNotComparePassword = 10005 // 密码不匹配
	ErrorFailEncryption     = 10006 // 加密失败

	ErrorAuthCheckTokenFail    = 30001 // Token鉴权失败
	ErrorAuthCheckTokenTimeout = 30002 // Token已超时
	ErrorAuthToken             = 30003 // Token生成失败

	ErrorDatabase = 40001 // 数据库操作出错，请重试
)
