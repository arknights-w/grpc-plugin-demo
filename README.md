# TextMessage Example

这是我对 go-plugin 仓库中, grpc 例子的极简化,并对项目目录进行了工程向优化, 当然你可以去原网站去观摩原有的代码, 我去掉了其中与 grpc 毫不相关的代码, 比如 netrpc, python-rpc 等, 让代码的可读性增加, 更有利于去了解整个项目的结构

这里是原例的地址
```
https://github.com/hashicorp/go-plugin/tree/master/examples/grpc
```


```sh
# This builds the plugin written in Go,
# tencent is filename whitch you can change
$ go build -o ./bootstrap/client/tencent ./bootstrap/server/
# or
$ make build

# to run this plugin, 
# you just need to setup bootstrap/client/main
# you can do this
$ cd bootstrap/client/
$ go run .
# more wisely, try this
$ make run
```

## Updating the Protocol

For Go:
这里我做了修改, 原因是原版本的 protoc 版本过老, 可以看见Version2, 当前最新版 Version7

```sh
$ protoc -I proto/ proto/text_message.proto --go_out=proto/
$ protoc -I proto/ proto/text_message.proto --go-grpc_out=proto/

# otherwise,you can use this command
$ make protoc
```

## 一些其他需要注意的东西

一个插件的服务器端, 可以承载多个插件

一个插件的客户端, 可以打开不同的服务器端

服务器端有 Plugins 配置项(实际上客户端也有), 你可以在其中配多个 Plugin

客户端配置有 cmd 可以选择不同的可执行文件, 每一个可执行文件都是一个独立的服务器

Plugins 作为公共的配置参数,可以将数据存进数据库,每次调用时动态读取

当然在 短信插件中, 一个服务器只有一个 Plugin, 原因是减少代码增量后,对原有代码的重新编译

假设我们现在用的是 腾讯的短信服务，我们将其转为 阿里的短信服务，那么我们需要改动的代码仅限于 service/中