package controllers

import (
	"fmt"
	"net/http"
	commonDto "notifications-service/Contracts/Common"
	dto "notifications-service/Contracts/Telegram"
	commonDomain "notifications-service/Domain/Common/Data"
	service "notifications-service/Domain/Telegram"
	domain "notifications-service/Domain/Telegram/Data"
	controllers "notifications-service/Web/Controllers"
	"time"

	"github.com/gin-gonic/gin"
)

type ITelegramController interface {
	SendMessage(ctx *gin.Context)
	//GetById(ctx *gin.Context)
}

type TelegramController struct {
	service service.ITelegramService
}

func New(service service.ITelegramService) ITelegramController {
	return &TelegramController{service}
}

// SendMessage implements ITelegramController.
// @Summary      Show an account
// @Description  get string by ID
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Account ID"
// @Success      200  {object}  model.Account
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /accounts/{id} [get]
func (c *TelegramController) SendMessage(ctx *gin.Context) {

	var dto dto.TelegramSendMessageModelDto
	if err := controllers.BindJSON(ctx, &dto); err != nil {
		//appError := domainErrors.NewAppError(err, domainErrors.ValidationError)
		_ = ctx.Error(err)
		return
	}

	fmt.Println("Request to send message was received: %v", dto)

	response, err := c.service.SendMessage(ctx, toDomainModel(&dto))
	if err != nil {
		fmt.Errorf("Unabled to send the message. Error: %w", err)
		_ = ctx.Error(err)
		return
	}

	fmt.Println("Message was sent. Response: %v", response)

	//response := domainToResponseMapper(response)
	ctx.JSON(http.StatusOK, response)

}

func toDomainModel(req *dto.TelegramSendMessageModelDto) *domain.TelegramNotificationSendRequest {
	var sendAt time.Time

	if req.SendAt == nil || req.SendAt.IsZero() {
		sendAt = time.Now().UTC()
	} else {
		sendAt = *req.SendAt
	}

	return &domain.TelegramNotificationSendRequest{
		Request: commonDomain.RequestBase{
			References: mapSlice(req.References, func(dto commonDto.ReferenceDto) commonDomain.Reference {
				return commonDomain.Reference{
					Id:   dto.Id,
					Type: dto.Type,
				}
			}),
			SendAt: sendAt,
		},
		Message: domain.TelegramMessage{
			Text: req.Message,
			Chat: domain.TelegramChat{
				Id: int64(req.ChatId),
			},
		},
	}
}

func mapSlice[T any, M any](a []T, f func(T) M) []M {
	n := make([]M, len(a))
	for i, e := range a {
		n[i] = f(e)
	}
	return n
}
