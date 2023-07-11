package bootstrap

import (
	"gikslab-practical-test/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB = nil

func Database() {
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	if err := db.AutoMigrate(&models.Skill{}, &models.User{}, &models.Activity{}); err != nil {
		panic(err.Error())
	}

	DB = db
}
