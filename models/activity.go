package models

import "time"

type Activity struct {
	Base
	Skill       Skill
	Title       string    `json:"title"`
	Description string    `json:"description"`
	StartDate   time.Time `json:"startdate"`
	EndDate     time.Time `json:"enddate"`
}
