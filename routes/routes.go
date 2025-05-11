package routes

import (
	"example.com/todogo/controller"
	"github.com/gin-gonic/gin"
)

func MainRoutes(r *gin.Engine) {
	r.GET("/", controller.Ping)

	post := r.Group("/post")
	{
		post.POST("/", controller.CreatePost)
		post.GET("/", controller.GetPosts)
		post.GET("/:id", controller.GetPost)
		post.PUT("/:id", controller.PutPost)
		post.DELETE("/:id", controller.DeletePost)
	}
}
