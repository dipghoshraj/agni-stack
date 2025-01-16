package cosmicmodel

import "time"

type LokaSthiti string

const (
	LokaSthitiActive   LokaSthiti = "active"
	LokaSthitiFailed   LokaSthiti = "failed"
	LokaSthitiPending  LokaSthiti = "pending"
	LokaSthitiInactive LokaSthiti = "inactive"
)

type LokaNode struct {
	ID            string     `json:"id" gorm:"primaryKey"`
	Status        LokaSthiti `json:"status"`
	Capacity      int64      `json:"capacity"`  // in MB
	UsedSpace     int64      `json:"usedSpace"` // in MB
	LastHeartbeat time.Time  `json:"lastHeartbeat"`
	CreatedAt     time.Time  `json:"createdAt"`
	UpdatedAt     time.Time  `json:"updatedAt"`
	VolumeName    string     `json:"volumeName"`
	Port          string     `json:"port"`
	Host          string     `json:"host"`
}
