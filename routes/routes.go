package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jatinsaini249/ms-blog/controllers"
	"github.com/jatinsaini249/ms-blog/services"
)

// Initialize => To initialize and setup routes
func Initialize() *gin.Engine {
	router := gin.Default()

	// Basic Authentication middleware
	authorized := router.Group("ms-blog", gin.BasicAuth(gin.Accounts{
		"jatin.saini": "12345",
	}))

	postCtrl := controllers.NewPostController(services.NewPostService())
	commentCtrl := controllers.NewCommentController(services.NewCommentService())

	// routes to handle blog posts
	authorized.GET("blogs", postCtrl.HandleGetBlogPosts)
	authorized.POST("blog", postCtrl.HandleSaveBlogs)
	authorized.PUT("blog/:id", postCtrl.HandleUpdateBlogPost)
	authorized.DELETE("blog/:id", postCtrl.HandleDeleteBlogPost)

	//routes to handle comments on blog post
	authorized.POST("post/:id/comment", commentCtrl.HandleSaveComment)
	authorized.POST("post/:id/comment/:parentId", commentCtrl.HandleSaveComment)
	authorized.GET("post/:id/comments", commentCtrl.HandleGetPostComments)
	authorized.GET("comment/:id", commentCtrl.HandleGetComment)
	authorized.PUT("comment/:id", commentCtrl.HandleUpdateComment)
	authorized.DELETE("comment/:id", commentCtrl.HandleDeleteComment)

	return router
}
