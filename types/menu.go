package types

type SelectMenuReq struct {
	Keywords string `json:"keywords" binding:"required"`
}

type CreateUserMenuReq struct {
	Keywords string `json:"keywords" binding:"required"`
	Content  string `json:"content" binding:"required"`
}

type ListUserMenuReq struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
}

type MenuResp struct {
	ID       uint   `json:"id"`
	Keywords string `json:"keywords"`
	Content  string `json:"content"`
	CreateAt string `json:"create_at"`
}
