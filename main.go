package main

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"time"
	"user-api/configs"
	"user-api/docs"
	userHandler "user-api/internal/handler"
	"user-api/internal/service"
	"user-api/internal/storage"
	"user-api/pkg/middlewares"
)

func main() {

	e := echo.New()

	docs.SwaggerInfo.Host = "localhost:8080/"

	baseGroup := BuildEchoEssentials(e)

	mc, ctx := BuildMongoEssentials()

	defer func() {
		if err := mc.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	userRepository := storage.NewRepository(mc.Database(configs.Configs.DBName).Collection(configs.Configs.CollectionName))

	userService := service.NewService(userRepository)

	userHandler.NewHandler(baseGroup, userRepository, userService)

	log.Fatalln(e.Start(":8080"))
}

func BuildMongoEssentials() (*mongo.Client, context.Context) {
	opts := options.Client().ApplyURI(configs.Configs.MongoConnectionURI)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	mc, err := mongo.Connect(ctx, opts)
	if err != nil {
		panic(err)
	}

	return mc, ctx
}

func BuildEchoEssentials(e *echo.Echo) *echo.Group {
	e.HideBanner = true
	baseGroup := e.Group("/") // api routing
	baseGroup.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete, http.MethodOptions},
	}))
	baseGroup.Use(middlewares.PanicExceptionHandling())

	baseGroup.GET("swagger/*", echoSwagger.WrapHandler)

	return baseGroup
}
