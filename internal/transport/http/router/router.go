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

	// system
	iris.Handle("GET", "/health_check", serv.HealthCheck)

	// user
	iris.Handle("POST", "/users", serv.GetAllUsers) // get all users

	return iris
}
