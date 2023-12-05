package types

type UserServiceReq struct {
	UserName string `form:"user_name" json:"user_name"`
	Password string `form:"password" json:"password"`
}

type UserUpdatePwdReq struct {
	OriginPwd string `json:"originPwd" form:"originPwd"`
	UpdatePwd string `json:"updatePwd" form:"updatePwd"`
}

type UseUpdateInfoReq struct {
	UpdateName string `json:"update_name" form:"update_name"`
}

type TokenDataResp struct {
	User  interface{} `json:"user"`
	Token string      `json:"token"`
}

type UserResp struct {
	ID       uint   `json:"id" form:"id"`               // 用户ID
	UserName string `json:"user_name" form:"user_name"` // 用户名
	CreateAt int64  `json:"create_at" form:"create_at"` // 创建
}
