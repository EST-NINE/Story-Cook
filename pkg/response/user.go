package response

import (
	"story-cook-be/model"
)

type UserResp struct {
	ID       uint   `json:"id"`        // 用户ID
	UserName string `json:"user_name"` // 用户名
	Kitchen  string `json:"kitchen"`   // 厨房名
	CreateAt string `json:"create_at"` // 创建
	Count    int    `json:"count"`     // 当天剩余合成次数
}

type TokenDataResp struct {
	User  interface{} `json:"user"`
	Token string      `json:"token"`
}

func BuildUserResp(user *model.User) *UserResp {
	return &UserResp{
		ID:       user.ID,
		UserName: user.UserName,
		Kitchen:  user.Kitchen,
		CreateAt: user.CreatedAt.Format("2006-01-02 15:04:05"),
	}
}
