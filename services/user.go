package services

import (
	"gikslab-practical-test/bootstrap"
	"gikslab-practical-test/dto"
	"gikslab-practical-test/helpers"
	"gikslab-practical-test/models"

	"gorm.io/gorm/clause"
)

func RegisterUser(body dto.RegistrationBody) error {
	db := bootstrap.DB

	var profile models.Profile
	db.First(&profile, "name = ?", body.Profile)

	var skills []models.Skill
	for _, s := range body.Skill {
		skills = append(skills, models.Skill{
			SkillName: s,
		})
	}

	db.Clauses(clause.OnConflict{
		DoNothing: true,
	}).CreateInBatches(&skills, len(skills))

	db.Find(&skills, "skill_name IN ?", body.Skill)

	h, _ := helpers.HashPassword(body.Password)
	u := models.User{
		Name:     body.Name,
		Email:    body.Email,
		Username: body.Username,
		Password: h,
		Profile:  profile,
		Skill:    skills,
	}

	return db.Create(&u).Error
}
