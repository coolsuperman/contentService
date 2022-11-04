package datamanager

import (
	"github.com/google/wire"
)

var ProviderRpcSet = wire.NewSet(NewMysqlHelper, NewRedisHelper)
