package models

type User struct {
	Base
	Name      string
	Email     string
	Username  string `gorm:"uniqueIndex"`
	Password  string
	Profile   Profile `gorm:"foreignKey:ProfileID"`
	ProfileID uint64
	Skill     []Skill `gorm:"many2many:user_skills;"`
}
