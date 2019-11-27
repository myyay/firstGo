### 【配置】

#### 下载安装包  https://golang.google.cn/  

#### 环境变量修改 
GOROOT 是安装目录不用改

GOPATH 改为工作目录，以后下来的包都会放到工作目录。 比如D:\GoProjects 可以在目录下建立以下3个目录 src pkg 和 bin

PATH 要加上GOPATH/bin


#### 安装下载gopm

执行  `go get -u -v github.com/gpmgo/gopm`

#### 安装 golang tools

执行  `gopm get -u -v -g golang.org/x/tools/cmd/goimports`

#### 表格驱动测试

`basic/basic/triangle_test.go`

`container/norepeating/nonrepeating_test.go.go`

源码中的注释写了一些命令的使用方法 包括普通测试 和 性能测试

需要下载 http://graphviz.gitlab.io/


#### 文档 go doc

`queue/queue.go`

需要安装go tools包 https://github.com/golang/tools
1.

翻墙  go get -v -u golang.org/x/tools/cmd/godoc

不翻墙 
git clone https://github.com/golang/tools $GOPATH/src/golang.org/x/tools 
git clone https://github.com/golang/tools %GOPATH%/src/golang.org/x/tools 

或者用以上的gopm也可以  gopm get -u -v -g golang.org/x/tools/cmd/godoc

2.

go build $GOPATH/src/golang.org/x/tools/cmd/godocg
mv $GOPATH/src/golang.org/x/tools/cmd/godoc/godoc $GOPATH/bin/




