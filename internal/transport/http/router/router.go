package router

import (
	"core/internal/services"
	"core/internal/transport/http/handler"
	"github.com/kataras/iris/v12"
)

type Router struct {
	iris *iris.Application
}

func NewRouter(iris *iris.Application, services *services.Services) *iris.Application {
	serv := handler.Handler{
		Service: services,
	}

	v1 := iris.Party("/v1")

	// System
	system := v1.Party("/system")
	system.Handle("GET", "/health_check", serv.HealthCheck)

	// User
	v1.Handle("GET", "/users", serv.GetAllUsers)
	v1.Handle("GET", "/user/{id:int64}", serv.GetUser)

	return iris
}
