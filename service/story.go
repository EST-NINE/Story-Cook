package service

import (
	"SparkForge/pkg/response"
	"SparkForge/repository/db/dao"
	"SparkForge/repository/db/model"
	"errors"
	"github.com/gin-gonic/gin"

	"SparkForge/pkg/util"
	"SparkForge/types"
)

type StorySrv struct {
}

// CreateStory 创建故事
func (s *StorySrv) CreateStory(ctx *gin.Context, req *types.CreateStoryReq) error {
	claims, _ := ctx.Get("claims")
	userInfo := claims.(*util.Claims)

	user, err := dao.NewUserDao(ctx).FindUserByUserId(userInfo.Id)
	if err != nil {
		util.LogrusObj.Infoln(err)
		return err
	}

	storyDao := dao.NewStoryDao(ctx)
	_, err = storyDao.FindStoryByTitleAndUserId(userInfo.Id, req.Title)
	if err == nil {
		err = errors.New("已经创建过该标题的故事哦")
		return err
	}

	story := model.Story{
		User:     *user,
		Title:    req.Title,
		Keywords: req.Keywords,
		Mood:     req.Mood,
		Content:  req.Content,
	}

	err = storyDao.CreateStory(&story)
	if err != nil {
		util.LogrusObj.Infoln(err)
		return err
	}

	return nil
}

// ListStory 得到对应用户的故事
func (s *StorySrv) ListStory(ctx *gin.Context, req *types.ListStoryReq) (resp []*response.StoryResp, total int64, err error) {
	claims, _ := ctx.Get("claims")
	userInfo := claims.(*util.Claims)

	stories, total, err := dao.NewStoryDao(ctx).ListStory(req.Page, req.Limit, userInfo.Id)
	if err != nil {
		util.LogrusObj.Infoln(err)
		return
	}

	listStoryResp := make([]*response.StoryResp, 0)
	for _, story := range stories {
		listStoryResp = append(listStoryResp, response.BuildStoryResp(story))
	}

	return listStoryResp, total, nil
}

// DeleteStory 删除故事
func (s *StorySrv) DeleteStory(ctx *gin.Context, req *types.DeleteStoryReq) error {
	claims, _ := ctx.Get("claims")
	userInfo := claims.(*util.Claims)

	err := dao.NewStoryDao(ctx).DeleteStory(userInfo.Id, req.Title)
	if err != nil {
		util.LogrusObj.Infoln(err)
		return err
	}

	return nil
}

// UpdateStory 更新故事
func (s *StorySrv) UpdateStory(ctx *gin.Context, req *types.UpdateStoryReq) (resp *response.StoryResp, err error) {
	claims, _ := ctx.Get("claims")
	userInfo := claims.(*util.Claims)

	storyDao := dao.NewStoryDao(ctx)
	_, err = storyDao.FindStoryByTitleAndUserId(userInfo.Id, req.UpdateTitle)
	if err == nil {
		err = errors.New("已经有这个标题的历史记录了哦")
		return
	}

	err = storyDao.UpdateStory(userInfo.Id, req)
	if err != nil {
		util.LogrusObj.Infoln(err)
		return
	}

	story, err := storyDao.FindStoryByTitleAndUserId(userInfo.Id, req.UpdateTitle)
	if err != nil {
		util.LogrusObj.Infoln(err)
		return
	}
	return response.BuildStoryResp(story), nil
}

// ListStoryByMood 根据mood分类查找story
func (s *StorySrv) ListStoryByMood(ctx *gin.Context, req *types.ListStoryByMoodReq) (resp []*response.StoryResp, total int64, err error) {
	claims, _ := ctx.Get("claims")
	userInfo := claims.(*util.Claims)

	stories, total, err := dao.NewStoryDao(ctx).ListStoryByMood(userInfo.Id, req)
	if err != nil {
		util.LogrusObj.Infoln(err)
		return
	}

	listStoryResp := make([]*response.StoryResp, 0)
	for _, story := range stories {
		listStoryResp = append(listStoryResp, response.BuildStoryResp(story))
	}

	return listStoryResp, total, nil
}

// ListStoryByTime 根据time分类查找story
func (s *StorySrv) ListStoryByTime(ctx *gin.Context, req *types.ListStoryByTimeReq) (resp []*response.StoryResp, total int64, err error) {
	claims, _ := ctx.Get("claims")
	userInfo := claims.(*util.Claims)

	stories, total, err := dao.NewStoryDao(ctx).ListStoryByTime(userInfo.Id, req)
	if err != nil {
		util.LogrusObj.Infoln(err)
		return
	}

	listStoryResp := make([]*response.StoryResp, 0)
	for _, story := range stories {
		listStoryResp = append(listStoryResp, response.BuildStoryResp(story))
	}

	return listStoryResp, total, nil
}
