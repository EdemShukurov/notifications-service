package notifications

import (
	"context"
	data "notifications-service/Repositories/Contracts/Notifications/Data"
)

// Definition for storing notifications
type INotificationsRepository interface {
	CreateOrUpdate(ctx context.Context, dbo data.NotificationDbo) error
}
