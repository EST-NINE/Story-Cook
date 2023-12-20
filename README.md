# SparkForge (Former HackWeek)
使用go编写

展示地址：[故事厨房](https://kitchen.heuluck.top/)

# 2023“摆脱社恐对不对”组项目（重构）
项目名：故事厨房

前端：请见[Heuluck](https://github.com/ncuhome/Story-Cook-FE)的仓库

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
| jwt-go     | v3.2.0  |
| logrus     | v1.9.3  |

## 项目结构
```
├─api                    # 用于定义接口函数
├─config                 # 配置文件
├─dao                    # 数据库访问对象
├─logs                   # 日志打印
├─middleware             # 中间件
├─model                  # 数据库模型
├─pkg
│  ├─errCode             # 封装错误码
│  ├─response            # 封装统一的响应结构体和响应函数
│  └─util                # 工具函数
├─router                 # 路由逻辑处理
├─service                # 接口函数的实现
└─types                  # 定于请求结构体，便于前端发送数据
```


