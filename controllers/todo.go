package controllers

import (
	"fmt"
	"net/http"

	"github.com/ahdaan98/config"
	"github.com/ahdaan98/models"
	"github.com/gin-gonic/gin"
)

func Show(c *gin.Context) {
	var Lists []models.Todo
	config.DB.Find(&Lists)

	c.HTML(http.StatusOK, "index.html", gin.H{
		"Lists": Lists,
	})
}

func AddTask(c *gin.Context) {
	task := c.PostForm("task")

	newTask := models.Todo{Task: task}

	if err := config.DB.Create(&newTask).Error; err != nil {
		fmt.Println(err)
	}

	c.Redirect(http.StatusSeeOther, "/")
}

func DeleteTask(c *gin.Context) {
	id := c.Query("id")

	if err := config.DB.Delete(&models.Todo{}, id).Error; err != nil {
		fmt.Println(err)
		c.String(http.StatusInternalServerError, "Error deleting task")
		return
	}

	c.Redirect(http.StatusSeeOther, "/")
}

func CompleteTask(c *gin.Context){
	id := c.Query("id")

	var task models.Todo
	if err:=config.DB.Model(&task).Where("id = ?", id).Update("completed", "1").Error; err != nil {
		fmt.Println(err)
	}

	c.Redirect(http.StatusSeeOther,"/")
}

func Uncomplete(c *gin.Context){
	id := c.Query("id")

	var task models.Todo
	if err:=config.DB.Model(&task).Where("id = ?", id).Update("completed", "0").Error; err != nil {
		fmt.Println(err)
	}

	c.Redirect(http.StatusSeeOther,"/")
}