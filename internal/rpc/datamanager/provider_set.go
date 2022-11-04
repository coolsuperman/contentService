package datamanager

import (
	"github.com/google/wire"
)

var ProviderDataManagerSet = wire.NewSet(NewMysqlHelper, NewRedisHelper)
