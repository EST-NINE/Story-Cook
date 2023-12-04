# HackWeek

# 接口说明
## 1.1：注册

- 请求路径：api/v1/user/register
- 请求方法：POST
- 请求体

|   参数    |           含义           |              备注              |
| :-------: | :----------------------: | :----------------------------: |
| user_name |         用户姓名         | 只能包含汉字、大小写字母、数字 |
| password  | 经前端加密过后的用户密码 |  经前端加密，后端存储时也加密  |

```
{
     "user_name":"lxy",
     "password":"123456"
}
```



- 响应体

|  参数  |         含义         | 备注 |
| :----: | :------------------: | ---- |
| status | 后端响应返回的状态码 |      |
|  data  |     数据或者信息     |      |
|  msg   |       响应信息       |      |
| error  |     详细错误信息     |      |

```
{
    "status": 200,
    "data": "操作成功",
    "msg": "操作成功",
    "error": ""
}
```



## 1.2：登录

- 请求路径：api/v1/user/login
- 请求方法：POST
- 请求体

|   参数    |   含义   |              备注              |
| :-------: | :------: | :----------------------------: |
| user_name | 用户姓名 | 只能包含汉字、大小写字母、数字 |
| password  | 用户密码 |  经前端加密，后端存储时也加密  |

```
{
     "user_name":"lxy",
     "password":"123456"
}
```



- 响应体

|    参数     |         含义         |     备注     |
| :---------: | :------------------: | :----------: |
|   status    | 后端响应返回的状态码 |              |
|    data     |     数据或者信息     |              |
| data->user  |       用户信息       |              |
| data->token |      认证token       | 有效期为一天 |
|     msg     |       响应信息       |              |
|    error    |     详细错误信息     |              |

```
{
    "status": 200,
    "data": {
        "user": {
            "id": 2,
            "user_name": "lxy",
            "create_at": 1701557314
        },
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MiwidXNlcm5hbWUiOiJseHkiLCJleHAiOjE3MDE2NDM3MjUsImlzcyI6InRvLWRvLWxpc3QifQ.XOQFVHIcdttR0BpQFANUcbLmgstmZhQh5BM6hwm2Qek"
    },
    "msg": "操作成功",
    "error": ""
}
```

## 2.1：更改密码

- 请求路径：api/v1/user/updatePwd
- 请求方法：POST
- 请求体(需要token认证)

|   参数   |   含义   |         备注         |
| :------: | :------: | :------------------: |
| password | 用户密码 |  |

```
{
     "password":"123456"
}
```



- 响应体

|  参数  |         含义         | 备注 |
| :----: | :------------------: | :--: |
| status | 后端响应返回的状态码 |      |
|  data  |     数据或者信息     |      |
|  msg   |       响应信息       |      |
| error  |     详细错误信息     |      |

```
{
    "status": 200,
    "data": null,
    "msg": "修改成功!",
    "error": ""
}
```