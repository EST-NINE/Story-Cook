package main

import (
	"log"
	"story-cook-be/config"
	"story-cook-be/dao"
	"story-cook-be/pkg/util"
	"story-cook-be/router"
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
}
