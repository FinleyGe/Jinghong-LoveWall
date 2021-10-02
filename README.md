# Jinghong-LoveWall
精弘表白墙项目
是精弘网络技术部 试用期项目

# API

## [1] [用户注册] [POST]

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