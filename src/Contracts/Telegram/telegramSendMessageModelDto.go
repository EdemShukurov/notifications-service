package telegram

import (
	common "notifications-service/Contracts/Common"
	"time"
)

type TelegramSendMessageModelDto struct {
	SendAt     *time.Time            `json:"sendAt,omitempty"`
	ChatId     int                   `json:"chatId" binding:"required"`
	Message    string                `json:"message" binding:"required"`
	References []common.ReferenceDto `json:"references,omitempty"`
}
