package routers

import (
	"ding/core/http/controllers"

	"github.com/gin-gonic/gin"
)

func Register() *gin.Engine {
	router := gin.Default()

	index := new(controllers.IndexController)
	router.GET("/", index.Index)

	ding := new(controllers.DingController)
	router.POST("/ding/bot", ding.Bot)

	return router
}
