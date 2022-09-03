package routes

import (
	"github.com/a1nn1997/go-auth/controllers"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {

	//user routes
	app.Post("/api/register", controllers.Register)
	app.Post("/api/edituser", controllers.EditUser)
	app.Post("/api/login", controllers.Login)
	app.Get("/api/user", controllers.User)
	app.Post("/api/logout", controllers.Logout)

	//post routes
	app.Post("/api/post", controllers.CreatePost)
	app.Post("/api/post/:id", controllers.EditPost)
	app.Get("/api/post", controllers.GetPost)
	app.Get("/api/post/:id", controllers.GetPostById)
	app.Get("/api/postbyuser", controllers.PostByUserName)

	//Post and Subreddit routes
	app.Get("/api/post/bysubreddit/:id", controllers.PostBySubRedditId)
	//Subreddit routes
	app.Post("/api/subreddit", controllers.CreateSubreddit)
	app.Get("/api/subreddit", controllers.GetSubreddit)
	app.Get("/api/subreddit/:id", controllers.GetSubredditById)

	//votes routes
	app.Post("/api/votes", controllers.VotePost)

	//comment routes
	app.Post("/api/comment", controllers.AddComment)
	app.Get("/api/comment", controllers.GetComment)
	app.Get("/api/comment/:id", controllers.GetCommentById)
	app.Post("/api/comment/:id", controllers.EditComment)
	app.Get("/api/comment/bypostid/:id", controllers.GetCommentByPostId)
	app.Get("/api/comment/byuser", controllers.GetCommentByUserId)
	
}
