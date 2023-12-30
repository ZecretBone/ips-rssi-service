package di

import (
	wireminio "github.com/RyuChk/wire-provider/minio"
	wiremongo "github.com/RyuChk/wire-provider/mongodb"
	"github.com/ZecretBone/ips-rssi-service/internal/config"
	"github.com/ZecretBone/ips-rssi-service/internal/repository/cache"
	"github.com/ZecretBone/ips-rssi-service/internal/repository/minio"
	"github.com/ZecretBone/ips-rssi-service/internal/repository/mongodb"
	statcollectionrepo "github.com/ZecretBone/ips-rssi-service/internal/repository/mongodb/statCollectionRepo"
	statcollection "github.com/ZecretBone/ips-rssi-service/internal/services/statCollection"
	"github.com/google/wire"
)

var DatabaseSet = wire.NewSet(
	minio.ProvideMinioService,
	mongodb.ProvideMongoDBService,
	cache.ProvideCacheService,
)

var ProviderSet = wire.NewSet(
	statcollectionrepo.ProvideStatCollectionRepo,
	statcollection.ProvideStatCollectionService,
)

var ConfigSet = wire.NewSet(
	config.ProvideMongoxConfig,
	config.ProvideMinioXConfig,
	config.ProvideCacheConfig,
	config.ProvideStatCollectionServiceConfig,
)

type Locator struct {
	MongoDBConn  wiremongo.Connection
	MinioXConn   wireminio.Connection
	RSSIService  statcollection.Service
	CacheService cache.Service
}
