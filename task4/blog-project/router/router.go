package router

import (
	"blog-main/controller"
	"blog-main/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	api := r.Group("api")
	{
		user := api.Group("/user")
		{
			user.POST("/register", controller.Register)
			user.POST("/login", controller.Login)
		}
		post := api.Group("/post")
		{
			post.Use(middleware.JWTAuthMiddleware())
			{
				post.POST("/create", controller.CreatePost)
				post.GET("/get", controller.GetPosts)
				post.GET("/get/:id", controller.GetPost)
				post.DELETE("/delete/:id", controller.DeletePost)
				post.PUT("/update/:id", controller.UpdatePost)
			}
		}
		comment := api.Group("/comment")
		{
			comment.Use(middleware.JWTAuthMiddleware())
			{
				comment.POST("/create", controller.CreateComment)
				comment.GET("/get", controller.GetComments)
				comment.DELETE("/delete/:id", controller.DeleteComment)
				comment.PUT("/update/:id", controller.UpdateComment)
				comment.GET("/get/:id", controller.GetComment)
				comment.GET("/get/post/:id", controller.GetCommentsByPost)
			}
		}
	}
	return r
}
