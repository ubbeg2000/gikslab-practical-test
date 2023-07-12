package bootstrap

import (
	"gikslab-practical-test/helpers"
	"gikslab-practical-test/models"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB = nil

func Database() {
	db, err := gorm.Open(sqlite.Open(os.Getenv("DATABASE_FILE")), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	if err := db.AutoMigrate(&models.Skill{}, &models.User{}, &models.Activity{}); err != nil {
		panic(err.Error())
	}

	h, _ := helpers.HashPassword("root")
	firstUser := models.User{
		Name:     "root",
		Email:    "root@root.com",
		Username: "root",
		Password: h,
		Profile:  "expert",
	}

	db.Create(&firstUser)

	DB = db
}
