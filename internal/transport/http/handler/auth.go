package handler

import (
	"github.com/kataras/iris/v12"
	"strings"
)

func (h *Handler) AuthMiddleware(ctx iris.Context) {
	rawToken := ctx.GetHeader("Authorization")

	if len(rawToken) == 0 {
		ctx.StatusCode(401)
		_ = ctx.JSON(iris.Map{
			"status": iris.Map{
				"message": "Not token required",
			},
			"data": nil,
		})
		return
	}

	contain := strings.Contains(rawToken, "Bearer")
	if !contain {
		ctx.StatusCode(401)
		_ = ctx.JSON(iris.Map{
			"status": iris.Map{
				"message": "Insert token is not validate",
			},
			"data": nil,
		})
		return
	}

	sliceToken := strings.Split(rawToken, " ")
	if len(sliceToken) == 1 || len(sliceToken) > 2 {
		ctx.StatusCode(401)
		_ = ctx.JSON(iris.Map{
			"status": iris.Map{
				"message": "Insert token is not validate",
			},
			"data": nil,
		})
		return
	}

	isAuth := h.Service.Auth.IsAuth(sliceToken[1])
	if !isAuth {
		ctx.StatusCode(401)
		_ = ctx.JSON(iris.Map{
			"status": iris.Map{
				"message": "Bad token required",
			},
			"data": nil,
		})
		return
	}

	ctx.Next()
}
