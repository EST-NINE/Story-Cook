package main

import (
	"fmt"
	"log"

	"github.com/robfig/cron"

	"SparkForge/cache"
	"SparkForge/config"
	"SparkForge/db/dao"
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
	loading()
	r := router.NewRouter()
	err := r.Run(config.HttpPort)
	if err != nil {
		log.Fatalln(err)
	}
}

func loading() {
	config.InitFile()
	util.InitLog()
	dao.InitMysql()
	cache.InitRedis()

	go func() {
		c := cron.New()
		c.AddFunc("0 0 0 * * *", func() {
			err := cache.DeleteUserCountKeys()
			if err != nil {
				log.Println(err)
			}
		})
		c.Start()
		select {}
	}()
	fmt.Println("loading success!")
}
