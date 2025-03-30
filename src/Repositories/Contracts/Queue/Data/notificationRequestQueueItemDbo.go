package repositories

import (
	"time"
)

// The wrapper over Request that is stored in a processing queue
type NotificationRequestQueueItemDbo struct {
	Id          string    `bson:"_id"`
	Status      string    `bson:"Status"` // pending | processing
	Type        string    `bson:"Type"`
	CreatedAt   time.Time `bson:"CreatedAt"`
	SendAt      time.Time `bson:"SendAt"`
	RequestJson string    `bson:"Request"`
}
