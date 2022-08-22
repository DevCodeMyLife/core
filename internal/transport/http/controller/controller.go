package controller

import (
	"context"
	"core/internal/services"
	"core/internal/transport/http/router"
	"github.com/iris-contrib/middleware/cors"
	"github.com/ivahaev/go-logger"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/accesslog"
	"os"
)

//type Controller struct {
//	Services *services.Services
//}

func NewController(ctx context.Context, services *services.Services) {
	ac := makeAccessLog()
	defer func(ac *accesslog.AccessLog) {
		err := ac.Close()
		if err != nil {
			logger.Error(err)
		}
	}(ac)

	app := iris.New()
	app.UseRouter(ac.Handler)

	iris.RegisterOnInterrupt(func() {
		err := app.Shutdown(ctx)
		if err != nil {
			logger.Error(err)
		}
	})

	crs := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowCredentials: true,
		Debug:            true,
	})
	app.UseRouter(crs)

	irisRouter := router.NewRouter(app, services)

	err := irisRouter.Listen(":8080", iris.WithoutInterruptHandler, iris.WithoutServerError(iris.ErrServerClosed))
	if err != nil {
		logger.Error(err)
	}
}

func makeAccessLog() *accesslog.AccessLog {
	ac := accesslog.File("./access.log")
	ac.AddOutput(os.Stdout)

	ac.Delim = '|'
	ac.TimeFormat = "2006-01-02 15:04:05"
	ac.Async = false
	ac.IP = true
	ac.BytesReceivedBody = true
	ac.BytesSentBody = true
	ac.BytesReceived = false
	ac.BytesSent = false
	ac.BodyMinify = true
	ac.RequestBody = true
	ac.ResponseBody = false
	ac.KeepMultiLineError = true
	ac.PanicLog = accesslog.LogHandler

	ac.SetFormatter(&accesslog.JSON{
		Indent:    "  ",
		HumanTime: true,
	})

	return ac
}
