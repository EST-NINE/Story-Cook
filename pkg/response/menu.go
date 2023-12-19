package response

import "SparkForge/repository/db/model"

type MenuResp struct {
	ID       uint   `json:"id"`
	Keywords string `json:"keywords"`
	Content  string `json:"content"`
	CreateAt string `json:"create_at"`
}

func BuildMenuResp(menu *model.Menu) *MenuResp {
	return &MenuResp{
		ID:       menu.ID,
		Keywords: menu.Keywords,
		Content:  menu.Content,
		CreateAt: menu.CreatedAt.Format("2006-01-02 15:04:05"),
	}
}

func BuildUserMenuResp(userMenu *model.UserMenu) *MenuResp {
	return &MenuResp{
		ID:       userMenu.ID,
		Keywords: userMenu.Keywords,
		Content:  userMenu.Content,
		CreateAt: userMenu.CreatedAt.Format("2006-01-02 15:04:05"),
	}
}
