package data

import (
	common "notifications-service/Domain/Common/Data"
	"time"
)

// NotificationDbo common info
type NotificationDbo struct {
	Id          string
	Type        string
	SendAt      *time.Time
	Status      string
	Details     any
	ProcessedAt *time.Time
	//ExpiresAt   *time.Time
	CreatedAt  *time.Time
	References []common.Reference // TODO: ReferenceDbo
	Errors     []ErrorDbo
}
