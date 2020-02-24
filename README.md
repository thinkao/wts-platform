# wts-platform

> wts_web


``` bash
# install dependencies
npm install

# serve with hot reload at localhost:8080
npm run dev

# build for production with minification
npm run build

# build for production and view the bundle analyzer report
npm run build --report

# run unit tests
npm run unit

# run e2e tests
npm run e2e

# run all tests
npm test
```




> server


在`.bashrc`添加环境变量,`sourse .bashrc`使其生效

```
export GO111MODULE=on
export GOPROXY=https://goproxy.cn
```


修改mysql的账号密码和添加数据库与conf下的配置文件保持一致
```
bee run
或
go run main.go
```

在运行前可以先`go build main.go`下载完所有依赖



