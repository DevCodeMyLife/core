package handler

import (
	"github.com/kataras/iris/v12"
)

func (h *Handler) Cors(ctx iris.Context) {
	ctx.Header("Access-Control-Allow-Origin", "asdasdasd")
	ctx.Header("Access-Control-Allow-Credentials", "true")

	if ctx.Method() == iris.MethodOptions {
		ctx.Header("Access-Control-Methods",
			"POST, PUT, PATCH, DELETE")

		ctx.Header("Access-Control-Allow-Headers",
			"Access-Control-Allow-Origin,Content-Type")

		ctx.Header("Access-Control-Max-Age",
			"86400")

		ctx.StatusCode(iris.StatusNoContent)
		return
	}

	ctx.Next()
}
