package mapper

import (
	"github.com/ZecretBone/ips-rssi-service/apps/constants"
	"github.com/ZecretBone/ips-rssi-service/apps/rssi/models"
	"github.com/ZecretBone/ips-rssi-service/apps/rssi/models/request"
	v1 "github.com/ZecretBone/ips-rssi-service/internal/gen/proto/ips/rssi/v1"
	sharedv1 "github.com/ZecretBone/ips-rssi-service/internal/gen/proto/ips/shared/rssi/v1"
)

var (
	FromStatCollectionStateEnumToConstant = map[sharedv1.StatCollectionStage]constants.CollectionStage{
		sharedv1.StatCollectionStage_STAT_COLLECTION_STAGE_SINGLE:   constants.CollectionStageSingle,
		sharedv1.StatCollectionStage_STAT_COLLECTION_STAGE_MULTIPLE: constants.CollectionStageMultiple,
	}
)

func ToRSSIRequestModel(req *v1.CollectDataRequest) request.StatCollectionRequest {
	output := request.StatCollectionRequest{
		Signals: make([]models.RSSI, len(req.Signals)),
		Position: models.Position{
			X: float64(req.Position.X),
			Y: float64(req.Position.Y),
			Z: float64(req.Position.Z),
		},
		Duration:  int(req.Duration),
		StartedAt: req.StartedAt.AsTime(),
		EndedAt:   req.EndedAt.AsTime(),
		CreatedAt: req.CreatedAt.AsTime(),
	}

	for i, v := range req.Signals {
		output.Signals[i] = models.RSSI{
			SSID:        v.Ssid,
			MacAddress:  v.MacAddress,
			Strength:    v.Strength,
			PollingRate: int(v.PollingRate),
			CreatedAt:   v.CreatedAt.AsTime(),
		}
	}

	return output
}

func ToRSSIModel(req *v1.CollectDataRequest) models.RSSIStatModel {
	output := models.RSSIStatModel{
		RSSIInfo: make([]models.RSSI, len(req.Signals)),
		DeviceInfo: models.DeviceInfo{
			DeviceID: req.DeviceInfo.DeviceId,
			Models:   req.DeviceInfo.Models,
		},
		Position: models.Position{
			X: float64(req.Position.X),
			Y: float64(req.Position.Y),
			Z: float64(req.Position.Z),
		},
		Stage:     FromStatCollectionStateEnumToConstant[req.Stage],
		Duration:  int(req.Duration),
		StartedAt: req.StartedAt.AsTime(),
		EndedAt:   req.EndedAt.AsTime(),
		CreatedAt: req.CreatedAt.AsTime(),
	}

	for i, v := range req.Signals {
		output.RSSIInfo[i] = models.RSSI{
			SSID:        v.Ssid,
			MacAddress:  v.MacAddress,
			Strength:    v.Strength,
			PollingRate: int(v.PollingRate),
			CreatedAt:   v.CreatedAt.AsTime(),
		}
	}

	return output
}
