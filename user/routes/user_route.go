package routes

import (
	"github.com/A-junaid-K/pixel_vogue/user/api"
	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {
	router.LoadHTMLGlob("templates/*.html")
	r := router.Group("/user")
	{
		r.POST("/signup", api.SignUp)
	}

}
