学习资源：https://www.lanqiao.cn/courses/834/learning/?id=2991

笔记：

0、目录及环境配置

windows +　docker(golang:latest镜像) + vscode

获取镜像：
```
docker pull golang:latest
```

创建容器
```
docker run -it -v d:\Workspaces\Go\:/go/src -p 3000:3000 --name go_learn golang:latest /bin/bash
```

使用vscode，可以直接连接远程资源管理器进入到docker 容器。选择刚刚创建好的`go_learn`容器，并进入。

查看go环境变量
```
#go env

GO111MODULE=""
...
GOPATH="/go"
GOROOT="/usr/local/go"
GOVERSION="go1.16.5"
...
```

代码目录：`/go/src/go_learn`

目录结构：
```
/go/src/go_learn
- data
    - data.json
- matchers 
    - rss.go
- search 
    - default.go
    - feed.go
    - match.go
    - search.go
- main.go
```

注意：
main.go 中引入路径为：
```go
import (
    "log"
    "os"

    _ "go_learn/matchers"
    "go_learn/search"
)
```

遇到的问题：

1、build的报错信息是：
![图片描述](https://dn-simplecloud.shiyanlou.com/courses/uid168132-20210712-1626101267503)

解决：
执行命令
```
go mod init

go mod tidy
```
然后执行`go build main.go`就可以了

2、这个应用的代码使用了 4 个文件夹，按字母顺序列出。文件夹 data 中有一个 JSON 文档，其内容是程序要拉取和处理的数据源。文件夹 matchers 中包含程序里用于支持搜索不同数据源的代码。目前程序只完成了支持处理 RSS 类型的数据源的匹配器。文件夹 search 中包含使用不同匹配器进行搜索的业务逻辑。最后，父级文件夹 sample 中有个 main.go 文件，这是整个程序的入口。

3.1 main包

在 Go 语言中，所有变量都被初始化为其零值。对于数值类型，零值是 0；对于字符串类型，零值是空字符串；对于布尔类型，零值是 false；对于指针，零值是 nil。对于引用类型来说，所引用的底层数据结构会被初始化为对应的零值。但是被声明为其零值的引用类型的变量，会返回 nil 作为其值。

3.2 search包

这个程序里的匹配器，是指包含特定信息、用于处理某类数据源的实例。在这个示例程序中有两个匹配器。框架本身实现了一个无法获取任何信息的默认匹配器，而在 matchers 包里实现了 RSS 匹配器。RSS 匹配器知道如何获取、读入并查找 RSS 数据源。随后我们会扩展这个程序，加入能读取 JSON 文档或 CSV 文件的匹配器。

3.2.1 search.go

(1)如果需要声明初始值为零值的变量，应该使用var关键字声明变量；如果提供确切的非零值初始化变量或者使用函数返回值创建变量，应该使用简化变量声明运算符。

(2)在 Go 语言中，通道（channel）和映射（map）与切片（slice）一样，也是引用类型，不过通道本身实现的是一组带类型的值，这组值用于在goroutine之间传递数据。通道内置同步机制，从而保证通信安全。

(3)在 Go 语言中，如果 main 函数返回，整个程序也就终止了。Go 程序终止时，还会关闭所有之前启动且还在运行的 goroutine。写并发程序的时候，最佳做法是，在 main 函数返回前，清理并终止所有之前启动的 goroutine。编写启动和终止时的状态都很清晰的程序，有助减少 bug，防止资源异常。

(4)使用sync包的WaitGroup跟踪所有启动的 goroutine。非常推荐使用WaitGroup来跟踪 goroutine 的工作是否完成。WaitGroup是一个计数信号量，我们可以利用它来统计所有的 goroutine 是不是都完成了工作

3.2.2 feed.go

关键字defer会安排随后的函数调用在函数返回时才执行。在使用完文件后，需要主动关闭文件。使用关键字defer来安排调用 Close 方法，可以保证这个函数一定会被调用。哪怕函数意外崩溃终止，也能保证关键字defer安排调用的函数会被执行。关键字defer可以缩短打开文件和关闭文件之间间隔的代码行数，有助提高代码可读性，减少错误。

3.2.3 match.go/default.go

命名接口的时候，也需要遵守 Go 语言的命名惯例。如果接口类型只包含一个方法，那么这个类型的名字以er结尾。我们的例子里就是这么做的，所以这个接口的名字叫作 Matcher。如果接口类型内部声明了多个方法，其名字需要与其行为关联。


