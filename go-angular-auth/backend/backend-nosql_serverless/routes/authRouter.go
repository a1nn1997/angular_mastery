package routes

import(
	controller "github.com/a1nn1997/redditclone/controller"
	"github.com/gin-gonic/gin"

)
func AuthRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.POST("/api/register",controller.SignUp())
	incomingRoutes.POST("/api/login",controller.Login())
}