package models

import "time"

type Target struct {
	ID          int64      `json:"id" db:"id"`
	MissionID   int64      `json:"mission_id" db:"mission_id"`
	Name        string     `json:"name" db:"name"`
	Country     string     `json:"country" db:"country"`
	Notes       string     `json:"notes" db:"notes"`
	Completed   bool       `json:"completed" db:"completed"`
	CompletedAt *time.Time `json:"completed_at" db:"completed_at"`
	CreatedAt   time.Time  `json:"created_at" db:"created_at"`
}


