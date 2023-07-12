package models

type Profile struct {
	Base
	Name string `gorm:"uniqueIndex"`
}
