package types

type UserServiceReq struct {
	UserName string `json:"user_name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserUpdatePwdReq struct {
	OriginPwd string `json:"originPwd" binding:"required"`
	UpdatePwd string `json:"updatePwd" binding:"required"`
}

type UseUpdateInfoReq struct {
	UpdateName string `json:"update_name"`
	Kitchen    string `json:"kitchen"`
}

type TokenDataResp struct {
	User  interface{} `json:"user"`
	Token string      `json:"token"`
}

type UserResp struct {
	ID       uint   `json:"id"`        // 用户ID
	UserName string `json:"user_name"` // 用户名
	Kitchen  string `json:"kitchen"`   // 厨房名
	CreateAt string `json:"create_at"` // 创建
	Count    uint64 `json:"count"`     // 当天剩余合成次数
}
