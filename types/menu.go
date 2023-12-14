package types

type SelectMenuReq struct {
	Keywords string `json:"keywords" binding:"required" example:"大学生+未完成的作业"`
}

type CreateUserMenuReq struct {
	Keywords string `json:"keywords" binding:"required" example:"大学生+未完成的作业"`
	Content  string `json:"content" binding:"required" example:"彩蛋:焦虑的夜晚"`
}

type ListUserMenuReq struct {
	Page  int `json:"page" example:"1"`
	Limit int `json:"limit" example:"10"`
}

type MenuResp struct {
	ID       uint   `json:"id"`
	Keywords string `json:"keywords"`
	Content  string `json:"content"`
	CreateAt string `json:"create_at"`
}
