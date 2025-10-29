package models

import (
	"time"
)

type Mission struct {
	ID          int64      `json:"id" db:"id"`
	CatID       *int64     `json:"cat_id" db:"cat_id"`
	Completed   bool       `json:"completed" db:"completed"`
	CompletedAt *time.Time `json:"completed_at" db:"completed_at"`
	CreatedAt   time.Time  `json:"created_at" db:"created_at"`
}

type RequestMission struct {
	Mission Mission  `json:"mission"`
	Targets []Target `json:"targets"`
}

type ResponseMission struct {
	Cat     Cat      `json:"cat"`
	Mission Mission  `json:"mission"`
	Targets []Target `json:"targets"`
}
