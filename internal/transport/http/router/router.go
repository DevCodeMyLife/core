package router

import (
	"core/internal/transport/http/handler"
	"github.com/kataras/iris/v12"
)

type Router struct {
	iris *iris.Application
}

func NewRouter(iris *iris.Application) *iris.Application {
	services := handler.Services{}

	iris.Handle("GET", "/health_check", services.HealthCheck)

	return iris
}
