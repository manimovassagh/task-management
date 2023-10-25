package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/manimovassagh/task-management/database"
	"github.com/manimovassagh/task-management/models"
)

func SetupRoutes(r *gin.Engine) {

	r.GET("/tasks", func(c *gin.Context) {
		var tasks []models.Task
		result := database.DB.Find(&tasks)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
			return
		}
		c.JSON(http.StatusOK, tasks)
	})

	r.GET("/", func(c *gin.Context) {

		c.JSON(http.StatusOK, gin.H{
			"status": "healthy: up and running",
		})
	})

	r.POST("/tasks", func(c *gin.Context) {
		var newTask models.Task

		if err := c.ShouldBindJSON(&newTask); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		result := database.DB.Create(&newTask)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
			return
		}

		c.JSON(http.StatusCreated, newTask)
	})

	r.GET("/search", func(c *gin.Context) {
		descriptionQuery := c.Query("description")

		if descriptionQuery == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Description query parameter is required"})
			return
		}

		var tasks []models.Task
		result := database.DB.Where("description LIKE ?", "%"+descriptionQuery+"%").Find(&tasks)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
			return
		}

		c.JSON(http.StatusOK, tasks)
	})
	r.GET("/search/:description", func(c *gin.Context) {
		descriptionParam := c.Param("description")

		var tasks []models.Task
		result := database.DB.Where("description LIKE ?", "%"+descriptionParam+"%").Find(&tasks)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
			return
		}

		c.JSON(http.StatusOK, tasks)
	})
}
