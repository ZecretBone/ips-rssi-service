package config

import (
	wireminio "github.com/RyuChk/wire-provider/minio"
	wiremongo "github.com/RyuChk/wire-provider/mongodb"
	"github.com/kelseyhightower/envconfig"
)

func ProvideMongoxConfig() wiremongo.Config {
	return provideConfig(wiremongo.Config{})
}

func ProvideMinioXConfig() wireminio.Config {
	return provideConfig(wireminio.Config{})
}

func ProvideCacheConfig() CacheConfig {
	return provideConfig(CacheConfig{})
}

func ProvideStatCollectionServiceConfig() StatCollectionServiceConfig {
	return provideConfig(StatCollectionServiceConfig{})
}

func ProvideApCollectionServiceConfig() ApCollectionServiceConfig {
	return provideConfig(ApCollectionServiceConfig{})
}

func provideConfig[T any](cfg T) T {
	envconfig.MustProcess("", &cfg)
	return cfg
}
