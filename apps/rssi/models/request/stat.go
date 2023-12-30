package request

import (
	"time"

	"github.com/ZecretBone/ips-rssi-service/apps/rssi/models"
)

type StatCollectionRequest struct {
	Signals   []models.RSSI   `json:"signals"`
	Position  models.Position `json:"position"`
	Duration  int             `json:"duration"` //capture duration in second
	StartedAt time.Time       `json:"started_at"`
	EndedAt   time.Time       `json:"ended_at"`
	CreatedAt time.Time       `json:"created_at"`
}
