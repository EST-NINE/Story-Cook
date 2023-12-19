package main

import (
	"SparkForge/repository/cache"
	"SparkForge/repository/db/dao"
	"log"

	"github.com/robfig/cron"

	"SparkForge/config"
	_ "SparkForge/docs" // 导入自动生成的docs文档
	"SparkForge/pkg/util"
	"SparkForge/router"
)

// @title		Story-Cook
// @version		1.0
// @description	Story-Cook API文档
// @host		localhost:8082
// @BasePath    /api/v1
func main() {
	r := router.NewRouter()
	err := r.Run(config.HttpPort)
	if err != nil {
		log.Fatalln(err)
	}
}

func init() {
	config.InitFile()
	util.InitLog()
	dao.InitMysql()
	cache.InitRedis()

	go func() {
		c := cron.New()
		_ = c.AddFunc("0 0 0 * * *", func() { _ = cache.DeleteUserCountKeys() }) // 每天零点更新可调用api的次数
		c.Start()
		select {}
	}()
}
