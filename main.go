package main

import (
	"github.com/ahdaan98/config"
	"github.com/ahdaan98/routes"
	"github.com/gin-gonic/gin"
)

func init(){
	config.DBconnect()
	config.SyncDb()
	R.LoadHTMLGlob("templates/*")
}

var R = gin.Default()

func main() {
	routes.TodoRoutes(R)

	R.Run()
}