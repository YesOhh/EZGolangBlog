EZGolangBlog 博客
=
使用 golang 编写的简单博客，只要把 markdown 文档放在配置文件中指定的地址即可。

## 使用须知
* 配置 conf.toml

    * markdown 文章存放路径
    * 需要展示的具体分类和不需要区分的默认分类
    * 文件名和分类间的分隔符
    * 运行的 ip 和端口，默认运行在 127.0.0.1:8080
    * 日志存放的文件夹地址

* 运行

    * 安装 go 并配置好环境变量
    * go run main.go
    > 因为墙的原因，可能无法更新 go mod，可以配置使用国内镜像源  
    > 配置方式见 https://goproxy.cn