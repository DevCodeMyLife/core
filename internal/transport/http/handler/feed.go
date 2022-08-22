package handler

import (
	"github.com/kataras/iris/v12"
)

func (h *Handler) GetAllFeeds(ctx iris.Context) {
	ctx.StatusCode(200)

	feeds := h.Service.Feed.GetAllFeeds()

	_ = ctx.JSON(iris.Map{
		"status": iris.Map{
			"message": nil,
		},
		"data": feeds,
	})
}

func (h *Handler) GetFeedByID(ctx iris.Context) {
	id, err := ctx.Params().GetInt64("id")
	if err != nil {
		ctx.StatusCode(400)
		_ = ctx.JSON(iris.Map{
			"status": iris.Map{
				"message": "Insert id is not int",
			},
			"data": nil,
		})
		return
	}

	feeds := h.Service.Feed.GetFeedByID(id)
	if feeds.ID == 0 {
		ctx.StatusCode(404)
		_ = ctx.JSON(iris.Map{
			"status": iris.Map{
				"message": "Feed not exist",
			},
			"data": nil,
		})
		return
	}

	ctx.StatusCode(200)
	_ = ctx.JSON(iris.Map{
		"status": iris.Map{
			"message": nil,
		},
		"data": feeds,
	})
}
