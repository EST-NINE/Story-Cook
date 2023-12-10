package config

import (
	"gopkg.in/ini.v1"

	"SparkForge/pkg/util"
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

	AppId     string
	ApiKey    string
	ApiSecret string
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
	LoadSpark(file)
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

func LoadSpark(file *ini.File) {
	AppId = file.Section("spark").Key("AppId").String()
	ApiKey = file.Section("spark").Key("ApiKey").String()
	ApiSecret = file.Section("spark").Key("ApiSecret").String()
}
