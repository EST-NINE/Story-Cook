package dao

import (
	"SparkForge/db/model"
	"context"
	"gorm.io/gorm"
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

// FindStoryIdByTitle 根据故事的标题找到故事id
func (dao *StoryDao) FindStoryIdByTitle(title string) (id uint, err error) {
	err = dao.DB.Model(&model.Story{}).Where("title = ?", title).First(&id).Error

	return
}

// FindStoryByIdAndUserId 根据用户id和故事id查找故事
func (dao *StoryDao) FindStoryByIdAndUserId(id, uid uint) (story *model.Story, err error) {
	err = dao.DB.Model(&model.Story{}).Where("id = ? AND uid = ?", id, uid).First(&story).Error

	return
}

// CreateStory 创建故事
func (dao *StoryDao) CreateStory(story *model.Story) (err error) {
	err = dao.DB.Model(&model.Story{}).Create(&story).Error

	return
}

// ListStory 得到故事列表
func (dao *StoryDao) ListStory(page, limit int, uid uint) (stories []model.Story, total int64, err error) {
	err = dao.DB.Model(&model.Story{}).Preload("User").Where("uid = ?", uid).
		Count(&total).
		Limit(limit).Offset((page - 1) * limit).
		Find(&stories).Error

	return
}
