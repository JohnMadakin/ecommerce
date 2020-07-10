package routes

import (
	"github.com/gin-gonic/gin"
	controllers "github.com/ecommerce/controllers"
  // middlewares "github.com/ecommerce/middleware"
)

func Routes(router *gin.Engine){
	router.GET("/", controllers.Welcome)
	router.POST("/signup")
}
