package models

import (
	"time"

	"github.com/ZecretBone/ips-rssi-service/apps/constants"
)

type RSSIStatModel struct {
	RSSIInfo   []RSSI                    `bson:"rssi_info"`
	DeviceInfo DeviceInfo                `bson:"device_info"`
	Stage      constants.CollectionStage `bson:"collection_stage"`
	Position   Position                  `bson:"position"`
	Duration   int                       `bson:"duration"`
	StartedAt  time.Time                 `bson:"started_at"`
	EndedAt    time.Time                 `bson:"ended_ata"`
	CreatedAt  time.Time                 `bson:"created_at"`
}
