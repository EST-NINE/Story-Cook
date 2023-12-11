package dao

import (
	"SparkForge/db/model"
)

// 执行数据迁移
func migration() {
	// 自动迁移模式
	err := ormDB.Set("gorm:table_options", "charset=utf8mb4").
		AutoMigrate(&model.User{}, &model.Story{}, &model.Menu{}, &model.UserMenu{})
	if err != nil {
		return
	}
}
