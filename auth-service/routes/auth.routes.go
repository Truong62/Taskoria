package routes

import (
	"github.com/Truong62/taskoria/auth-service/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.Engine) {
	auth := r.Group("/api/v1/auth")
	{
		auth.POST("/register", controllers.Register)
	}
}
