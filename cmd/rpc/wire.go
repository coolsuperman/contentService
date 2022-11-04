//go:build wireinject
// +build wireinject

package rpc

import (
	"contentService/internal/rpc"
	"contentService/internal/rpc/datamanager"
	"contentService/pkg/config"
	"github.com/google/wire"
)

func wireRpc() (*rpc.Rpc, func(), error) {
	panic(wire.Build(
		config.NewGConfig,
		datamanager.ProviderDataManagerSet,
		rpc.NewRpc,
	))
}
