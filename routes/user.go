package routes

import (
	"github.com/SanjaySinghRajpoot/interactWrapper/controller"
	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {

	// All user routes
	router.GET("api/getURL", controller.GetURL)
	router.POST("api/getInteractions", controller.GetInteractions)
}
