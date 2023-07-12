package models

import "time"

type Activity struct {
	Base
	Skill        Skill `gorm:"foreignKey:SkillID"`
	SkillID      uint64
	Title        string
	Description  string
	StartDate    time.Time
	EndDate      time.Time
	Participants []User `gorm:"many2many:activity_users;"`
}
