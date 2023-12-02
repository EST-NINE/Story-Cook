package types

type UserServiceReq struct {
	UserName string `form:"user_name" json:"user_name" binding:"required,min=2,max=16"`
	Password string `form:"password" json:"password" binding:"required,min=6,max=16"`
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
