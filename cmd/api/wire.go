//go:build wireinject
// +build wireinject

package api

import (
	"contentService/internal/api"
	"contentService/pkg/config"
	"github.com/google/wire"
)

func wireApi() (*api.Api, func(), error) {
	panic(wire.Build(
		config.NewGConfig,
		api.ProviderApiSet,
		api.NewApi,
	))
}
