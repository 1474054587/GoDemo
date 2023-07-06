# first_work_jty
姜同远的go作业
## 2 并发控制
> /OtherWorks/work2/main.go
## 3 Go的优缺点
> /OtherWorks/work3/go.md
## 4 TCP SOCKET
> /OtherWorks/work4
## 7 Context源码阅读
> /OtherWorks/work7/context.md
## 1,5,6 USER模块
### 配置
1. 创建本项目所用的数据库:
```MySQL
CREATE DATABASE first_work_jty DEFAULT CHARSET=utf8mb4;
```
2. 配置`/conf/config.ini`
```ini
port = 启动服务的端口号
release = false

[mysql]
user = 数据库用户名
password = 数据库密码
host = 数据库host地址
port = 数据库端口
db = first_work_jty
charset = utf8mb4
```
### 接口
1. 登录页面
   - type: GET
   - path1: /
   - path2: /login
2. 登录
   - type: POST
   - path: /login
   - params: username,password
3. 注册
   - type: POST
   - path: /login/register
   - params: username,password
4. 修改密码
   - type: POST
   - path: /login/updatePassword
   - params: username,oldPassword,newPassword
5. 注销账号
   - type: POST
   - path: /login/deleteUser
   - params: id