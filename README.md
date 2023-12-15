# SparkForge (Former HackWeek)
使用GoLang编写

展示地址：[故事厨房](https://kitchen.heuluck.top/)

# 2023“摆脱社恐对不对”组项目（重构）
项目名：故事厨房

前端：请见[Heuluck](https://github.com/Heuluck/Team-Story-Cook-FE)的仓库

# 接口文档
启动项目后，直接访问 http://localhost:8082/swagger/index.html

## 项目运行
### 手动执行
**本项目使用`Go Mod`管理依赖。**

**下载依赖**
```shell
go mod tidy
```

**运行**
```shell
go run main.go
```

## 主要功能介绍
- 用户注册登录(jwt-go)
- 用户基本信息修改和获取，修改密码
- 调用星火api生成故事
- 历史记录的加入，删除，浏览，选择等
- 彩蛋的加入，删除，浏览，选择等
- 分页功能

## 主要依赖
| 名称         | 版本      |
|------------|---------|
| golang     | 1.21.0  |
| gin        | v1.9.1  |
| gorm       | v1.9.16 |
| mysql      | v1.7.0  |
| redis      | v6.15.9 |
| jwt-go     | v3.2.0  |
| logrus     | v1.9.3  |
| go-swagger | v1.16.2 |

## 项目结构
```
├─api                    # 用于定义接口函数
├─config                 # 配置文件
├─docs                   # swagger接口文档
├─logs                   # 日志打印
├─middleware             # 中间件
├─pkg
│  ├─controller          # 处理响应
│  ├─errMsg              # 封装错误码
│  └─util                # 工具函数
├─repository
│  ├─cache               # 缓存操作
│  └─db
│      ├─dao             # 数据库访问对象
│      └─model           # 数据库模型
├─router                 # 路由逻辑处理
├─service                # 接口函数的实现
└─types                  # 将数据序列化为 json 的函数，便于返回给前端

```

## 【重要】配置项目
根据实际情况在config包下配置config.ini文件(配置文件可以将`config.ini.example`重命名为`config.ini`)
```ini
[service]
AppMode = debug
HttpPort = :8082
# 运行端口号 8082端口

[redis]
RedisDb = redis
# redis ip地址和端口号
RedisAddr =
# redis 密码
RedisPw =
# redis 名字
RedisDbName =

[mysql]
Db = mysql
# mysql ip地址
DbHost =
# mysql 端口号
DbPort =
# mysql 用户名
DbUser =
# mysql 密码
DbPassWord =
# mysql 名字
DbName =

[spark]
# Spark 应用ID
AppId = 
# Spark API密钥
ApiKey =
# Spark API秘钥
ApiSecret = 
```

## 简要说明
- mysql 是存储主要的数据。
- redis 用来存储调用api的次数。


