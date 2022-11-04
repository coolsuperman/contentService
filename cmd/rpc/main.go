package rpc

import (
	"contentService/internal/rpc/content"
	"contentService/pkg/config"
	"github.com/spf13/cobra"
)

func main() {
	rpcServer := content.NewRpcContentServer(config.Config.Rpc.RPCPort, config.Config.Rpc.ListenIP)
	rpcServer.Run()
}

var RpcCmd = &cobra.Command{
	Use:   "rpc",
	Short: "文章系统Rpc服务",
	Long:  "练习做文章系统Rpc服务",
	Run: func(cmd *cobra.Command, args []string) {
		rpcServer := content.NewRpcContentServer(config.Config.Rpc.RPCPort, config.Config.Rpc.ListenIP)
		rpcServer.Run()
	},
}
