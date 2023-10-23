package main

import (
	"machship-go/adapter/redis"
	"machship-go/api"
	"machship-go/core/usecase"
	"machship-go/util/logger"

	"github.com/gin-gonic/gin"
	goredis "github.com/go-redis/redis/v8"
)

func main() {
	l := logger.NewLogger()
	l.Info("Starting Github API Service...")
	rdb := goredis.NewClient(&goredis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})

	repo := redis.NewRedisRepository(rdb)
	usecaseService := usecase.NewGithubService(repo)
	handler := api.NewHandler(usecaseService, l)

	r := gin.Default()
	r.POST("/retrieveUsers", handler.GetGithubUsers)
	r.Run()
}
