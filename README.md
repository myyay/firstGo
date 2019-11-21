#### 配置

##### 下载安装包  https://golang.google.cn/  

##### 环境变量修改 
GOROOT 是安装目录不用改

GOPATH 改为工作目录，以后下来的包都会放到工作目录。 比如D:\GoProjects 可以在目录下建立以下3个目录 src pkg 和 bin

PATH 要加上GOPATH/bin


##### 安装下载gopm

执行  `go get -u -v github.com/gpmgo/gopm`

##### 安装 golang tools

执行  `gopm get -u -v -g golang.org/x/tools/cmd/goimports`

##### 表格驱动测试

`basic/basic/triangle_test.go`

`container/norepeating/nonrepeating_test.go.go`

源码中的注释写了一些命令的使用方法 包括普通测试 和 性能测试



