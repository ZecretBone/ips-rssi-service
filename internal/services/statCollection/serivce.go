package statcollection

import (
	"context"
	"math"

	"github.com/ZecretBone/ips-rssi-service/apps/rssi/models"
	"github.com/ZecretBone/ips-rssi-service/internal/config"
	statcollectionrepo "github.com/ZecretBone/ips-rssi-service/internal/repository/mongodb/statCollectionRepo"
)

type Service interface {
	AddSignalStatToDB(ctx context.Context, stat models.RSSIStatModel) error
}

type StatCollectionService struct {
	statCollectionRepo statcollectionrepo.Repository
	cfg                config.StatCollectionServiceConfig
}

func ProvideStatCollectionService(statCollectionRepo statcollectionrepo.Repository, cfg config.StatCollectionServiceConfig) Service {
	return &StatCollectionService{
		statCollectionRepo: statCollectionRepo,
		cfg:                cfg,
	}
}

func (s *StatCollectionService) AddSignalStatToDB(ctx context.Context, stat models.RSSIStatModel) error {
	for _, v := range stat.RSSIInfo {
		v.AverageStrenth = float32(AverateStrength(v.Strength))
	}

	if err := s.statCollectionRepo.InsertOne(ctx, stat); err != nil {
		return err
	}
	return nil
}

func AverateStrength(strength_list []float64) float64 {
	var t float64
	for _, v := range strength_list {
		t += v
	}

	return math.Ceil(float64(t) / float64(len(strength_list)))
}
