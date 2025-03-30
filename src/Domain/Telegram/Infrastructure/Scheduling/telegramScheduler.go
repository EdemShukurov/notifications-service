// package telegram

// import (
// 	"context"
// 	"time"

// 	"github.com/robfig/cron"
// 	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
// )

// // Service is the interface that provides Telegram sender methods
// type ITelegramScheduler interface {
// 	SendMessage(ctx context.Context, message string, chatId int64) error
// }

// type TelegramScheduler struct {
// 	client *tgbotapi.BotAPI
// }

// // New returns a new instance of a Telegram sender service.
// // For more information about telegram api token:
// //
// //	-> https://pkg.go.dev/github.com/go-telegram-bot-api/telegram-bot-api#NewBotAPI
// func New(apiToken string) (error) {
// 	c := cron.New()
// 	if err != nil {
// 		return err
// 	}

// 	// add a job to the scheduler
// 	_, err = s.NewJob(
// 		gocron.DurationJob(
// 			10*time.Second,
// 		),
// 		gocron.NewTask(
// 			func(db *database.OChainDatabase) {
// 				CheckAndHandlePortalUpdate(cfg, db)
// 			},
// 			db,
// 		),
// 	)

// 	if err != nil {
// 		return OChainScheduler{}, err
// 	}

// 	return OChainScheduler{
// 		Scheduler: s,
// 		db:        db,
// 		cfg:       cfg,
// 	}, nil

// func (t *TelegramScheduler) ScheduleNextFireTime(ctx context.Context, nextTime time.Time) error {

// 	return nil
// }
