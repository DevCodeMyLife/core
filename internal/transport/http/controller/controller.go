package controller

import (
	"context"
	"core/internal/services"
	"core/internal/transport/http/router"
	"github.com/ivahaev/go-logger"

	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris/v12"
)

type Controller struct {
	Iris     *iris.Application
	Services *services.Services
}

func NewController(ctx context.Context, services *services.Services) {
	server := &Controller{
		Iris:     iris.Default(),
		Services: services,
	}

	iris.RegisterOnInterrupt(func() {
		err := server.Iris.Shutdown(ctx)
		if err != nil {
			logger.Error(err)
		}
	})

	crs := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"*"},
		Debug:            false,
	})

	server.Iris.UseRouter(crs)

	irisRouter := router.NewRouter(server.Iris, services)

	err := irisRouter.Listen(":8080", iris.WithoutInterruptHandler, iris.WithoutServerError(iris.ErrServerClosed))
	if err != nil {
		logger.Error(err)
	}
}
