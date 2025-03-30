package queue

import (
	"context"
	data "notifications-service/Repositories/Contracts/Queue/Data"
)

// Queue definition for all notifications
type INotificationRequestsQueue interface {
	Enqueue(ctx context.Context, dbo data.NotificationRequestQueueItemDbo) error
	Peek(ctx context.Context, t string) (*data.NotificationRequestQueueItemDbo, error)
	Delete(ctx context.Context, id string) error
}
