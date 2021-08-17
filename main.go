package main

import (
	"context"
	"fmt"
	"food-app/config"
	"food-app/handler"
	"food-app/middleware"
	"food-app/repository/postgres"
	"food-app/repository/redis"
	"food-app/service"
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
