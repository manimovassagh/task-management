package main

import (
	"github.com/gin-gonic/gin"
	"github.com/manimovassagh/task-management/database"
	"github.com/manimovassagh/task-management/router"
)

func main() {
	r := gin.Default()
	database.ConnectDB()

	router.SetupRoutes(r)

	r.Run(":4545")
}
