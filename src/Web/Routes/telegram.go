package routes

import (
	telegram "notifications-service/Web/Controllers/Telegram"

	"github.com/gin-gonic/gin"
)

func TelegramRoutes(router *gin.RouterGroup, controller telegram.ITelegramController) {
	rg := router.Group("/telegram")
	// rg.Use(middlewares.AuthJWTMiddleware())
	// {
	// 	rg.POST("/sendMessage", controller.SendMessage)
	// }

	rg.POST("/sendMessage", controller.SendMessage)
}
