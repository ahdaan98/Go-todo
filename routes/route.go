package routes

import (
	"github.com/ahdaan98/controllers"
	"github.com/gin-gonic/gin"
)

func TodoRoutes(c *gin.Engine) {
	c.GET("/",controllers.Show)
	c.POST("/add",controllers.AddTask)
	c.GET("/delete",controllers.DeleteTask)
	c.GET("/complete",controllers.CompleteTask)
	c.GET("/uncomplete",controllers.Uncomplete)
}