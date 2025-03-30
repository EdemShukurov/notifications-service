package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	services "notifications-service/Web/Services"
	routes "notifications-service/Web/routes"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	mongoClient, err := services.AddMongoDbService("mongodb://localhost:27017")
	if err != nil {
		panic(fmt.Errorf("failed to initialize DB: %w", err))
	}

	router := gin.Default()
	//router.Use(cors.Default())
	// DB, err := repository.InitDB()
	// if err != nil {
	// 	panic(fmt.Errorf("failed to initialize DB: %w", err))
	// }

	// router.Use(middlewares.ErrorHandler())
	// router.Use(middlewares.GinBodyLogMiddleware)
	// router.Use(middlewares.CommonHeaders)
	// routes.ApplicationRouter(router, DB)
	routes.ApplicationRouter(router, mongoClient)

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "3000"
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	s := &http.Server{
		Addr:           ":" + port,
		Handler:        router,
		ReadTimeout:    18000 * time.Second,
		WriteTimeout:   18000 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	fmt.Printf("Service is running http://localhost:%s\n", port)
	if err := s.ListenAndServe(); err != nil {
		panic(strings.ToLower(err.Error()))
	}

	router.Run(":" + port)
}
