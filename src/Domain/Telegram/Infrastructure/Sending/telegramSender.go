package telegram

import (
	"context"
	"fmt"
	data "notifications-service/Domain/Telegram/Data"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// Service is the interface that provides Telegram sender methods
type ITelegramSender interface {
	SendMessage(ctx context.Context, message data.TelegramMessage) error
}

type TelegramSender struct {
	client *tgbotapi.BotAPI
}

// New returns a new instance of a Telegram sender service.
// For more information about telegram api token:
//
//	-> https://pkg.go.dev/github.com/go-telegram-bot-api/telegram-bot-api#NewBotAPI
func New(apiToken string) (ITelegramSender, error) {
	client, err := tgbotapi.NewBotAPI(apiToken)
	if err != nil {
		return nil, err
	}

	t := &TelegramSender{
		client: client,
	}

	return t, nil
}

func (t *TelegramSender) SendMessage(ctx context.Context, message data.TelegramMessage) error {
	text := message.Text
	chatId := message.Chat.Id

	msg := tgbotapi.NewMessage(chatId, text)
	//msg.ParseMode = parseMode

	select {
	case <-ctx.Done():
		return ctx.Err()
	default:

		_, err := t.client.Send(msg)
		if err != nil {
			err.Error()
			return fmt.Errorf("send message to chat %d: %w", chatId, err)
		}
	}

	return nil
}
