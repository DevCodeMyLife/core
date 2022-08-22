package handler

import (
	"github.com/kataras/iris/v12"
	"time"
)

func (h *Handler) HealthCheck(ctx iris.Context) {
	ctx.StatusCode(200)

	_ = ctx.JSON(iris.Map{
		"status": iris.Map{
			"message": nil,
		},
		"data": iris.Map{
			"time": time.Now().Unix(),
		},
	})
}
