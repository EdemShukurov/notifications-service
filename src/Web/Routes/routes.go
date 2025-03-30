package routes

import (
	services "notifications-service/Web/Services"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.mongodb.org/mongo-driver/mongo"
)

func ApplicationRouter(router *gin.Engine, mongoClient *mongo.Client) {
	v1 := router.Group("/v1")

	v1.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	//t.me/globy_test_notifications_dev_bot
	token := "7680623680:AAEfNoPXzfBKYUdAT1wylxEeI7MEi6DNGcE"
	TelegramRoutes(v1, services.AddTelegramService(token, mongoClient))
}
