# WebDav 多文件上传
支持各平台以及Drone 
## Parameters
- -u URL 服务地址
- -a User 用户命
- -p Password 密码
- -f Files 要上传的文件,多个以逗号隔开
- -r Remote Dir 远程目录
## Windows
```shell
# 构建
go build -o updav.exe 
# 使用
.\updav.exe -u http://192.168.8.6:8081 -a admin -p password -f a.txt,b.txt -r /project/test
```
## Linux
```shell
# 构建
go build -o updav 
# 使用
.\updav -u http://192.168.8.6:8081 -a admin -p password -f a.txt,b.txt -r /project/test
```
