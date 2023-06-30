# first_work_jty
姜同远的go作业

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
db = bubble
charset = utf8mb4
```
### 编译
```bash
go build
```
### 执行
Mac/Unix：
```bash
./first_work_jty
```
Windows:
```bash
first_work_jty.exe
```
### 接口
1. 登录
   - type: POST
   - path: /login
   - params: username,password
2. 注册
    - type: POST
    - path: /login/register
    - params: username,password
3. 修改密码
    - type: POST
    - path: /login/updatePassword
    - params: username,oldPassword,newPassword
4. 注销账号
    - type: POST
    - path: /login/deleteUser
    - params: id