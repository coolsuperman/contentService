package rpc

import (
	"contentService/internal/rpc/content"
	"contentService/internal/rpc/datamanager"
	"contentService/pkg/config"
)

type Rpc struct {
	config        *config.GConfig
	mysqlInstance *datamanager.MysqlHelper
	redisInstance *datamanager.RedisHelper
}

func NewRpc(conf config.GConfig, mysqlClient *datamanager.MysqlHelper, redisClient *datamanager.RedisHelper) *Rpc {
	return &Rpc{
		config:        &conf,
		mysqlInstance: mysqlClient,
		redisInstance: redisClient,
	}
}

func (r *Rpc) Run() {
	rpcContentServer := content.NewRpcContentServer(r.config.Rpc.RPCPort, r.config.Rpc.ListenIP, r.mysqlInstance, r.redisInstance)
	rpcContentServer.Run()
}
