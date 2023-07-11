package dto

import "time"

type ActivityParticipant struct {
	ID      uint64 `json:"id"`
	Name    string `json:"name"`
	Profile string `json:"profile"`
	Skill   string `json:"skill"`
}

type Activity struct {
	SkillID      uint64                `json:"skill_id"`
	SkillName    string                `json:"skill_name"`
	Title        string                `json:"title"`
	Description  string                `json:"description"`
	StartDate    time.Time             `json:"startdate"`
	EndDate      time.Time             `json:"enddate"`
	Participants []ActivityParticipant `json:"participants"`
}

type RegisterActivityResponse BaseResponse

type RegisterActivityBody struct {
	Skill       string    `json:"skill"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	StartDate   time.Time `json:"startdate"`
	EndDate     time.Time `json:"enddate"`
}

type UpdateActivityResponse BaseResponse

type UpdateActivityBody struct {
	Skill       string    `json:"skill"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	StartDate   time.Time `json:"startdate"`
	EndDate     time.Time `json:"enddate"`
}

type DeleteActivityResponse BaseResponse
