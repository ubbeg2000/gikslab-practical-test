package services

import (
	"errors"
	"gikslab-practical-test/bootstrap"
	"gikslab-practical-test/dto"
	"gikslab-practical-test/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func RegisterActivity(body dto.RegisterActivityBody) error {
	db := bootstrap.DB

	var skill models.Skill
	if err := db.First(&skill, "skill_name = ?", body.Skill).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	var users []models.User
	db.Joins("JOIN user_skills ON user_skills.user_id = users.id AND user_skills.skill_id = ?", skill.ID).Find(&users, "name IN ?", body.Participants)

	activity := models.Activity{
		SkillID:      skill.ID,
		Skill:        skill,
		Title:        body.Title,
		Description:  body.Description,
		StartDate:    body.StartDate,
		EndDate:      body.EndDate,
		Participants: users,
	}
	return db.Create(&activity).Error
}

func ListActivities(skillID uint64, page int, limit int, sortBy string, sortOrder string) []dto.Activity {
	db := bootstrap.DB

	var activities []models.Activity
	db.Preload("Participants.Skill").Preload("Skill").
		Limit(limit).Offset((page-1)*limit).
		Order(clause.OrderByColumn{
			Column: clause.Column{Name: sortBy},
			Desc:   sortOrder == "desc",
		}).
		Order(clause.OrderByColumn{
			Column: clause.Column{Name: "start_date"},
			Desc:   false,
		}).
		Find(&activities, "skill_id = ?", skillID)

	var dtoActivities []dto.Activity = []dto.Activity{}
	for _, a := range activities {
		var dtoParticipants []dto.ActivityParticipant
		for _, p := range a.Participants {
			var skills []string
			for _, s := range p.Skill {
				skills = append(skills, s.SkillName)
			}

			dtoParticipants = append(dtoParticipants, dto.ActivityParticipant{
				ID:      p.ID,
				Name:    p.Name,
				Profile: p.Profile,
				Skill:   skills,
			})
		}
		dtoActivities = append(dtoActivities, dto.Activity{
			SkillID:      a.SkillID,
			SkillName:    a.Skill.SkillName,
			Title:        a.Title,
			Description:  a.Description,
			StartDate:    a.StartDate,
			EndDate:      a.EndDate,
			Participants: dtoParticipants,
		})
	}

	return dtoActivities
}

func UpdateActivity(activityID uint64, body dto.UpdateActivityBody) error {
	db := bootstrap.DB

	var activity models.Activity
	var skill models.Skill
	var users []models.User

	if err := db.First(&activity, activityID).Error; err != nil {
		return err
	}

	if body.Title != "" {
		activity.Title = body.Title
	}

	if body.Description != "" {
		activity.Description = body.Description
	}

	if body.StartDate.IsZero() {
		activity.StartDate = body.StartDate
	}

	if body.EndDate.IsZero() {
		activity.EndDate = body.EndDate
	}

	if body.Skill != "" {
		if err := db.First(&skill, "skill_name = ?", body.Skill).Error; errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}

		activity.SkillID = skill.ID
		activity.Skill = skill
	}

	if len(body.Participants) != 0 {
		db.Model(&activity).Association("Participants").Clear()
		db.Joins("JOIN user_skills ON user_skills.user_id = users.id AND user_skills.skill_id = ?", skill.ID).Find(&users, "name IN ?", body.Participants)
		activity.Participants = users
	}

	return db.Save(&activity).Error
}

func DeleteActivity(activityID uint64) error {
	db := bootstrap.DB

	var activity models.Activity
	if err := db.First(&activity, activityID).Error; err != nil {
		return err
	}

	db.Model(&activity).Association("Participants").Clear()
	return db.Model(&activity).Association("Participants").Delete(&activity)
}
