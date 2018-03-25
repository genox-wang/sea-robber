package router

import (
	"sea-robber-api/app/controller"
)

func route() {

	api := router.Group("/api")
	{
		api.POST("/user/", controller.UserCreatePost)
		api.POST("/user/:uuid", controller.UserUpdatePost)

		api.GET("/rank/", controller.RankAllGet)
	}
}
