package dao

import (
	model2 "story-cook-be/model"
)

// 执行数据迁移
func migration() {
	// 自动迁移模式
	err := ormDB.Set("gorm:table_options", "charset=utf8mb4").
		AutoMigrate(&model2.User{}, &model2.Story{}, &model2.Menu{}, &model2.UserMenu{}, &model2.UserCount{})
	if err != nil {
		return
	}
}
