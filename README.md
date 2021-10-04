# Jinghong-LoveWall
精弘表白墙项目
是精弘网络技术部 试用期项目
#如何使用后端测试？
0. 安装好Go 1.13
1. clone整个项目
   ```
   git clone git@github.com:Forever-Jimmy/Jinghong-LoveWall.git
   ```

2. 进入项目根目录，也就是`main.go`文件所在的目录。
3. 执行:
    ```
    go run main.go
    ```
默认端口为8080
可以使用`localhost:8080`访问

例如要实现用户注册
向`localhost:8080/api/register`POST即可
# API

## [1] [用户注册] [POST]
URL: domain:port/api/register
#### 传入数据
```
{
    "email" : "node520@example.com", // 电子邮件地址
    "username" : "node520",  // 用户昵称
    "password" : "5AA765D61D8327DE"  // 用户密码，用MD5（16位大写）加密，样例所示为"password"加密后的密文
}
```

#### 返回数据
```
{
    "return_value": 0 // 表示用户注册成功
                    -1 // 表示用户已存在（电子邮件已经注册）
                    -2 // 表示服务器错误
                    -3 // 表示请求错误
}
```

## [2] [用户登录] [POST]
URL: domain:port/api/login
#### 传入数据

```
{
    "email" : "node520@example.com", // 电子邮件地址
    "password" : "5AA765D61D8327DE"  // 用户密码
}
```

#### 返回数据

```
{
    "return_value" : 0, // 登录成功
                     -1, // 用户不存在
                     -2, // 密码错误
                     -3, // 重复登录
                     -4, // 服务器错误
                     -5, // 重复登录
                     -6, // 参数错误
    "token" : "xxxx" // token 客户端保存，需要token时需要上传。如果登录失败，返回的token是空
}
```


## [3] 用户登出 [POST]

URL: domain:port/api/logout

#### 传入数据
```
{
    "token" : "xxx"  // 传入token
}
```

#### 返回数据
```
{
    "return_value" : 0 // 成功登出
                     -1 // 参数错误
                     -2 // 服务器错误
                     -3 // token错误

}
```


## [4] [新建信息] [POST]
URL: domain:port/api/new_message
#### 传入数据
```
{
    "token" : "xxx",
    "content" : "内容"  // 信息内容
    "anonymous": 0  // 不匿名
                 1  // 匿名
    "permit_comment" : 0 // 不允许评论
                       -1 // 允许评论
}
```

```
{
    "return_value" : 0 // 成功
                     -1 // 参数错误
                     -2 // token错误
                     -3 // 服务器错误

}
```

## [5] 获取指定的信息 [GET]
URL: domain:port/api/message

#### 传入数据
```
{
    "id": "1" // 传入信息的id
}
```

#### 返回数据
```
{
    "return_value" : 0 // 成功
                     -1 // 请求错误
                     -2 // 服务器错误
                     -3 // 无该信息
    "uid" : 0 // 0 表示该信息是匿名的，防止泄露，客户端无法得知信息的用户
    "content" : "内容" // 信息的内容
    "permit_comment" : 0 // 不允许评论
                       1 // 允许评论
    "create_time"  // 创建时间
    "update_time" // 更新时间

}
```