package dao

import (
	"context"
	"gorm.io/gorm"

	"SparkForge/db/model"
	"SparkForge/types"
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
func (dao *StoryDao) CreateStory(story *model.Story) (err error) {
	err = dao.DB.Model(&model.Story{}).Create(&story).Error

	return
}

// ListStory 得到故事列表
func (dao *StoryDao) ListStory(page, limit int, uid uint) (stories []model.Story, total int64, err error) {
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

// SelectStory 根据mood分类故事
func (dao *StoryDao) SelectStory(uid uint, mood string) (stories []model.Story, total int64, err error) {
	err = dao.DB.Model(&model.Story{}).Preload("User").Where("uid = ? AND mood = ?", uid, mood).
		Count(&total).
		Find(&stories).Error

	return
}
