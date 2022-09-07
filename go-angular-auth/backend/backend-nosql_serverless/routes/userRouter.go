package routes

import(
	controller "github.com/a1nn1997/redditclone/controller"
	"github.com/a1nn1997/redditclone/middleware"
	"github.com/gin-gonic/gin"

)
func UserRoutes(incomingRoutes *gin.Engine){
	//user apis
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/users/all",controller.GetUsers())
	incomingRoutes.GET("/users/:user_id",controller.GetUserById())
	incomingRoutes.POST("/api/edituser",controller.EditUser())
	
	// post apis
	incomingRoutes.GET("/api/post",controller.GetPost())
	incomingRoutes.GET("/api/post/:id",controller.GetPostById())
	incomingRoutes.POST("/api/post",controller.CreatePost())
	incomingRoutes.POST("/api/post/:id",controller.EditPost())
	incomingRoutes.GET("/api/postbyuser",controller.PostByUserName())
	incomingRoutes.GET("/api/post/bysubreddit/:id",controller.PostBySubRedditId())
	
	// Subreddit apis
	incomingRoutes.GET("/api/subreddit",controller.GetSubreddit())
	incomingRoutes.GET("/api/subreddit/:id",controller.GetSubredditById())
	incomingRoutes.POST("/api/subreddit",controller.CreateSubreddit())
	
	// comments apis
	incomingRoutes.GET("/api/comment",controller.GetComment())
	incomingRoutes.GET("/api/comment/:id",controller.GetCommentById())
	incomingRoutes.POST("/api/comment",controller.AddComment())
	incomingRoutes.POST("/api/comment/:id",controller.EditComment())
	incomingRoutes.GET("/api/comment/byuser",controller.GetCommentByUserId())
	incomingRoutes.GET("/api/comment/bypostid/:id",controller.GetCommentByPostId())
	
	// vote apis
	incomingRoutes.POST("/api/votes", controller.VotePost())
}
