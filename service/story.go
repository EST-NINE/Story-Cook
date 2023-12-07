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
