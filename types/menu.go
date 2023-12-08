package types

type SelectMenuReq struct {
	Keywords string `json:"keywords" binding:"required"`
}

type MenuResp struct {
	ID       uint   `json:"id"`
	Keywords string `json:"keywords"`
	Content  string `json:"content"`
}
