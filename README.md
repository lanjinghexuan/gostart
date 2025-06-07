# gostart
初学者的go项目

项目启动文件分别在
api/cmd/main.go
server/cmd/main.go

项目启动需要进入目录下运行go run main.go
否则会导致检查配置文件错误

配置文件为
/serve/inits/config/dev.yaml
内容需要自行编写
下面为示例

Mysql:
Host: 127.0.0.1
Port: 3306
User: root
Pass: root
DB: db1
Redis:
Host: 127.0.0.1
Port: 6379

