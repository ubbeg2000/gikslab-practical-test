package dto

import (
	"errors"
	"time"
)

type ActivityParticipant struct {
	ID      uint64   `json:"id"`
	Name    string   `json:"name"`
	Profile string   `json:"profile"`
	Skill   []string `json:"skill"`
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

type RegisterActivityBody struct {
	Skill        string    `json:"skill"`
	Title        string    `json:"title"`
	Description  string    `json:"description"`
	StartDate    time.Time `json:"startdate"`
	EndDate      time.Time `json:"enddate"`
	Participants []string  `json:"participants"`
}

func (b RegisterActivityBody) Validate() error {
	if b.Skill == "" {
		return errors.New("skill must not be empty")
	}

	if b.Title == "" {
		return errors.New("title must not be empty")
	}

	if b.StartDate.IsZero() {
		return errors.New("startdate must not be empty")
	}

	if b.EndDate.IsZero() {
		return errors.New("enddate must not be empty")
	}

	return nil
}

type UpdateActivityBody struct {
	Skill        string    `json:"skill"`
	Title        string    `json:"title"`
	Description  string    `json:"description"`
	StartDate    time.Time `json:"startdate"`
	EndDate      time.Time `json:"enddate"`
	Participants []string  `json:"participants"`
}
