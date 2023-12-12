package main

import (
	"SparkForge/cache"
	"SparkForge/config"
	"SparkForge/db/dao"
	"SparkForge/pkg/util"
	"SparkForge/router"
	"fmt"
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
	fmt.Println("loading success!")
}
