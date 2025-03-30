package telegram

import (
	common "notifications-service/Domain/Common/Data"
)

// Telegram message send request business model
type TelegramNotificationSendRequest struct {
	Request common.RequestBase
	Message TelegramMessage
}

// Telegram message details
type TelegramMessage struct {
	Text string
	Chat TelegramChat
}

// A Telegram TelegramChat indicates the conversation to which the message belongs
type TelegramChat struct {
	Id int64
}
