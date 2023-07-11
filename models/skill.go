package models

type Skill struct {
	Base
	SkillName string `gorm:"index,unique"`
}
