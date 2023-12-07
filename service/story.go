package service

import (
	"SparkForge/db/dao"
	"SparkForge/db/model"
	"SparkForge/pkg/ctl"
	"SparkForge/pkg/util"
	"SparkForge/types"
	"context"
	"sync"
)

type StorySrv struct {
}

var StorySrvIns *StorySrv
var StorySrvOnce sync.Once

func GetStorySrv() *StorySrv {
	StorySrvOnce.Do(func() {
		StorySrvIns = &StorySrv{}
	})
	return StorySrvIns
}

// CreateStory 创建故事
func (s *StorySrv) CreateStory(c context.Context, req *types.CreateStoryReq) (resp interface{}, err error) {
	userInfo, err := ctl.GetUserInfo(c)
	if err != nil {
		util.LogrusObj.Infoln(err)
		return
	}

	user, err := dao.NewUserDao(c).FindUserByUserId(userInfo.Id)
	if err != nil {
		util.LogrusObj.Infoln(err)
		return
	}

	story := model.Story{
		User:     *user,
		Title:    req.Title,
		Keywords: req.Keywords,
		Mood:     req.Mood,
		Content:  req.Content,
	}

	err = dao.NewStoryDao(c).CreateStory(&story)
	if err != nil {
		util.LogrusObj.Infoln(err)
		return
	}

	return ctl.SuccessResp(), nil
}

// ListStory 得到对应用户的故事
func (s *StorySrv) ListStory(c context.Context, req *types.ListStoryReq) (resp interface{}, err error) {
	userInfo, err := ctl.GetUserInfo(c)
	if err != nil {
		util.LogrusObj.Infoln(err)
		return
	}

	stories, total, err := dao.NewStoryDao(c).ListStory(req.Page, req.Limit, userInfo.Id)
	if err != nil {
		util.LogrusObj.Infoln(err)
		return
	}

	listStoryResp := make([]*types.ListStoryResp, 0)
	for _, story := range stories {
		listStoryResp = append(listStoryResp, &types.ListStoryResp{
			Title:     story.Title,
			Mood:      story.Mood,
			Keywords:  story.Keywords,
			Content:   story.Content,
			CreatedAt: story.CreatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	return ctl.ListResp(listStoryResp, total), nil
}
