package rssicollection

import (
	"context"

	//"github.com/ZecretBone/ips-rssi-service/apps/rssi/models"
	"github.com/ZecretBone/ips-rssi-service/internal/config"
	rssiv1 "github.com/ZecretBone/ips-rssi-service/internal/gen/proto/ips/rssi/v1"
	apcollectionrepo "github.com/ZecretBone/ips-rssi-service/internal/repository/mongodb/apCollectionRepo"
)

type Service interface {
	IsExpectedApExisted(ctx context.Context, req *rssiv1.GetCoordinateRequest) (bool, error)
	RegisterNewAp(ctx context.Context, req *rssiv1.RegisterApRequest) error
}

type RssiCollectionService struct {
	apCollectionRepo apcollectionrepo.Repository
	cfg              config.ApCollectionServiceConfig
}

func ProvideRssiCollectionService(apCollectionRepo apcollectionrepo.Repository, cfg config.ApCollectionServiceConfig) Service {
	return &RssiCollectionService{
		apCollectionRepo: apCollectionRepo,
		cfg:              cfg,
	}
}

func (s *RssiCollectionService) IsExpectedApExisted(ctx context.Context, req *rssiv1.GetCoordinateRequest) (bool, error) {
	found, err := s.apCollectionRepo.IsExpectedApExisted(ctx, req)
	if err != nil {
		return false, err
	}
	if !found {
		return false, nil
	}
	return true, nil
}

func (s *RssiCollectionService) RegisterNewAp(ctx context.Context, req *rssiv1.RegisterApRequest) error {
	if err := s.apCollectionRepo.InsertOne(ctx, req); err != nil {
		return err
	}
	return nil
}
