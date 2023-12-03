package config

import (
	"SparkForge/pkg/util"
	"gopkg.in/ini.v1"
)

var (
	AppMode  string
	HttpPort string

	RedisDb     string
	RedisAddr   string
	RedisPw     string
	RedisDbName string

	Db         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassWord string
	DbName     string
)

func InitFile() {
	file, err := ini.Load("./config/config.ini")
	if err != nil {
		util.LogrusObj.Println(err)
		panic(err)
	}
	LoadServer(file)
	LoadMysqlData(file)
	LoadRedis(file)
}

func LoadServer(file *ini.File) {
	AppMode = file.Section("service").Key("AppMode").String()
	HttpPort = file.Section("service").Key("HttpPort").String()
}

func LoadMysqlData(file *ini.File) {
	Db = file.Section("mysql").Key("Db").String()
	DbHost = file.Section("mysql").Key("DbHost").String()
	DbPort = file.Section("mysql").Key("DbPort").String()
	DbUser = file.Section("mysql").Key("DbUser").String()
	DbPassWord = file.Section("mysql").Key("DbPassWord").String()
	DbName = file.Section("mysql").Key("DbName").String()
}

func LoadRedis(file *ini.File) {
	RedisAddr = file.Section("redis").Key("RedisAddr").String()
	RedisPw = file.Section("redis").Key("RedisPw").String()
	RedisDbName = file.Section("redis").Key("RedisDbName").String()
}
