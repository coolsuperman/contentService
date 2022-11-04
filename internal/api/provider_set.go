package api

import (
	"contentService/internal/api/conn"
	"contentService/internal/api/content"
	"github.com/google/wire"
)

var ProviderApiSet = wire.NewSet(NewApi, conn.NewRpcConnClient, content.NewApiContent)
