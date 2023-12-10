package dao

import (
	"context"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"

	"SparkForge/config"
	"SparkForge/pkg/util"
)

var db *gorm.DB

func InitMysql() {
	conn := strings.Join([]string{config.DbUser, ":", config.DbPassWord, "@tcp(", config.DbHost, ":", config.DbPort, ")/", config.DbName, "?charset=utf8mb4&parseTime=true&loc=Local"}, "")
	var ormLogger logger.Interface
	if gin.Mode() == "debug" {
		ormLogger = logger.Default.LogMode(logger.Info)
	} else {
		ormLogger = logger.Default
	}
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       conn,  // DSN data source name
		DefaultStringSize:         256,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据版本自动配置
	}), &gorm.Config{
		Logger: ormLogger, // 打印日志
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 表名不加s
		},
	})
	if err != nil {
		util.LogrusObj.Error(err)
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(20)                  // 设置数据库连接池中空闲连接的最大数量
	sqlDB.SetMaxOpenConns(100)                 // 设置数据库连接池中最大连接数
	sqlDB.SetConnMaxLifetime(time.Second * 30) // 设置数据库连接的最大生命周期
	migration()
}

func NewDBClient(c context.Context) *gorm.DB {
	return db.WithContext(c)
}
