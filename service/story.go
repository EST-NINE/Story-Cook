package service

import (
	"SparkForge/db/dao"
	"SparkForge/db/model"
	"SparkForge/pkg/ctl"
	"SparkForge/pkg/util"
	"SparkForge/types"
	"context"
	"errors"
)

type StorySrv struct {
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

	storyDao := dao.NewStoryDao(c)
	_, err = storyDao.FindStoryByTitleAndUserId(userInfo.Id, req.Title)
	if err == nil {
		err = errors.New("已经创建过该标题的故事哦")
		return
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

	listStoryResp := make([]*types.StoryResp, 0)
	for _, story := range stories {
		listStoryResp = append(listStoryResp, &types.StoryResp{
			ID:        story.ID,
			Title:     story.Title,
			Mood:      story.Mood,
			Keywords:  story.Keywords,
			Content:   story.Content,
			CreatedAt: story.CreatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	return ctl.ListResp(listStoryResp, total), nil
}

// DeleteStory 删除故事
func (s *StorySrv) DeleteStory(c context.Context, req *types.DeleteStoryReq) (resp interface{}, err error) {
	userInfo, err := ctl.GetUserInfo(c)
	if err != nil {
		util.LogrusObj.Infoln(err)
		return
	}

	err = dao.NewStoryDao(c).DeleteStory(userInfo.Id, req.Title)
	if err != nil {
		util.LogrusObj.Infoln(err)
		return
	}

	return ctl.SuccessResp(), nil
}

// UpdateStory 更新故事
func (s *StorySrv) UpdateStory(c context.Context, req *types.UpdateStoryReq) (resp interface{}, err error) {
	userInfo, err := ctl.GetUserInfo(c)
	if err != nil {
		util.LogrusObj.Infoln(err)
		return
	}

	storyDao := dao.NewStoryDao(c)
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
	storyResp := &types.StoryResp{
		ID:        story.ID,
		Title:     story.Title,
		Mood:      story.Mood,
		Keywords:  story.Keywords,
		Content:   story.Content,
		CreatedAt: story.CreatedAt.Format("2006-01-02 15:04:05"),
	}
	return ctl.SuccessWithDataResp(storyResp), nil
}

// SelectStory 根据mood分类查找story
func (s *StorySrv) SelectStory(c context.Context, req *types.SelectStoryReq) (resp interface{}, err error) {
	userInfo, err := ctl.GetUserInfo(c)
	if err != nil {
		util.LogrusObj.Infoln(err)
		return
	}

	stories, total, err := dao.NewStoryDao(c).SelectStory(userInfo.Id, req.Mood)
	if err != nil {
		util.LogrusObj.Infoln(err)
		return
	}

	listStoryResp := make([]*types.StoryResp, 0)
	for _, story := range stories {
		listStoryResp = append(listStoryResp, &types.StoryResp{
			ID:        story.ID,
			Title:     story.Title,
			Mood:      story.Mood,
			Keywords:  story.Keywords,
			Content:   story.Content,
			CreatedAt: story.CreatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	return ctl.ListResp(listStoryResp, total), nil
}
