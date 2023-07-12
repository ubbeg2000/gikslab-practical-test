package services

import (
	"gikslab-practical-test/bootstrap"
	"gikslab-practical-test/dto"
	"gikslab-practical-test/models"
)

func ListSkills() []dto.Skill {
	db := bootstrap.DB

	var skills []models.Skill
	db.Find(&skills)

	var dtoSkills []dto.Skill = []dto.Skill{}
	for _, s := range skills {
		dtoSkills = append(dtoSkills, dto.Skill{
			ID:        s.ID,
			SkillName: s.SkillName,
		})
	}

	return dtoSkills
}
