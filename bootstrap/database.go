package bootstrap

import (
	"gikslab-practical-test/helpers"
	"gikslab-practical-test/models"
	"os"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var DB *gorm.DB = nil

var profiles []models.Profile = []models.Profile{
	{
		Base: models.Base{
			ID:        1,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		Name: "board",
	},
	{
		Base: models.Base{
			ID:        2,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		Name: "expert",
	},
	{
		Base: models.Base{
			ID:        3,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		Name: "trainer",
	},
	{
		Base: models.Base{
			ID:        4,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		Name: "competitor",
	},
}

func Database() {
	db, err := gorm.Open(sqlite.Open(os.Getenv("DATABASE_FILE")), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	if err := db.AutoMigrate(&models.Skill{}, &models.User{}, &models.Activity{}); err != nil {
		panic(err.Error())
	}

	h, _ := helpers.HashPassword("root")

	db.Clauses(clause.OnConflict{
		DoNothing: true,
	}).CreateInBatches(profiles, len(profiles))

	firstUser := models.User{
		Name:     "root",
		Email:    "root@root.com",
		Username: "root",
		Password: h,
		Profile: models.Profile{
			Base: models.Base{
				ID:        1,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			Name: "board",
		},
	}
	db.Clauses(clause.OnConflict{
		DoNothing: true,
	}).Create(&firstUser)

	DB = db
}
