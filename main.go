package main

import (
	"SparkForge/cache"
	"SparkForge/config"
	"SparkForge/db/dao"
	"SparkForge/pkg/util"
	"SparkForge/router"
	"fmt"
	"github.com/robfig/cron"
	"log"
)

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
