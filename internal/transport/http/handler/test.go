package handler

import (
	"github.com/ivahaev/go-logger"
	"github.com/kataras/iris/v12"
)

func (s *Services) HealthCheck(ctx iris.Context) {
	ctx.StatusCode(200)

	err := ctx.JSON(iris.Map{
		"status": iris.Map{
			"message": nil,
		},
		"data": nil,
	})
	if err != nil {
		logger.Error(err)
	}

	return
}
