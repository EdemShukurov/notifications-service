package telegram

import (
	"context"
	"encoding/json"
	"fmt"
	common "notifications-service/Domain/Common/Data"
	domain "notifications-service/Domain/Telegram/Data"
	sender "notifications-service/Domain/Telegram/Infrastructure/Sending"
	notificationsRepo "notifications-service/Repositories/Contracts/Notifications"
	notificationDboContracts "notifications-service/Repositories/Contracts/Notifications/Data"
	queue "notifications-service/Repositories/Contracts/Queue"
	queueData "notifications-service/Repositories/Contracts/Queue/Data"
	"time"

	gocron "github.com/go-co-op/gocron/v2"
	"github.com/rs/xid"
)

// Service is the interface that provides Telegram notifications methods
type ITelegramService interface {
	SendMessage(ctx context.Context, model *domain.TelegramNotificationSendRequest) (*common.Response, error)
}

type TelegramService struct {
	// readonly fields
	sender     sender.ITelegramSender
	queue      queue.INotificationRequestsQueue
	scheduler  gocron.Scheduler
	repository notificationsRepo.INotificationsRepository
}

func New(sender sender.ITelegramSender, queue queue.INotificationRequestsQueue, repository notificationsRepo.INotificationsRepository) ITelegramService {
	scheduler, err := gocron.NewScheduler()
	if err != nil {
		// TODO: handle error
	}

	service := &TelegramService{
		sender:     sender,
		queue:      queue,
		scheduler:  scheduler,
		repository: repository,
	}

	service.Start()
	return service
}

func (s *TelegramService) Start() {
	// add a job to the scheduler
	j, err := s.scheduler.NewJob(
		gocron.DurationJob(
			10*time.Second,
		),
		gocron.NewTask(func(ctx context.Context) {
			s.RunSendingJob(ctx)
		}),
	)

	if err != nil {
		// handle error
	}

	// each job has a unique id
	fmt.Println(j.ID())

	// start the scheduler
	s.scheduler.Start()
}

func (s *TelegramService) Stop() {
	// stop the scheduler
	s.scheduler.Shutdown()
}

func (s *TelegramService) RunSendingJob(ctx context.Context) {
	for {
		dbo, err := s.queue.Peek(ctx, "telegram")
		if err != nil {
			fmt.Errorf("Queue peek operation was failed. Error: %w", err)
			break
		}

		if dbo == nil {
			fmt.Println("No requests in the queue to process at this moment")
			break
		}

		s.ProcessRequest(ctx, dbo)
	}
}

func (s *TelegramService) SendMessage(ctx context.Context, model *domain.TelegramNotificationSendRequest) (*common.Response, error) {
	// Marshal the Person object into a JSON string
	requestJson, err := json.Marshal(model)
	if err != nil {
		return nil, err
	}

	queueItem := queueData.NotificationRequestQueueItemDbo{
		Id:          xid.New().String(),
		CreatedAt:   time.Now().UTC(),
		Status:      "pending",
		Type:        "telegram",
		SendAt:      model.Request.SendAt,
		RequestJson: string(requestJson),
	}

	err = s.queue.Enqueue(ctx, queueItem)
	if err != nil {
		return nil, err
	}

	response := &common.Response{
		Id:        "test",
		CreatedAt: time.Now().UTC(),
	}

	return response, nil
}

func (s *TelegramService) ProcessRequest(ctx context.Context, dbo *queueData.NotificationRequestQueueItemDbo) error {

	var request domain.TelegramNotificationSendRequest
	err := json.Unmarshal([]byte(dbo.RequestJson), &request)
	if err != nil {
		fmt.Errorf("Unmarshall operation was failed. Error: %w", err)
		return err
	}

	notification := notificationDboContracts.NotificationDbo{
		Id:          dbo.Id,
		Type:        "telegram",
		SendAt:      &dbo.SendAt,
		Details:     request.Message,
		ProcessedAt: nil,
		Status:      "running",
		//ExpiresAt:   nil,
		CreatedAt:  timePtr(time.Now().UTC()),
		References: request.Request.References,
		Errors:     nil,
	}

	err = s.sender.SendMessage(ctx, request.Message)
	if err != nil {
		notification.Status = "failed"
		notification.Errors = make([]notificationDboContracts.ErrorDbo, 1)
		notification.Errors = append(notification.Errors, notificationDboContracts.ErrorDbo{
			Type:    "telegram_bot",
			Message: err.Error(),
		})

		fmt.Errorf("Unable to send message. Error: %w", err)

		s.repository.CreateOrUpdate(ctx, notification)

		return err
	}

	notification.ProcessedAt = timePtr(time.Now().UTC())
	notification.Status = "succeeded"

	s.repository.CreateOrUpdate(ctx, notification)

	s.queue.Delete(ctx, dbo.Id)

	return nil
}

func timePtr(t time.Time) *time.Time {
	return &t
}
