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
