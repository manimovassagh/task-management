package database

import (
	"time"

	"github.com/manimovassagh/task-management/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := "user:password@tcp(127.0.0.1:4444)/task?parseTime=true"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the database: " + err.Error())
	}

	db.AutoMigrate(&models.Task{})
	sampleTask1 := models.Task{
		Title:       "Sample Task 1",
		Description: "This is the first sample task",
		DueDate:     "2023-11-01",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	db.Create(&sampleTask1)

	DB = db
}
