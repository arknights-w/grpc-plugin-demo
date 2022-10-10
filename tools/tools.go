package tools

// 这个类就是用来对外暴露的

import (
	"context"
	"fmt"
	"log"
	"os/exec"
	"time"

	"github.com/hashicorp/go-plugin"

	"grpc-plugin/config"
	"grpc-plugin/plugins"
)

// monitor 内部函数，监控数据更替，create中调用
func monitor(ctx context.Context,cli *plugin.Client) {
	ticker := time.NewTicker(5 * time.Second)
Loop:
	for {
		select {
		case <-ticker.C:
			//do something
			if false {
				break Loop
			}
		case <-ctx.Done():
			//手动的
			fmt.Println("---------------------------")
			cli.Kill()
			fmt.Println("关闭了服务器")
			fmt.Println("---------------------------")
			break Loop
		}
	}
}

// 传入配置文件信息，返回 client 和 close 函数
func CreateClient(c_path string, c_name string, c_type string) (plugin.ClientProtocol, context.CancelFunc) {


	// 这里包了两层
	// 里面一层是用来打开服务器的(这里我们总是让插件二进制程序名为plugins)
	// 外面一层是用来日志交互的
	cmd := exec.Command("./plugins",
		"-path", c_path,
		"-name", c_name,
		"-type", c_type)
	cmd = exec.Command("sh",
		"-c", cmd.String())
	
	conf := config.GetConfig(c_path, c_name, c_type)
	plugins_get := plugins.GetPlugins(conf)
	// 这个 client 可用于管理 插件子进程 的生命周期
	// HandshakeConfig	通讯密钥
	// Plugins			插件集，选择需要的插件进行加载
	// Cmd				rpc命令
	// AllowedProtocols	允许协议
	client := plugin.NewClient(&plugin.ClientConfig{
		HandshakeConfig: plugins.Handshake_sm,
		Plugins:         plugins_get,
		Cmd:             cmd,
		AllowedProtocols: []plugin.Protocol{
			plugin.ProtocolNetRPC, plugin.ProtocolGRPC},
	})

	ctx, cancel := context.WithCancel(context.Background())
	go monitor(ctx,client)

	// 启动服务器端,同时返回客户端
	// 注意服务器是在客户端中启起来的
	// client-bootstrap
	cp, err := client.Client()
	if err != nil {
		log.Println("clientProtocol create failed")
	}

	return cp, cancel
}
