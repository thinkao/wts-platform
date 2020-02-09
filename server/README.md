# server

### note

在`.bashrc`添加环境变量,`sourse .bashrc`使其生效

```
export GO111MODULE=on
export GOPROXY=https://goproxy.cn
```


### run project

修改mysql的账号密码和添加数据库与conf下的配置文件保持一致
```
bee run
或
go run main.go
```

在运行前可以先`go build main.go`下载完所有依赖

