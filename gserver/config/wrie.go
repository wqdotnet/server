// +build wireinject

package config

import (
	"github.com/google/wire"
)

func CreateApp(cfgname string) error {
	panic(wire.Build(providerSet))
}

var providerSet = wire.NewSet(
	ProviderSet,
)
