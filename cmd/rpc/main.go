package api

import (
	"contentService/internal/rpc/content"
	"contentService/pkg/config"
)

func main() {
	rpcServer := content.NewRpcContentServer(config.Config.Rpc.RPCPort, config.Config.Rpc.ListenIP)
	rpcServer.Run()
}
