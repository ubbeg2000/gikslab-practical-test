package models

type User struct {
	Base
	Name     string
	Email    string
	Username string `gorm:"uniqueIndex"`
	Password string
	Profile  string
	Skill    []Skill `gorm:"many2many:user_skills;"`
}
