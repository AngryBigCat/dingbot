package controllers

import "github.com/gin-gonic/gin"

type IndexController struct {
	BaseController
}

func (cc IndexController) Index(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello World",
	})
}
