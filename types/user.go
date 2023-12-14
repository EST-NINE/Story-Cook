package types

type UserServiceReq struct {
	UserName string `json:"user_name" binding:"required" example:"john"`
	Password string `json:"password" binding:"required" example:"12345678"`
}

type UserUpdatePwdReq struct {
	OriginPwd string `json:"originPwd" binding:"required" example:"12345678"`
	UpdatePwd string `json:"updatePwd" binding:"required" example:"123456789"`
}

type UserUpdateInfoReq struct {
	UpdateName string `json:"update_name" example:"John Doe"`
	Kitchen    string `json:"kitchen" example:"John Doe的厨房"`
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
