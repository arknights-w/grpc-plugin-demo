# TextMessage Example

这是多插件管理方案

对原有grpc-plugin的改写在master分支中

sendMsg分支则将kv案例改为了短信案例

这里是原例的地址
```
https://github.com/hashicorp/go-plugin/tree/master/examples/grpc
```


```sh
# 这个命令用于编译插件为一个二进制文件
# 二进制文件存放于test中，因为测试需要
$ make build
```

## Updating the Protocol

For Go:
这里我做了修改, 原因是原版本的 protoc 版本过老, 可以看见Version2, 当前最新版 Version7

```sh
$ make protoc
```

## 一些其他需要注意的东西

这里说一下项目结构
```
/bootstrap  插件main函数入口，是插件启动器
            二进制文件就是由他编译得到

/config     这包是用于处理配置文件的

/plugins    所有的插件都放在这里
    
    /plugin1    这是其中一个插件

        /proto      protobuf文件位置

        /service    业务接口实现

/test       一些测试

/tools      对外暴露的函数，比如对插件服务器的创建、销毁插件等
```