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

// FindStoryByIdAndUserId 根据用户id和故事id查找故事
func (dao *StoryDao) FindStoryByIdAndUserId(uid, id uint) (story *model.Story, err error) {
	err = dao.DB.Model(&model.Story{}).Where("uid = ? AND id = ? ", uid, id).First(&story).Error

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

// DeleteStory 删除故事
func (dao *StoryDao) DeleteStory(uid, id uint) error {
	story, err := dao.FindStoryByIdAndUserId(uid, id)
	if err != nil {
		return err
	}

	return dao.Delete(&story).Error
}
