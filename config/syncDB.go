package config

import "github.com/ahdaan98/models"

func SyncDb() {
	DB.AutoMigrate(&models.Todo{})
}