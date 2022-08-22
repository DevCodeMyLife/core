package handler

import "github.com/kataras/iris/v12"

func (h *Handler) GetAllUsers(ctx iris.Context) {
	token, err := ctx.Request().Cookie("access_token")
	if err != nil {
		ctx.StatusCode(404)
		_ = ctx.JSON(iris.Map{
			"status": iris.Map{
				"code":    1,
				"message": "Your token is nil",
			},
			"data": nil,
		})
		return
	}

	users := h.Service.User.GetAllUsers(token)

	ctx.StatusCode(200)
	_ = ctx.JSON(iris.Map{
		"status": iris.Map{
			"message": nil,
		},
		"data": users,
	})
}
