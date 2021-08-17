package main

import (
	"context"
	"fmt"
	"game-api/config"
	"game-api/handler"
	"game-api/middleware"
	"game-api/repository/postgres"
	"game-api/repository/redis"
	"game-api/service"
	"github.com/gin-gonic/gin"
)

func main() {
	log := config.Logger()
	repo, err := postgres.NewRepositories()
	if err != nil {
		log.Fatal(err)
	}
	defer repo.Close()
	_ = repo.AutoMigrate()

	//todo: redis db
	_, err = redis.NewRedisDB(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	service := service.NewService(repo)
	handler := handler.NewHandler(service)
	r := gin.Default()
	r.Use(middleware.CORSMiddleware()) //For CORS

	//user routes
	r.POST("/users", handler.SaveUser)
	r.GET("/users", handler.GetUsers)
	r.GET("/users/:user_id", handler.GetUser)

	//authentication routes
	r.POST("/login", handler.Login)
	r.POST("/logout", handler.Logout)

	//Starting the service
	appPort := config.Env().GetServerPort()
	log.Fatal(r.Run(":" + fmt.Sprint(appPort)))
}
