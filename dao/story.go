package dao

import (
	"context"
	"errors"
	"story-cook-be/model"
	"time"

	"gorm.io/gorm"

	"story-cook-be/types"
)

type StoryDao struct {
	*gorm.DB
}

func NewStoryDao(c context.Context) *StoryDao {
	if c == nil {
		c = context.Background()
	}
	return &StoryDao{NewDBClient(c)}
}

// FindStoryByIdAndUserId 根据用户id和故事id查找故事
func (dao *StoryDao) FindStoryByIdAndUserId(uid, id uint) (story *model.Story, err error) {
	err = dao.DB.Model(&model.Story{}).Where("uid = ? AND id = ? ", uid, id).First(&story).Error

	return
}

// FindStoryByTitleAndUserId 根据故事title查找故事
func (dao *StoryDao) FindStoryByTitleAndUserId(uid uint, title string) (story *model.Story, err error) {
	err = dao.DB.Model(&model.Story{}).Where("uid = ? AND title = ? ", uid, title).First(&story).Error

	return
}

// CreateStory 创建故事
func (dao *StoryDao) CreateStory(story *model.Story) error {
	return dao.DB.Model(&model.Story{}).Create(&story).Error
}

// ListStory 得到故事列表
func (dao *StoryDao) ListStory(page, limit int, uid uint) (stories []*model.Story, total int64, err error) {
	err = dao.DB.Model(&model.Story{}).Preload("User").Where("uid = ?", uid).
		Count(&total).
		Order("created_at DESC").
		Limit(limit).Offset((page - 1) * limit).
		Find(&stories).Error

	return
}

// DeleteStory 删除故事
func (dao *StoryDao) DeleteStory(uid uint, title string) error {
	story, err := dao.FindStoryByTitleAndUserId(uid, title)
	if err != nil {
		return err
	}

	return dao.Delete(&story).Error
}

// UpdateStory 更新故事
func (dao *StoryDao) UpdateStory(uid uint, req *types.UpdateStoryReq) error {
	story := new(model.Story)
	err := dao.DB.Model(&model.Story{}).Where("uid = ? AND title = ?", uid, req.Title).First(&story).Error
	if err != nil {
		return err
	}

	if req.UpdateTitle != "" {
		story.Title = req.UpdateTitle
	}

	if req.UpdateContent != "" {
		story.Content = req.UpdateContent
	}

	return dao.Save(story).Error
}

// ListStoryByMood 根据mood分类查找story
func (dao *StoryDao) ListStoryByMood(uid uint, req *types.ListStoryByMoodReq) (stories []*model.Story, total int64, err error) {
	err = dao.DB.Model(&model.Story{}).Preload("User").Where("uid = ? AND mood = ?", uid, req.Mood).
		Count(&total).
		Order("created_at DESC").
		Offset((req.Page - 1) * req.Limit).
		Limit(req.Limit).
		Find(&stories).Error

	return
}

// ListStoryByTime 根据time分类查找story
func (dao *StoryDao) ListStoryByTime(uid uint, req *types.ListStoryByTimeReq) (stories []*model.Story, total int64, err error) {
	// 获取当前时间
	currentTime := time.Now()
	// 设置时间查询的起始时间
	startTime := currentTime

	// 根据时间标识设置不同的起始时间
	switch req.TimeFlag {
	case "daily":
		startTime = currentTime.AddDate(0, 0, -1) // 前一天
	case "weekly":
		startTime = currentTime.AddDate(0, 0, -7) // 前一周
	case "monthly":
		startTime = currentTime.AddDate(0, -1, 0) // 前一个月
	case "yearly":
		startTime = currentTime.AddDate(-1, 0, 0) // 前一年
	default:
		err = errors.New("没有对应查询的时间标识哦")
		return
	}

	err = dao.DB.Model(&model.Story{}).Preload("User").Where("uid = ? AND created_at >= ?", uid, startTime).
		Count(&total).
		Order("created_at DESC").
		Offset((req.Page - 1) * req.Limit).
		Limit(req.Limit).
		Find(&stories).Error

	return
}
